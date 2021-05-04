package tools

//import (
//	"bencurio/inverter_exporter/models"
//	"fmt"
//	mqtt "github.com/eclipse/paho.mqtt.golang"
//	log "github.com/sirupsen/logrus"
//	"os"
//	"os/signal"
//	"syscall"
//)
//
//func NewMQTT(config *models.Config) (mqtt.Client, error) {
//	mqttBroker := fmt.Sprintf("%s://%s:%d", config.MQTT.Schema, config.MQTT.Broker, config.MQTT.Port)
//
//	options := mqtt.NewClientOptions()
//	options.AddBroker(mqttBroker)
//	options.SetClientID(config.MQTT.ClientID)
//
//	if len(config.MQTT.Username) > 0 {
//		options.SetUsername(config.MQTT.Username)
//		options.SetPassword(config.MQTT.Password)
//	}
//
//	options.SetAutoReconnect(true)
//
//	options.OnConnect = func(client mqtt.Client) {
//		log.Infof("mqtt.OnConnect")
//	}
//	options.OnConnectionLost = func(client mqtt.Client, err error) {
//		log.Warnf("mqtt.OnConnectionLost: %v", err)
//	}
//	options.OnReconnecting = func(client mqtt.Client, options *mqtt.ClientOptions) {
//		log.Warnf("mqtt.OnReconnecting")
//	}
//	options.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
//		log.Warnf("mqtt.DefaultPublishHandler: %v", message)
//	})
//
//	client := mqtt.NewClient(options)
//	if token := client.Connect(); token.Wait() && token.Error() != nil {
//		log.Fatalf("mqtt.Connect: %v", token.Error())
//	}
//
//	closeSignals := make(chan os.Signal, 1)
//	signal.Notify(closeSignals, syscall.SIGINT, syscall.SIGTERM)
//
//	go func() {
//		<-closeSignals
//		log.Infof("mqtt.Close")
//		defer client.Disconnect(250)
//	}()
//
//	return client, nil
//}
