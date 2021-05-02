package sensor

import (
	"bencurio/inverter_exporter/homeassistant/common"
)

// @see https://www.home-assistant.io/integrations/tag.mqtt/

type Tag struct {
	ConfigTopic   string         `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Topic         string         `validate:"required" yaml:"Topic,omitempty" json:"topic,omitempty"`
	ValueTemplate string         `yaml:"ValueTemplate,omitempty" json:"value_template,omitempty"`
	Device        *common.Device `validate:"required" yaml:"Device,omitempty" json:"device,omitempty"`
}
