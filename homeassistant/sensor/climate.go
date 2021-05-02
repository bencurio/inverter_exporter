package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/climate.mqtt/

type Climate struct { //nolint:maligned
	ConfigTopic                    string               `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	ActionTemplate                 string               `yaml:"ActionTemplate,omitempty" json:"action_template,omitempty"`
	ActionTopic                    string               `yaml:"ActionTopic,omitempty" json:"action_topic,omitempty"` // TODO enum
	AuxCommandTopic                string               `yaml:"AuxCommandTopic,omitempty" json:"aux_command_topic,omitempty"`
	AuxStateTemplate               string               `yaml:"AuxStateTemplate,omitempty" json:"aux_state_template,omitempty"`
	AuxStateTopic                  string               `yaml:"AuxStateTopic,omitempty" json:"aux_state_topic,omitempty"`
	Availability                   common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode               string               `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic              string               `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	AwayModeCommandTopic           string               `yaml:"AwayModeCommandTopic,omitempty" json:"away_mode_command_topic,omitempty"`
	AwayModeStateTemplate          string               `yaml:"AwayModeStateTemplate,omitempty" json:"away_mode_state_template,omitempty"`
	AwayModeStateTopic             string               `yaml:"AwayModeStateTopic,omitempty" json:"away_mode_state_topic,omitempty"`
	CurrentTemperatureTemplate     string               `yaml:"CurrentTemperatureTemplate,omitempty" json:"current_temperature_template,omitempty"`
	CurrentTemperatureTopic        string               `yaml:"CurrentTemperatureTopic,omitempty" json:"current_temperature_topic,omitempty"`
	Device                         common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	FanModeCommandTemplate         string               `yaml:"FanModeCommandTemplate,omitempty" json:"fan_mode_command_template,omitempty"`
	FanModeCommandTopic            string               `yaml:"FanModeCommandTopic,omitempty" json:"fan_mode_command_topic,omitempty"`
	FanModeStateTemplate           string               `yaml:"FanModeStateTemplate,omitempty" json:"fan_mode_state_template,omitempty"`
	FanModeStateTopic              string               `yaml:"FanModeStateTopic,omitempty" json:"fan_mode_state_topic,omitempty"`
	FanModes                       string               `yaml:"FanModes,omitempty" json:"fan_modes,omitempty" default:"[\"auto\", \"low\", \"medium\", \"high\"]"`
	HoldCommandTemplate            string               `yaml:"HoldCommandTemplate,omitempty" json:"hold_command_template,omitempty"`
	HoldCommandTopic               string               `yaml:"HoldCommandTopic,omitempty" json:"hold_command_topic,omitempty"`
	HoldStateTemplate              string               `yaml:"HoldStateTemplate,omitempty" json:"hold_state_template,omitempty"`
	HoldStateTopic                 string               `yaml:"HoldStateTopic,omitempty" json:"hold_state_topic,omitempty"`
	HoldMoves                      string               `yaml:"HoldMoves,omitempty" json:"hold_moves,omitempty"`
	Initial                        int                  `yaml:"Initial,omitempty" json:"initial" default:"21"`
	JSONAttributesTemplate         string               `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic            string               `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	MaxTemp                        float32              `yaml:"MaxTemp,omitempty" json:"max_temp,omitempty"`
	MinTemp                        float32              `yaml:"MinTemp,omitempty" json:"min_temp,omitempty"`
	ModeCommandTemplate            string               `yaml:"ModeCommandTemplate,omitempty" json:"mode_command_template,omitempty"`
	ModeCommandTopic               string               `yaml:"ModeCommandTopic,omitempty" json:"mode_command_topic,omitempty"`
	ModeStateTemplate              string               `yaml:"ModeStateTemplate,omitempty" json:"mode_state_template,omitempty"`
	ModeStateTopic                 string               `yaml:"ModeStateTopic,omitempty" json:"mode_state_topic,omitempty"`
	Modes                          string               `yaml:"Modes,omitempty" json:"modes,omitempty" default:"[\"auto\", \"off\", \"cool\", \"heat\", \"dry\", \"fan_only\"]"`
	Name                           string               `yaml:"Name,omitempty" json:"name" default:"MQTT Alarm"`
	PayloadAvailable               string               `yaml:"PayloadAvailable,omitempty" json:"payload_available" default:"online"`
	PayloadNotAvailable            string               `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available" default:"offline"`
	PayloadOff                     string               `yaml:"PayloadOff,omitempty" json:"payload_off,omitempty" default:"OFF"`
	PayloadOn                      string               `yaml:"PayloadOn,omitempty" json:"payload_on,omitempty" default:"ON"`
	PowerCommandTopic              string               `yaml:"PowerCommandTopic,omitempty" json:"power_command_topic,omitempty"`
	Precision                      float32              `yaml:"Precision,omitempty" json:"precision,omitempty" default:"0.1"`
	QOS                            int                  `yaml:"QOS,omitempty" json:"qos" default:"0"`
	Retain                         bool                 `yaml:"Retain,omitempty" json:"retain" default:"false"`
	SendIfOFF                      bool                 `yaml:"SendIfOFF,omitempty" json:"send_if_off" default:"true"`
	SwingModeCommandTemplate       string               `yaml:"SwingModeCommandTemplate,omitempty" json:"swing_mode_command_template,omitempty"`
	SwingModeCommandTopic          string               `yaml:"SwingModeCommandTopic,omitempty" json:"swing_mode_command_topic,omitempty"`
	SwingModeStateTemplate         string               `yaml:"SwingModeStateTemplate,omitempty" json:"swing_mode_state_template,omitempty"`
	SwingModeStateTopic            string               `yaml:"SwingModeStateTopic,omitempty" json:"swing_mode_state_topic,omitempty"`
	SwingModes                     string               `yaml:"SwingModes,omitempty" json:"swing_modes,omitempty" default:"[\"on\", \"off\"]"`
	TemperatureCommandTemplate     string               `yaml:"TemperatureCommandTemplate,omitempty" json:"temperature_command_template,omitempty"`
	TemperatureCommandTopic        string               `yaml:"TemperatureCommandTopic,omitempty" json:"temperature_command_topic,omitempty"`
	TemperatureHighCommandTemplate string               `yaml:"TemperatureHighCommandTemplate,omitempty" json:"temperature_high_command_template,omitempty"`
	TemperatureHighCommandTopic    string               `yaml:"TemperatureHighCommandTopic,omitempty" json:"temperature_high_command_topic,omitempty"`
	TemperatureHighStateTemplate   string               `yaml:"TemperatureHighStateTemplate,omitempty" json:"temperature_high_state_template,omitempty"`
	TemperatureHighStateTopic      string               `yaml:"TemperatureHighStateTopic,omitempty" json:"temperature_high_state_topic,omitempty"`
	TemperatureLowCommandTemplate  string               `yaml:"TemperatureLowCommandTemplate,omitempty" json:"temperature_low_command_template,omitempty"`
	TemperatureLowCommandTopic     string               `yaml:"TemperatureLowCommandTopic,omitempty" json:"temperature_low_command_topic,omitempty"`
	TemperatureLowStateTemplate    string               `yaml:"TemperatureLowStateTemplate,omitempty" json:"temperature_low_state_template,omitempty"`
	TemperatureLowStateTopic       string               `yaml:"TemperatureLowStateTopic,omitempty" json:"temperature_low_state_topic,omitempty"`
	TemperatureStateTemplate       string               `yaml:"TemperatureStateTemplate,omitempty" json:"temperature_state_template"`
	TemperatureStateTopic          string               `yaml:"TemperatureStateTopic,omitempty" json:"temperature_state_topic"`
	TemperatureUnit                string               `yaml:"TemperatureUnit,omitempty" json:"temperature_unit"`
	TempStep                       string               `yaml:"TempStep,omitempty" json:"temp_step"`
	UniqueID                       string               `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
	ValueTemplate                  string               `yaml:"ValueTemplate,omitempty" json:"value_template,omitempty"`
}

type ClimateActionState string

const (
	CLIMATE_ACTION_STATE_OFF     ClimateActionState = "off"
	CLIMATE_ACTION_STATE_HEATING ClimateActionState = "heating"
	CLIMATE_ACTION_STATE_COOLING ClimateActionState = "cooling"
	CLIMATE_ACTION_STATE_DRYING  ClimateActionState = "drying"
	CLIMATE_ACTION_STATE_IDLE    ClimateActionState = "idle"
	CLIMATE_ACTION_STATE_FAN     ClimateActionState = "fan"
)

type ClimateFanMode string

const (
	CLIMATE_FAN_MODE_AUTO   ClimateFanMode = "auto"
	CLIMATE_FAN_MODE_LOW    ClimateFanMode = "low"
	CLIMATE_FAN_MODE_MEDIUM ClimateFanMode = "medium"
	CLIMATE_FAN_MODE_HIGH   ClimateFanMode = "high"
)
