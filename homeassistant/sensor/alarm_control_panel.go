package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// AlarmControlPanel The mqtt alarm panel platform enables the possibility to control MQTT capable alarm panels.
// The Alarm icon will change state after receiving a new state from state_topic. If these messages are published with
// RETAIN flag, the MQTT alarm panel will receive an instant state update after subscription and will start with the
// correct state. Otherwise, the initial state will be unknown.
type AlarmControlPanel struct { //nolint:maligned
	ConfigTopic            string               `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Availability           *common.Availability `yaml:"Availability,omitempty" json:"availability,omitempty"`
	AvailabilityMode       string               `yaml:"AvailabilityMode,omitempty" json:"availability_mode" default:"online"` // TODO enum
	AvailabilityTopic      string               `yaml:"AvailabilityTopic,omitempty" json:"availability_topic,omitempty"`
	Code                   string               `yaml:"Code,omitempty" json:"code,omitempty"`
	CodeArmRequired        bool                 `yaml:"CodeArmRequired,omitempty" json:"code_arm_required" default:"true"`
	CodeDisarmRequired     bool                 `yaml:"CodeDisarmRequired,omitempty" json:"code_disarm_required" default:"true"`
	CommandTemplate        string               `yaml:"CommandTemplate,omitempty" json:"command_template" default:"action"` // TODO enum
	CommandTopic           string               `validate:"required" yaml:"CommandTopic,omitempty" json:"command_topic,omitempty"`
	Device                 *common.Device       `yaml:"Device,omitempty" json:"device,omitempty"`
	JSONAttributesTemplate string               `yaml:"JsonAttributesTemplate,omitempty" json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string               `yaml:"JsonAttributesTopic,omitempty" json:"json_attributes_topic,omitempty"`
	Name                   string               `yaml:"Name,omitempty" json:"name" default:"MQTT Alarm"`
	PayloadArmAway         string               `yaml:"PayloadArmAway,omitempty" json:"payload_arm_away" default:"ARM_AWAY"`
	PayloadArmHome         string               `yaml:"PayloadArmHome,omitempty" json:"payload_arm_home" default:"ARM_HOME"`
	PayloadArmNight        string               `yaml:"PayloadArmNight,omitempty" json:"payload_arm_night" default:"ARM_NIGHT"`
	PayloadArmCustomBypass string               `yaml:"PayloadArmCustomBypass,omitempty" json:"payload_arm_custom_bypass" default:"ARM_CUSTOM_BYPASS"`
	PayloadAvailable       string               `yaml:"PayloadAvailable,omitempty" json:"payload_available" default:"online"`
	PayloadDisarm          string               `yaml:"PayloadDisarm,omitempty" json:"payload_disarm" default:"DISARM"`
	PayloadNotAvailable    string               `yaml:"PayloadNotAvailable,omitempty" json:"payload_not_available" default:"offline"`
	QOS                    int                  `yaml:"QOS,omitempty" json:"qos" default:"0"`
	Retain                 bool                 `yaml:"Retain,omitempty" json:"retain" default:"false"`
	StateTopic             string               `validate:"required" yaml:"StateTopic,omitempty" json:"state_topic,omitempty"`
	UniqueID               string               `yaml:"UniqueID,omitempty" json:"unique_id,omitempty"`
	ValueTemplate          string               `yaml:"ValueTemplate,omitempty" json:"value_template,omitempty"`
}

// @see https://www.home-assistant.io/integrations/alarm_control_panel.mqtt

// AlarmControlPanelCommandState defines which commands are accepted by the Home Assistant.
type AlarmControlPanelCommandState string

// These states can be used for the alarm panel.
const (
	ALARM_CONTROL_PANEL_STATE_DISARMED            AlarmControlPanelCommandState = "disarmed"
	ALARM_CONTROL_PANEL_STATE_ARMED_HOME          AlarmControlPanelCommandState = "armed_home"
	ALARM_CONTROL_PANEL_STATE_ARMED_AWAY          AlarmControlPanelCommandState = "armed_away"
	ALARM_CONTROL_PANEL_STATE_ARMED_NIGHT         AlarmControlPanelCommandState = "armed_night"
	ALARM_CONTROL_PANEL_STATE_ARMED_CUSTOM_BYPASS AlarmControlPanelCommandState = "armed_custom_bypass"
	ALARM_CONTROL_PANEL_STATE_PENDING             AlarmControlPanelCommandState = "pending"
	ALARM_CONTROL_PANEL_STATE_TRIGGERED           AlarmControlPanelCommandState = "triggered"
	ALARM_CONTROL_PANEL_STATE_ARMING              AlarmControlPanelCommandState = "arming"
	ALARM_CONTROL_PANEL_STATE_DISARMING           AlarmControlPanelCommandState = "disarming"
)
