package inverter_exporter

import (
	"testing"
)

func TestNewProtocol(t *testing.T) {
	_, err := NewProtocol("./config/protocol_bad.yaml")
	if err == nil {
		t.Errorf("No error occurred if an incorrect file path was specified.")
	}

	_, err = NewProtocol("./config/protocol.example.yaml")
	if err != nil {
		t.Error(err)
	}
}
