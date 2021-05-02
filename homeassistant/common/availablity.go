package common

type Availability struct {
	PayloadAvailable    string `json:"payload_available,omitempty" default:"online"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty" default:"offline"`
	Topic               string `json:"topic,omitempty"`
}
