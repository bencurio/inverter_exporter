package sensor

// @see https://www.home-assistant.io/integrations/device_tracker.mqtt/

type DeviceTracker struct {
	ConfigTopic    string `validate:"required" yaml:"ConfigTopic" json:"config_topic,omitempty"`
	Devices        string `validate:"required" yaml:"Devices,omitempty" json:"devices"` // TODO list?
	QOS            int    `yaml:"QOS,omitempty" json:"qos,omitempty" default:"0"`
	PayloadHome    string `yaml:"PayloadHome,omitempty" json:"payload_open,omitempty" default:"home"`
	PayloadNotHome string `yaml:"PayloadNotHome,omitempty" json:"payload_stop,omitempty" default:"not_home"`
	SourceType     string `yaml:"SourceType,omitempty" json:"source_type,omitempty"` // TODO enum?
}
