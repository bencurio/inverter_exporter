package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/switch.mqtt/

type Switch struct { //nolint:maligned
	ConfigTopic            string               `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability           *common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode       string               `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic      string               `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	CommandTopic           string               `yaml:"CommandTopic,omitempty" json:"command_topic,omitempty"`
	Device                 *common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	Icon                   string               `yaml:"Icon,omitempty" json:"icon"`
	JSONAttributesTemplate string               `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string               `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	Name                   string               `yaml:"Name,omitempty" json:"name" default:"MQTT Alarm"`
	Optimistic             bool                 `yaml:"Optimistic,omitempty" json:"optimistic" default:"false"`
	PayloadAvailable       string               `yaml:"PayloadAvailable,omitempty" json:"payload_available" default:"online"`
	PayloadNotAvailable    string               `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available" default:"offline"`
	PayloadOff             string               `yaml:"PayloadOff,omitempty" json:"payload_off,omitempty" default:"OFF"`
	PayloadOn              string               `yaml:"PayloadOn,omitempty" json:"payload_on,omitempty" default:"ON"`
	QOS                    int                  `yaml:"QOS,omitempty" json:"qos" default:"0"`
	Retain                 bool                 `yaml:"Retain,omitempty" json:"retain" default:"false"`
	StateOff               string               `yaml:"StateOff,omitempty" json:"state_off,omitempty" default:"payload_off"`
	StateOn                string               `yaml:"StateOn,omitempty" json:"state_on,omitempty" default:"payload_on"`
	StateTopic             string               `yaml:"StateTopic,omitempty" json:"state_topic,omitempty"`
	UniqueID               string               `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
	ValueTemplate          string               `yaml:"ValueTemplate,omitempty" json:"value_template,omitempty"`
}
