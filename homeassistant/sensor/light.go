package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/climate.mqtt/

type Light struct { //nolint:maligned
	ConfigTopic              string               `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability             *common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode         string               `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic        string               `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	BrightnessCommandTopic   string               `yaml:"BrightnessCommandTopic,omitempty" json:"brightness_command_topic,omitempty"`
	BrightnessScale          int                  `yaml:"BrightnessScale,omitempty" json:"brightness_scale,omitempty" default:"255"`
	BrightnessStateTopic     string               `yaml:"BrightnessStateTopic,omitempty" json:"brightness_state_topic,omitempty"`
	BrightnessValueTemplate  string               `yaml:"BrightnessValueTemplate,omitempty" json:"brightness_value_template,omitempty"`
	ColorTempCommandTemplate string               `yaml:"ColorTempCommandTemplate,omitempty" json:"color_temp_command_template,omitempty"`
	ColorTempCommandTopic    string               `yaml:"ColorTempCommandTopic,omitempty" json:"color_temp_command_topic,omitempty"`
	ColorTempStateTopic      string               `yaml:"ColorTempStateTopic,omitempty" json:"color_temp_state_topic,omitempty"`
	ColorTempValueTemplate   string               `yaml:"ColorTempValueTemplate,omitempty" json:"color_temp_value_template,omitempty"`
	CommandTopic             string               `yaml:"CommandTopic,omitempty" json:"command_topic,omitempty"`
	Device                   *common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	EffectCommandTopic       string               `yaml:"EffectCommandTopic,omitempty" json:"effect_command_topic,omitempty"`
	EffectList               string               `yaml:"EffectList,omitempty" json:"effect_list,omitempty"`
	EffectStateTopic         string               `yaml:"EffectStateTopic,omitempty" json:"effect_state_topic,omitempty"`
	EffectValueTemplate      string               `yaml:"EffectValueTemplate,omitempty" json:"effect_value_template,omitempty"`
	HSCommandTopic           string               `yaml:"HSCommandTopic,omitempty" json:"hs_command_topic,omitempty"`
	HSStateTopic             string               `yaml:"HSStateTopic,omitempty" json:"hs_state_topic,omitempty"`
	HSValueTemplate          string               `yaml:"HSValueTemplate,omitempty" json:"hs_value_template,omitempty"`
	JSONAttributesTemplate   string               `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JsonAttributesTopic      string               `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	MaxMireds                int                  `yaml:"MaxMireds,omitempty" json:"max_mireds,omitempty"`
	MinMireds                int                  `yaml:"MinMireds,omitempty" json:"min_mireds,omitempty"`
	Name                     string               `yaml:"Name,omitempty" json:"name" default:"MQTT Light"`
	OnCommandType            string               `yaml:"OnCommandType,omitempty" json:"on_command_type,omitempty"`
	Optimistic               bool                 `yaml:"Optimistic,omitempty" json:"optimistic" default:"false"`
	PayloadAvailable         string               `yaml:"PayloadAvailable,omitempty" json:"payload_available" default:"online"`
	PayloadNotAvailable      string               `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available" default:"offline"`
	PayloadOff               string               `yaml:"PayloadOff,omitempty" json:"payload_off,omitempty" default:"OFF"`
	PayloadOn                string               `yaml:"PayloadOn,omitempty" json:"payload_on,omitempty" default:"ON"`
	QOS                      int                  `yaml:"QOS,omitempty" json:"qos" default:"0"`
	Retain                   bool                 `yaml:"Retain,omitempty" json:"retain" default:"false"`
	RGBCommandTopic          string               `yaml:"RGBCommandTopic,omitempty" json:"rgb_command_topic,omitempty"`
	RGBStateTopic            string               `yaml:"RGBStateTopic,omitempty" json:"rgb_state_topic,omitempty"`
	RGBValueTemplate         string               `yaml:"RGBValueTemplate,omitempty" json:"rgb_value_template,omitempty"`
	Schema                   string               `yaml:"Schema,omitempty" json:"schema,omitempty" default:"default"`
	StateTopic               string               `yaml:"StateTopic,omitempty" json:"state_topic,omitempty"`
	StateValueTemplate       string               `yaml:"StateValueTemplate,omitempty" json:"state_value_template,omitempty"`
	UniqueID                 string               `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
	WhiteValueCommandTopic   string               `yaml:"WhiteValueCommandTopic,omitempty" json:"white_value_command_topic,omitempty"`
	WhiteValueScale          int                  `yaml:"WhiteValueScale,omitempty" json:"white_value_scale,omitempty" default:"255"`
	WhiteValueStateTopic     string               `yaml:"WhiteValueStateTopic,omitempty" json:"white_value_state_topic,omitempty"`
	WhiteValueTemplate       string               `yaml:"WhiteValueTemplate,omitempty" json:"white_value_template,omitempty"`
	XYCommandTopic           string               `yaml:"XYCommandTopic,omitempty" json:"xy_command_topic,omitempty"`
	XYStateTopic             string               `yaml:"XYStateTopic,omitempty" json:"xy_state_topic,omitempty"`
	XYValueTemplate          string               `yaml:"XYValueTemplate,omitempty" json:"xy_value_template,omitempty"`
}
