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

func init() {
	// TODO: Some of the logging-related settings will be made configurable later.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	log.Infof("log.GetLevel: %s", log.GetLevel())
}

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "./config/config.yaml", "Config file")
	flag.Parse()

	//config, err := models.NewConfig(configFile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	config, err := inverter_exporter.NewConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	Run(config)
}

func Run(config *inverter_exporter.Config) {

	sensorsMemDB := map[string]interface{}{}

	sensors, err := memdb.NewMemDB(&sensorsMemDB)
	if err != nil {
		log.Fatal(err)
	}

	if err := sensors.Clean(); err != nil {
		return
	}

	switch config.Inverter.Communication.Type {
	case inverter_exporter.CommunicationTypeRS232:
		if _, err := inverter.NewInverterRS232(*config, &sensors); err != nil {
			log.Errorf("NewInverterRS232: %v", err)
			return
		}
	}

	log.Infof("Waiting for data from the Inverter")
	time.Sleep(time.Second * 15)

	s, _ := sensors.GetAll()
	if len(s) == 0 {
		log.Fatal("No data was received from the inverter. Incorrect settings?")
	}

	log.Infof("Working ...")
	for {
		time.Sleep(time.Second * 60)
	}
}
