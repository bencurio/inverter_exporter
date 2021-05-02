package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// Sensor platform uses the MQTT message payload as the sensor value. If messages in this state_topic are published with
// RETAIN flag, the sensor will receive an instant update with last known value. Otherwise, the initial state will
// be undefined.
type Sensor struct {
	ConfigTopic            string               `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability           *common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode       string               `yaml:"AvailabilityMode,omitempty" json:"availability_mode,omitempty" default:"online"` // TODO enum
	AvailabilityTopic      string               `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	Device                 *common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	DeviceClass            string               `yaml:"DeviceClass,omitempty" json:"device_class,omitempty"`
	ExpireAfter            int                  `yaml:"ExpireAfter,omitempty" json:"expire_after,omitempty"`
	ForceUpdate            bool                 `yaml:"ForceUpdate,omitempty" json:"force_update,omitempty" default:"false"`
	Icon                   string               `yaml:"Icon,omitempty" json:"icon,omitempty"`
	JSONAttributesTemplate string               `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string               `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	Name                   string               `yaml:"Name,omitempty" json:"name,omitempty" default:"MQTT Alarm"`
	PayloadAvailable       string               `yaml:"PayloadAvailable,omitempty" json:"payload_available,omitempty" default:"online"`
	PayloadNotAvailable    string               `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available,omitempty" default:"offline"`
	QOS                    int                  `yaml:"QOS,omitempty" json:"qos,omitempty" default:"0"`
	StateTopic             string               `validate:"required" yaml:"StateTopic,omitempty" json:"state_topic,omitempty"`
	UniqueID               string               `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
	UnitOfMeasurement      string               `yaml:"UnitOfMeasurement,omitempty" json:"unit_of_measurement,omitempty"`
	ValueTemplate          string               `yaml:"ValueTemplate,omitempty" json:"value_template,omitempty"`
}

// @see https://www.home-assistant.io/integrations/sensor.mqtt/
