package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/cover.mqtt/

type Cover struct { //nolint:maligned
	ConfigTopic            string              `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability           common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode       string              `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic      string              `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	CommandTopic           string              `yaml:"CommandTopic,omitempty" json:"command_topic,omitempty"`
	Device                 common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	DeviceClass            string              `yaml:"DeviceClass,omitempty" json:"device_class,omitempty"`
	JSONAttributesTemplate string              `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string              `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	Name                   string              `yaml:"Name,omitempty" json:"name" default:"MQTT Alarm"`
	Optimistic             string              `yaml:"Optimistic,omitempty" json:"optimistic" default:"false"`
	PayloadAvailable       string              `yaml:"PayloadAvailable,omitempty" json:"payload_available" default:"online"`
	PayloadClose           string              `yaml:"PayloadClose,omitempty" json:"payload_close" default:"CLOSE"`
	PayloadNotAvailable    string              `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available" default:"offline"`
	PayloadOpen            string              `yaml:"PayloadOpen,omitempty" json:"payload_open" default:"OPEN"`
	PayloadStop            string              `yaml:"PayloadStop,omitempty" json:"payload_stop" default:"STOP"`
	PositionClosed         int                 `yaml:"PositionClosed,omitempty" json:"position_closed" default:"0"`
	PositionOpen           int                 `yaml:"PositionOpen,omitempty" json:"position_open" default:"100"`
	PositionTemplate       string              `yaml:"PositionTemplate,omitempty" json:"position_template,omitempty"`
	PositionTopic          string              `yaml:"PositionTopic,omitempty" json:"position_topic,omitempty"`
	QOS                    int                 `yaml:"QOS,omitempty" json:"qos" default:"0"`
	Retain                 bool                `yaml:"Retain,omitempty" json:"retain" default:"false"`
	SetPositionTemplate    string              `yaml:"SetPositionTemplate,omitempty" json:"set_position_template,omitempty"`
	SetPositionTopic       string              `yaml:"SetPositionTopic,omitempty" json:"set_position_topic,omitempty"`
	StateClosed            string              `yaml:"StateClosed,omitempty" json:"state_closed" default:"closed"`
	StateClosing           string              `yaml:"StateClosing,omitempty" json:"state_closing" default:"closing"`
	StateOpen              string              `yaml:"StateOpen,omitempty" json:"state_open" default:"open"`
	StateOpening           string              `yaml:"StateOpening,omitempty" json:"state_opening" default:"opening"`
	StateStopped           string              `yaml:"StateStopped,omitempty" json:"state_stopped" default:"stopped"`
	StateTopic             string              `yaml:"StateTopic,omitempty" json:"state_topic,omitempty"`
	TiltClosedValue        int                 `yaml:"TiltClosedValue,omitempty" json:"tilt_closed_value" default:"0"`
	TiltCommandTemplate    string              `yaml:"TiltCommandTemplate,omitempty" json:"tilt_command_template,omitempty"`
	TiltCommandTopic       string              `yaml:"TiltCommandTopic,omitempty" json:"tilt_command_topic,omitempty"`
	TiltMax                int                 `yaml:"TiltMax,omitempty" json:"tilt_max" default:"100"`
	TiltMin                int                 `yaml:"TiltMin,omitempty" json:"tilt_min" default:"0"`
	TiltOpenedValue        int                 `yaml:"TiltOpenedValue,omitempty" json:"tilt_opened_value" default:"100"`
	TiltOptimistic         bool                `yaml:"TiltOptimistic,omitempty" json:"tilt_optimistic" default:"true"`
	TiltStatusTemplate     string              `yaml:"TiltStatusTemplate,omitempty" json:"tilt_status_template,omitempty"`
	TiltStatusTopic        string              `yaml:"TiltStatusTopic,omitempty" json:"tilt_status_topic,omitempty"`
	UniqueID               string              `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
	ValueTemplate          string              `yaml:"ValueTemplate,omitempty" json:"value_template,omitempty"`
}
