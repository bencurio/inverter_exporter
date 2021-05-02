package homeassistant

import (
	"bencurio/inverter_exporter/homeassistant/sensor"
	"encoding/json"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/yaml.v2"
	"reflect"
)

type Config struct {
	HomeAssistant []ConfigSensor `yaml:"homeassistant"`
}

type ConfigSensor struct {
	Type   Sensor      `yaml:"type"`
	Key    string      `yaml:"key"`
	Config interface{} `yaml:"config"`
}

type HomeAssistant interface {
	ConfigureSensor(sensor ConfigSensor) error
	UpdateSensor(sensor interface{}, value interface{}) error
	SubscribeSensor(sensor interface{}) error
}

type AbstractHomeAssistant struct{}

func NewHomeAssistant(client mqtt.Client) (HomeAssistant, error) {
	return &homeassistant{
		client: client,
	}, nil
}

type homeassistant struct {
	client mqtt.Client
}

func (h homeassistant) SubscribeSensor(sensor interface{}) error {
	// TODO Implement me
	panic("implement me")
}

func (h homeassistant) ConfigureSensor(sensor ConfigSensor) error {
	if !h.client.IsConnected() {
		return errors.New("mqtt client isn't connected")
	}

	sensorRaw, err := json.Marshal(sensor.Config)
	if err != nil {
		return err
	}

	sensorStruct, err := getSensorStructByTypeName(sensor.Type)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(sensorRaw, sensorStruct); err != nil {
		return err
	}

	// Only used for the MQTT topic.
	configTopic := reflect.ValueOf(sensorStruct).Elem().FieldByName("ConfigTopic").String()
	reflect.ValueOf(sensorStruct).Elem().FieldByName("ConfigTopic").SetString("")

	sensorPayload, err := json.Marshal(sensorStruct)
	if err != nil {
		return err
	}

	if token := h.client.Publish(configTopic, 0, false, sensorPayload); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

type tmpUpdateSensor struct {
	Name       string `yaml:"name" json:"name"`
	UniqueID   string `yaml:"UniqueID" json:"unique_id"`
	StateTopic string `yaml:"StateTopic" json:"state_topic"`
}

func (h homeassistant) UpdateSensor(sensor interface{}, value interface{}) error {
	if !h.client.IsConnected() {
		return errors.New("mqtt client isn't connected")
	}

	buffer, err := json.Marshal(sensor)
	if err != nil {
		return err
	}

	var tmpSensor tmpUpdateSensor
	if err := json.Unmarshal(buffer, &tmpSensor); err != nil {
		return err
	}

	var output string
	switch reflect.ValueOf(sensor).Type().String() {
	case "sensor.Sensor":
		output = fmt.Sprintf("%v", value)
	case "sensor.BinarySensor":
		output = fmt.Sprintf("%v", value)
	default:
		return errors.New("unsupported sensor")
	}

	if token := h.client.Publish(tmpSensor.StateTopic, 0, false, output); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

type tmpSensor struct {
	Type   Sensor      `yaml:"type"`
	Key    string      `yaml:"key"`
	Config interface{} `yaml:"config"`
}

func (c *ConfigSensor) UnmarshalYAML(unmarshal func(interface{}) error) error {
	configData := &tmpSensor{}

	if err := unmarshal(&configData); err != nil {
		return err
	}

	var realResult interface{}

	switch configData.Type {
	case ALARM_CONTROL_PANEL:
		realResult = &sensor.AlarmControlPanel{}
	case BINARY_SENSOR:
		realResult = &sensor.BinarySensor{}
	case CAMERA:
		realResult = &sensor.Camera{}
	case CLIMATE:
		realResult = &sensor.Climate{}
	case COVER:
		realResult = &sensor.Cover{}
	case DEVICE_TRACKER:
		realResult = &sensor.DeviceTracker{}
	case DEVICE_TRIGGER:
		realResult = &sensor.DeviceTrigger{}
	case FAN:
		realResult = &sensor.Fan{}
	case LIGHT:
		realResult = &sensor.Light{}
	case LOCK:
		realResult = &sensor.Lock{}
	case SCENE:
		realResult = &sensor.Scene{}
	case SENSOR:
		realResult = &sensor.Sensor{}
	case SWITCH:
		realResult = &sensor.Switch{}
	case TAG:
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
