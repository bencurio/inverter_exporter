package inverter_exporter

import (
	"testing"
)

func TestNewProtocol(t *testing.T) {
	_, err := NewProtocol("protocol_bad.yaml")
	if err == nil {
		t.Errorf("No error occurred if an incorrect file path was specified.")
	}

	protocol, err := NewProtocol("protocol.yaml")
	if err != nil {
		t.Error(err)
	}

	t.Errorf("%v", protocol)
}
