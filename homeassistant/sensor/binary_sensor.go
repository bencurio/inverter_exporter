package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// BinarySensor The mqtt binary sensor platform uses an MQTT message received to set the binary sensor’s state to on
// or off. The state will be updated only after a new message is published on state_topic matching payload_on or
// payload_off. If these messages are published with the retain flag set, the binary sensor will receive an instant
// state update after subscription and Home Assistant will display the correct state on startup. Otherwise, the initial
// state displayed in Home Assistant will be unknown.
type BinarySensor struct {
	ConfigTopic            string               `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability           *common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode       string               `yaml:"AvailabilityMode,omitempty" json:"availability_mode,omitempty" default:"online"`
	AvailabilityTopic      string               `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	Device                 *common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	DeviceClass            string               `yaml:"DeviceClass,omitempty" json:"device_class,omitempty"`
	ExpireAfter            int                  `yaml:"ExpireAfter,omitempty" json:"expire_after,omitempty"`
	ForceUpdate            bool                 `yaml:"ForceUpdate,omitempty" json:"force_update,omitempty" default:"false"`
	Icon                   string               `yaml:"Icon" json:"icon,omitempty"`
	JSONAttributesTemplate string               `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string               `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	Name                   string               `yaml:"Name,omitempty" json:"name,omitempty" default:"MQTT Alarm"`
	OffDelay               int                  `yaml:"OffDelay,omitempty" json:"off_delay,omitempty"`
	PayloadAvailable       string               `yaml:"PayloadAvailable,omitempty" json:"payload_available,omitempty" default:"online"`
	PayloadNotAvailable    string               `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available,omitempty" default:"offline"`
	PayloadOn              string               `yaml:"PayloadOn,omitempty" json:"payload_on,omitempty"`
	PayloadOff             string               `yaml:"PayloadOff,omitempty" json:"payload_off,omitempty"`
	QOS                    int                  `yaml:"QOS,omitempty" json:"qos,omitempty" default:"0"`
	StateTopic             string               `validate:"required" yaml:"StateTopic,omitempty" json:"state_topic,omitempty"`
	UniqueID               string               `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
	ValueTemplate          string               `yaml:"ValueTemplate,omitempty" json:"value_template,omitempty"`
}

// @see https://www.home-assistant.io/integrations/binary_sensor.mqtt/
