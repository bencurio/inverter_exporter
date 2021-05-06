package main

import (
	"bencurio/inverter_exporter"
	"bencurio/inverter_exporter/inverter"
	"bencurio/inverter_exporter/memdb"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "/etc/inverter_exporter/config.yaml", "Config file")
	flag.Parse()

	config, err := inverter_exporter.NewConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	running := make(chan struct{})
	sc := make(chan os.Signal)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		<-sc
		running <- struct{}{}
		Quit()
	}()

	Run(config, running)

	<-running
}

func Run(config *inverter_exporter.Config, running chan struct{}) {
	initLog(config)

	log.Infof("Welcome to Inverter Exporter! Version: 0.0.0-dev")

	log.Infof("Initialize MemDB for sensors")
	sensorsMemDB := memdb.MemDBStore{}
	sensors := initSensorsMemDB(sensorsMemDB)

	log.Infof("Initialize Inverter Monitor service")
	initInverterMonitor(config, &sensors)

	log.Infof("Trying to communicate with the inverter ...")
	time.Sleep(time.Second * 3)

	s, _ := sensors.GetAll()
	if len(s) == 0 {
		log.Fatal("No data was received from the inverter. Incorrect settings?")
	}
	log.Infof("Suceeded!")

	log.Infof("Initialize Exporter service")
	initExporters(config, &sensors)

	log.Infof("Working ...")
}

func initLog(config *inverter_exporter.Config) {
	if config.Inverter.Log != nil {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
		})

		log.Infof("Log level set to %s", config.Inverter.Log.Level)
		log.SetLevel(config.Inverter.Log.Level)

		if config.Inverter.Log.Level == log.DebugLevel {
			log.SetReportCaller(true)
		}

		switch config.Inverter.Log.Out {
		case "stdout":
			log.Infof("Log output set to STDOUT")
			log.SetOutput(os.Stdout)
		case "stderr":
			log.Infof("Log output set to STDERR")
			log.SetOutput(os.Stderr)
		case "file":
			if len(config.Inverter.Log.File) == 0 {
				log.Fatal("you must specify the log file name")
			}
			if file, err := os.OpenFile(config.Inverter.Log.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640); err != nil {
				log.Errorf("os.OpenFile: %v", err)
			} else {
				log.Infof("Log output set to %s", config.Inverter.Log.File)
				log.SetOutput(file)
			}
		default:
			log.Fatalf("invalid log out: %s", config.Inverter.Log.Out)
		}
	}
}

func initSensorsMemDB(sensorsMemDB memdb.MemDBStore) memdb.MemDB {
	sensors, err := memdb.NewMemDB(&sensorsMemDB)
	if err != nil {
		log.Fatal(err)
	}
	return sensors
}

func initInverterMonitor(config *inverter_exporter.Config, sensors *memdb.MemDB) {
	switch config.Inverter.Communication.Type {
	case inverter_exporter.CommunicationTypeRS232:
		log.Infof("Inverter communication mode set to RS232")
		if _, err := inverter.NewInverterRS232(*config, sensors); err != nil {
			log.Errorf("NewInverterRS232: %v", err)
			return
		}
	}
}

func initExporters(config *inverter_exporter.Config, sensors *memdb.MemDB) {
	for _, exporter := range config.Inverter.Exporters {
		switch exporter.Type {
		case inverter_exporter.ExporterTypeHomeAssistantMQTT:
			log.WithFields(log.Fields{
				"service": "homeassistant",
			}).Infof("Initialize exporter: Home Assistant via MQTT")
			exporterConfig := exporter.Config.(*inverter_exporter.ExporterConfigHomeAssistantMQTT)
			schema := exporter.Schema.(*inverter_exporter.HomeAssistantConfig)
			if _, err := inverter_exporter.NewHomeAssistant(config, exporterConfig, schema, sensors); err != nil {
				log.Fatal(err)
			}
		case inverter_exporter.ExporterTypePrometheusExporter:
			log.WithFields(log.Fields{
				"service": "prometheus",
			}).Infof("Initialize exporter: Prometheus Exporter")
			exporterConfig := exporter.Config.(*inverter_exporter.ExporterConfigPrometheusExporter)
			schema := exporter.Schema.(*inverter_exporter.PrometheusConfig)
			if _, err := inverter_exporter.NewPrometheusExporter(config, exporterConfig, schema, sensors); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("unsupported exporter: %s", exporter.Type)
		}
	}
}

func Quit() {
	log.Info("Exiting...")
	os.Exit(0)
}
