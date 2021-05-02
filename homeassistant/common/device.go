package common

type Device struct {
	Connections  string `json:"connections,omitempty"` // TODO
	Identifiers  string `json:"identifiers,omitempty"` // TODO
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
	Name         string `json:"name,omitempty"`
	SwVersion    string `json:"sw_version,omitempty"`
	ViaDevice    string `json:"via_device,omitempty"`
}
