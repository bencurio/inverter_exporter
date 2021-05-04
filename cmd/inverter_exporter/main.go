package main

import (
	"bencurio/inverter_exporter"
	"bencurio/inverter_exporter/inverter"
	"bencurio/inverter_exporter/memdb"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
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

	Run(config)
}

func Run(config *inverter_exporter.Config) {
	initLog(config)

	log.Infof(" - Inverter Exporter init")

	log.Info(" - MemDB Sensors init")
	sensorsMemDB := map[string]interface{}{}
	sensors := initSensorsMemDB(sensorsMemDB)

	log.Info(" - Inverter Monitor init")
	initInverterMonitor(config, &sensors)

	log.Infof("   - Waiting for the first communication with inverter")
	time.Sleep(time.Second * 15)

	s, _ := sensors.GetAll()
	if len(s) == 0 {
		log.Fatal("   - No data was received from the inverter. Incorrect settings?")
	}
	log.Infof("   - Working ...")

	log.Infof(" - Init Exporters")
	initExporters(config, &sensors)

	for {
		time.Sleep(time.Second * 60)
	}
}

func initLog(config *inverter_exporter.Config) {
	if config.Inverter.Log != nil {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})

		log.SetLevel(config.Inverter.Log.Level)

		if config.Inverter.Log.Level == log.DebugLevel {
			log.SetReportCaller(true)
		}

		switch config.Inverter.Log.Out {
		case "stdout":
			log.SetOutput(os.Stdout)
		case "stderr":
			log.SetOutput(os.Stderr)
		case "file":
			if len(config.Inverter.Log.File) == 0 {
				log.Fatal("you must specify the log file name")
			}
			if file, err := os.OpenFile(config.Inverter.Log.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640); err != nil {
				log.Errorf("os.OpenFile: %v", err)
			} else {
				log.SetOutput(file)
			}
		default:
			log.Fatalf("invalid log out: %s", config.Inverter.Log.Out)
		}
	}
}

func initSensorsMemDB(sensorsMemDB map[string]interface{}) memdb.MemDB {
	sensors, err := memdb.NewMemDB(&sensorsMemDB)
	if err != nil {
		log.Fatal(err)
	}
	return sensors
}

func initInverterMonitor(config *inverter_exporter.Config, sensors *memdb.MemDB) {
	switch config.Inverter.Communication.Type {
	case inverter_exporter.CommunicationTypeRS232:
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
			log.Infof("  - Exporter: Home Assitant via MQTT init")
			exporterConfig := exporter.Config.(*inverter_exporter.ExporterConfigHomeAssistantMQTT)
			schema := exporter.Schema.(*inverter_exporter.HomeAssistantConfig)
			if _, err := inverter_exporter.NewHomeAssistant(config, exporterConfig, schema, sensors); err != nil {
				log.Fatal(err)
			}
		case inverter_exporter.ExporterTypePrometheusExporter:
			log.Infof("  - Exporter: Prometheus Exporter")
			exporterConfig := exporter.Config.(*inverter_exporter.ExporterConfigPrometheusExporter)
			schema := exporter.Schema.(*inverter_exporter.PrometheusConfig)
			if _, err := inverter_exporter.NewPrometheusExporter(config, exporterConfig, schema, sensors); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("unsupported type of exporter: %s", exporter.Type)
		}
	}
}
