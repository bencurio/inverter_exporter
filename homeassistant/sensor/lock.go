package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/lock.mqtt/

type Lock struct {
	ConfigTopic            string               `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability           common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode       string               `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic      string               `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	CommandTopic           string               `validate:"required" yaml:"CommandTopic,omitempty" json:"command_topic,omitempty"`
	Device                 common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	JSONAttributesTemplate string               `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string               `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	Name                   string               `yaml:"Name,omitempty" json:"name" default:"MQTT Alarm"`
	Optimistic             string               `yaml:"Optimistic,omitempty" json:"optimistic" default:"false"`
	PayloadAvailable       string               `yaml:"PayloadAvailable,omitempty" json:"payload_available" default:"online"`
	PayloadLock            string               `yaml:"PayloadLock,omitempty" json:"payload_lock" default:"LOCK"`
	PayloadNotAvailable    string               `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available" default:"offline"`
	PayloadUnlock          string               `yaml:"PayloadUnlock,omitempty" json:"payload_unlock" default:"UNLOCK"`
	QOS                    int                  `yaml:"QOS,omitempty" json:"qos" default:"0"`
	Retain                 bool                 `yaml:"Retain,omitempty" json:"retain" default:"false"`
	StateLocked            string               `yaml:"StateLocked,omitempty" json:"state_locked" default:"LOCKED"`
	StateTopic             string               `yaml:"StateTopic,omitempty" json:"state_topic,omitempty"`
	StateUnlocked          string               `yaml:"StateUnlocked,omitempty" json:"state_unlocked" default:"UNLOCKED"`
	UniqueID               string               `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
	ValueTemplate          string               `yaml:"ValueTemplate,omitempty" json:"value_template,omitempty"`
}
