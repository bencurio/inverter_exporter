package inverter_exporter

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"os"
)

// ProtocolType specifies whether the given command is a query (get) or a setting (set).
type ProtocolType string

const (
	PROTOCOL_TYPE_GET ProtocolType = "get"
	PROTOCOL_TYPE_SET ProtocolType = "set"
)

// ProtocolRegexType the Regex types that can be specified in the mapping
type ProtocolRegexType string

const (
	PROTOCOL_REGEX_TYPE_REPLACE_ALL_STRING ProtocolRegexType = "ReplaceAllString"
	PROTOCOL_REGEX_TYPE_MATCH_STRING       ProtocolRegexType = "MatchString"
)

type ProtocolScannerType string

const (
	PROTOCOL_SCANNER_TYPE_SCAN_WORDS ProtocolScannerType = "ScanWords"
)

type Protocols struct {
	Protocols []*Protocol `validate:"required" yaml:"protocols"`
}

// Protocol contains the structure of the protocol.example.yaml file.
type Protocol struct {
	Command        string             `yaml:"command"`
	Type           ProtocolType       `validate:"required" yaml:"type"`
	ResponseLength int                `yaml:"response_length"`
	Mapping        []*ProtocolMapping `yaml:"mapping"`
}

type ProtocolMapping struct {
	Scanner   ProtocolScannerType       `yaml:"scanner" default:"ScanWords"`
	Key       string                    `validate:"required" yaml:"key"` // must be unique
	Index     int                       `validate:"required" yaml:"index"`
	Regex     *ProtocolMappingRegex     `yaml:"regex,omitempty"`
	Command   string                    `yaml:"command,omitempty"`   // In GET mode
	Responses *ProtocolMappingResponses `yaml:"responses,omitempty"` // In GET mode
}

// ProtocolMappingResponses can take two states. We can determine which state means what based on the Inverter's
// response. For example, if the Inverter's response is ACK, then in this case the state is OK, whereas NAK is FAIL.
type ProtocolMappingResponses struct {
	Ok   string `yaml:"ok"`
	Fail string `yaml:"fail"`
}

type ProtocolMappingRegex struct {
	Type    ProtocolRegexType `yaml:"type" default:"MatchString"`
	Pattern string            `yaml:"pattern"`
	Replace string            `yaml:"replace"`
}

func NewProtocol(protocolFile string) (*Protocols, error) {
	protocols := &Protocols{}

	stat, err := os.Stat(protocolFile)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, fmt.Errorf("'%s' is a directory, not a file", protocolFile)
	}

	file, err := os.Open(protocolFile)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&protocols); err != nil {
		return nil, err
	}

	validate = validator.New()
	if err := validate.Struct(protocols); err != nil {
		return nil, err
	}

	//// Check the "key" is unique
	var tmpKeys []string
	for _, protocol := range protocols.Protocols {
		for _, m := range protocol.Mapping {
			for _, tmpKey := range tmpKeys {
				if tmpKey == m.Key {
					return nil, fmt.Errorf("the '%s' key must be unique in protocol spec (%v)", m.Key, m)
				}
			}
			tmpKeys = append(tmpKeys, m.Key)
		}
	}

	//// Validate struct
	//// TODO Looking for better validation
	for _, protocol := range protocols.Protocols {
		if len(protocol.Type) == 0 {
			return nil, fmt.Errorf("the type is required (%v)", protocol)
		}
		if len(protocol.Mapping) == 0 {
			return nil, fmt.Errorf("the mapping is required (%v)", protocol)
		}
		for _, m := range protocol.Mapping {
			if len(m.Key) == 0 {
				return nil, fmt.Errorf("the key is required (%v)", m)
			}
		}
	}

	return protocols, nil
}
