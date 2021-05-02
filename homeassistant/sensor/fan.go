package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/fan.mqtt
// @see https://www.home-assistant.io/integrations/fan.mqtt/#full-configuration

type Fan struct {
	ConfigTopic              string              `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability             common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode         string              `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic        string              `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	CommandTopic             string              `validate:"required" yaml:"CommandTopic,omitempty" json:"command_topic,omitempty"`
	Device                   common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	JSONAttributesTemplate   string              `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic      string              `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	Name                     string              `yaml:"Name,omitempty" json:"name" default:"MQTT Alarm"`
	Optimistic               string              `yaml:"Optimistic,omitempty" json:"optimistic" default:"false"`
	OscillationCommandTopic  string              `yaml:"OscillationCommandTopic,omitempty" json:"oscillation_command_topic"`
	OscillationStateTopic    string              `yaml:"OscillationStateTopic,omitempty" json:"oscillation_state_topic"`
	OscillationValueTemplate string              `yaml:"OscillationValueTemplate,omitempty" json:"oscillation_value_template"`
	PayloadAvailable         string              `yaml:"PayloadAvailable,omitempty" json:"payload_available" default:"online"`
	PayloadHighSpeed         string              `yaml:"PayloadHighSpeed,omitempty" json:"payload_high_speed" default:"offline"`
	PayloadLowSpeed          string              `yaml:"PayloadLowSpeed,omitempty" json:"payload_low_speed" default:"offline"`
	PayloadMediumSpeed       string              `yaml:"PayloadMediumSpeed,omitempty" json:"payload_medium_speed" default:"offline"`
	PayloadNotAvailable      string              `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available" default:"offline"`
	PayloadOn                string              `yaml:"PayloadOn,omitempty" json:"payload_on,omitempty"`
	PayloadOff               string              `yaml:"PayloadOff,omitempty" json:"payload_off,omitempty"`
	PayloadOscillationOn     string              `yaml:"PayloadOscillationOn,omitempty" json:"payload_oscillation_on,omitempty"`
	PayloadOscillationOff    string              `yaml:"PayloadOscillationOff,omitempty" json:"payload_oscillation_off,omitempty"`
	QOS                      int                 `yaml:"QOS,omitempty" json:"qos" default:"0"`
	Retain                   bool                `yaml:"Retain,omitempty" json:"retain" default:"false"`
	SpeedCommandTopic        string              `yaml:"SpeedCommandTopic,omitempty" json:"speed_command_topic"`
	SpeedStateTopic          string              `yaml:"SpeedStateTopic,omitempty" json:"speed_state_topic"`
	SpeedValueTemplate       string              `yaml:"SpeedValueTemplate,omitempty" json:"speed_value_template"`
	Speeds                   string              `yaml:"Speeds,omitempty" json:"speeds"` // TODO enum
	StateTopic               string              `yaml:"StateTopic,omitempty" json:"state_topic,omitempty"`
	StateValueTemplate       string              `yaml:"StateValueTemplate,omitempty" json:"state_value_template"`
	UniqueID                 string              `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
}
