package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/scene.mqtt/

type Scene struct {
	ConfigTopic         string              `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability        common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode    string              `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic   string              `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	Icon                string              `yaml:"Icon,omitempty" json:"icon"`
	Name                string              `yaml:"Name,omitempty" json:"name" default:"MQTT Switch"`
	PayloadAvailable    string              `yaml:"PayloadAvailable,omitempty" json:"payload_available" default:"online"`
	PayloadNotAvailable string              `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available" default:"offline"`
	PayloadOn           string              `yaml:"PayloadOn,omitempty" json:"payload_on,omitempty" default:"ON"`
	QOS                 int                 `yaml:"QOS,omitempty" json:"qos,omitempty" default:"0"`
	Retain              bool                `yaml:"Retain,omitempty" json:"retain" default:"false"`
	UniqueID            string              `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
}
