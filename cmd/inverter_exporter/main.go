package main

import (
	"bencurio/inverter_exporter"
	"bencurio/inverter_exporter/memdb"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"os"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

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

	fmt.Printf("%v", config.Inverter.Communication)
	fmt.Printf("%v", config.Inverter.Communication.Config)
	fmt.Printf("%v", config.Inverter.Communication.Type)
	fmt.Printf("%v", config.Inverter.Communication.Protocol)
	fmt.Printf("%v", config.Inverter.Communication.ProtocolFile)
	fmt.Printf("%v", config.Inverter.Exporters)

}
