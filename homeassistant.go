package inverter_exporter

import (
	ha "bencurio/inverter_exporter/homeassistant"
	"bencurio/inverter_exporter/homeassistant/common"
	"bencurio/inverter_exporter/homeassistant/sensor"
	sens "bencurio/inverter_exporter/homeassistant/sensor"
	"encoding/json"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/util/gconv"
	"github.com/prologic/bitcask"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"reflect"
	"time"
)

type HomeAssistantConfig struct {
	HomeAssistant []HomeAssistantConfigSensor `yaml:"homeassistant"`
}

type HomeAssistantConfigSensor struct {
	Type   ha.Sensor   `yaml:"type"`
	Key    string      `yaml:"key"`
	Config interface{} `yaml:"config"`
}

type HomeAssistant interface {
	ConfigureSensor(sensor HomeAssistantConfigSensor) error
	UpdateSensor(sensor HomeAssistantConfigSensor, value interface{}) error
	SubscribeSensor(sensor HomeAssistantConfigSensor) error
	Run() error
}

type AbstractHomeAssistant struct{}

func NewHomeAssistant(config *Config, exporterConfig *ExporterConfigHomeAssistantMQTT, schema *HomeAssistantConfig, sensors *bitcask.Bitcask) (HomeAssistant, error) {
	hoas := &homeassistant{
		config:         config,
		exporterConfig: exporterConfig,
		schema:         schema,
		sensors:        sensors,
	}

	log.Infof(" - MQTT Client ID: %s", exporterConfig.ClientID)
	log.Infof(" - MQTT Broker: %s://%s:%d", exporterConfig.Schema, exporterConfig.Broker, exporterConfig.Port)

	if err := hoas.Run(); err != nil {
		return nil, err
	}

	return hoas, nil
}

type homeassistant struct {
	mqttClient     mqtt.Client
	config         *Config
	exporterConfig *ExporterConfigHomeAssistantMQTT
	schema         *HomeAssistantConfig
	sensors        *bitcask.Bitcask
}

func (h *homeassistant) Run() error {
	if err := h.openMqttConnection(); err != nil {
		return err
	}
	if err := h.setupSensors(); err != nil {
		return err
	}
	if err := h.sensorHandler(); err != nil {
		return err
	}
	return nil
}

func (h *homeassistant) openMqttConnection() error {
	mqttOptions := mqtt.NewClientOptions()

	mqttBroker := fmt.Sprintf("%s://%s:%d", h.exporterConfig.Schema, h.exporterConfig.Broker, h.exporterConfig.Port)
	mqttOptions.AddBroker(mqttBroker)

	if len(h.exporterConfig.ClientID) != 0 {
		mqttOptions.SetClientID(h.exporterConfig.ClientID)
	}
	if len(h.exporterConfig.Username) != 0 {
		mqttOptions.SetUsername(h.exporterConfig.Username)
	}
	if len(h.exporterConfig.Password) != 0 {
		mqttOptions.SetPassword(h.exporterConfig.Password)
	}

	mqttOptions.OnConnect = func(client mqtt.Client) {
		log.Infof("mqtt.OnConnect")
	}
	mqttOptions.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Warnf("mqtt.OnConnectionLost: %v", err)
	}
	mqttOptions.OnReconnecting = func(client mqtt.Client, options *mqtt.ClientOptions) {
		log.Warnf("mqtt.OnReconnecting")
	}
	mqttOptions.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		log.Warnf("mqtt.DefaultPublishHandler: %s", message)
	})

	mqttOptions.SetAutoReconnect(true)

	h.mqttClient = mqtt.NewClient(mqttOptions)
	if token := h.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("mqtt.Connect: %w", token.Error())
	}

	return nil
}

func (h *homeassistant) setupSensors() error {
	for _, snsr := range h.schema.HomeAssistant {
		if err := h.ConfigureSensor(snsr); err != nil {
			return err
		}
	}
	return nil
}

func (h *homeassistant) sensorHandlerLoop() error {
	for _, snsr := range h.schema.HomeAssistant {
		v, err := h.sensors.Get([]byte(snsr.Key))
		if err != nil {
			log.Warnf("sensor not found: %v", err)
		}
		if err := h.UpdateSensor(snsr, v); err != nil {
			log.Warnf("update error: %v", err)
		}
	}
	return nil
}

func (h *homeassistant) sensorHandler() error {
	go func() {
		if err := h.sensorHandlerLoop(); err != nil {
			return
		}
		for range time.NewTicker(time.Second * 10).C {
			if err := h.sensorHandlerLoop(); err != nil {
				return
			}
		}
	}()
	return nil
}

func (h *homeassistant) SubscribeSensor(sensor HomeAssistantConfigSensor) error {
	// TODO Implement me
	panic("implement me")
}

func (h *homeassistant) ConfigureSensor(sensor HomeAssistantConfigSensor) error {
	if !h.mqttClient.IsConnected() {
		return errors.New("mqtt client isn't connected")
	}

	device := &common.Device{}

	serialNumber, err := h.sensors.Get([]byte("SerialNumber"))
	if err == nil {
		device.Identifiers = gconv.String(serialNumber)
	}

	firmwareVersion, err := h.sensors.Get([]byte("FirmwareVersion"))
	if err == nil {
		device.SwVersion = gconv.String(firmwareVersion)
	}

	switch sensor.Type {
	case ha.BINARY_SENSOR:
		sensor.Config.(*sens.BinarySensor).Device = device
	case ha.SENSOR:
		sensor.Config.(*sens.Sensor).Device = device
	case ha.SWITCH:
		sensor.Config.(*sens.Switch).Device = device
	}

	// Only used for the MQTT topic.
	config := reflect.ValueOf(sensor.Config).Elem()
	configTopic := config.FieldByName("ConfigTopic").String()
	if config.FieldByName("ConfigTopic").CanSet() {
		config.FieldByName("ConfigTopic").SetString("")
	}

	sensorPayload, err := json.Marshal(sensor.Config)
	if err != nil {
		return err
	}

	log.Debugf("Publish to %s", configTopic)
	log.Debug(string(sensorPayload))

	if token := h.mqttClient.Publish(configTopic, 0, false, sensorPayload); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (h *homeassistant) UpdateSensor(sensor HomeAssistantConfigSensor, value interface{}) error {
	if !h.mqttClient.IsConnected() {
		return errors.New("mqtt client isn't connected")
	}

	stateTopic := reflect.ValueOf(sensor.Config).Elem().FieldByName("StateTopic").String()

	log.Debugf("%s = %v", stateTopic, gconv.String(value))

	if token := h.mqttClient.Publish(stateTopic, 0, false, gconv.String(value)); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

type tmpSensor struct {
	Type   ha.Sensor   `yaml:"type"`
	Key    string      `yaml:"key"`
	Config interface{} `yaml:"config"`
}

func (c *HomeAssistantConfigSensor) UnmarshalYAML(unmarshal func(interface{}) error) error {
	configData := &tmpSensor{}

	if err := unmarshal(&configData); err != nil {
		return err
	}

	var realResult interface{}

	switch configData.Type {
	case ha.ALARM_CONTROL_PANEL:
		realResult = &sensor.AlarmControlPanel{}
	case ha.BINARY_SENSOR:
		realResult = &sensor.BinarySensor{}
	case ha.CAMERA:
		realResult = &sensor.Camera{}
	case ha.CLIMATE:
		realResult = &sensor.Climate{}
	case ha.COVER:
		realResult = &sensor.Cover{}
	case ha.DEVICE_TRACKER:
		realResult = &sensor.DeviceTracker{}
	case ha.DEVICE_TRIGGER:
		realResult = &sensor.DeviceTrigger{}
	case ha.FAN:
		realResult = &sensor.Fan{}
	case ha.LIGHT:
		realResult = &sensor.Light{}
	case ha.LOCK:
		realResult = &sensor.Lock{}
	case ha.SCENE:
		realResult = &sensor.Scene{}
	case ha.SENSOR:
		realResult = &sensor.Sensor{}
	case ha.SWITCH:
		realResult = &sensor.Switch{}
	case ha.TAG:
		realResult = &sensor.Tag{}
	default:
		return fmt.Errorf("unsupportred sensor: %s", configData.Type)
	}

	tmpByte, err := yaml.Marshal(configData.Config)
	if err != nil {
		return fmt.Errorf("failed to marshal config data (%w)", err)
	}
	if err := yaml.Unmarshal(tmpByte, realResult); err != nil {
		return fmt.Errorf("failed to unmarshal config data (%w)", err)
	}

	c.Key = configData.Key
	c.Type = configData.Type
	c.Config = realResult

	return nil
}
