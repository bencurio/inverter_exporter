package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// DeviceTrigger platform uses an MQTT message payload to generate device trigger events
type DeviceTrigger struct {
	ConfigTopic    string        `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	AutomationType string        `validate:"required" yaml:"AutomationType,omitempty" json:"automation_type" default:"trigger"`
	Payload        string        `yaml:"Payload,omitempty" json:"payload,omitempty"`
	QOS            int           `yaml:"QOS,omitempty" json:"qos,omitempty" default:"0"`
	Topic          string        `validate:"required" yaml:"Topic,omitempty" json:"topic"`
	Type           string        `validate:"required" yaml:"Type,omitempty" json:"type"` // TODO enum?
	Subtype        string        `validate:"required" yaml:"Subtype,omitempty" json:"subtype"`
	Device         common.Device `validate:"required" yaml:"Device,omitempty" json:"device,omitempty"`
}

// @see https://www.home-assistant.io/integrations/device_trigger.mqtt/
