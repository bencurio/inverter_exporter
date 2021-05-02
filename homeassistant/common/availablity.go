package common

// Availability a list of MQTT topics subscribed to receive availability (online/offline) updates.
type Availability struct {
	PayloadAvailable    string `json:"payload_available,omitempty" default:"online"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty" default:"offline"`
	Topic               string `json:"topic,omitempty"`
}
