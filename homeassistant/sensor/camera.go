package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/camera.mqtt/

type Camera struct {
	ConfigTopic            string              `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability           common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode       string              `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic      string              `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	Device                 common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	JSONAttributesTemplate string              `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string              `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	Name                   string              `yaml:"Name,omitempty" json:"name" default:"MQTT Alarm"`
	Topic                  string              `validate:"required" yaml:"Topic,omitempty" json:"topic"`
	UniqueID               string              `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
}
