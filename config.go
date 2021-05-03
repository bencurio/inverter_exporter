package inverter_exporter

import (
	"bencurio/inverter_exporter/tools/serial"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type Config struct {
	Inverter *Inverter `validate:"required" yaml:"inverter"`
}

type Inverter struct {
	Log           *Log           `yaml:"log"`
	Communication *Communication `validate:"required" yaml:"communication"`
	Exporters     []*Exporter    `validate:"required" yaml:"exporters"`
}

type Log struct {
	Level log.Level `yaml:"level" default:"info"`
	Out   string    `yaml:"out" default:"stdout"`
	File  string    `yaml:"file,omitempty"`
}

type CommunicationType string

const (
	CommunicationTypeRS232 CommunicationType = "RS232"
	// TODO: RS485, Modbus RTU
)

type Communication struct {
	Type         CommunicationType `validate:"required" yaml:"type"`
	ProtocolFile string            `validate:"required" yaml:"protocol_file"`
	Protocol     Protocols         `yaml:",inline"`
	Config       interface{}       `validate:"required" yaml:"config"`
}

type ExporterType string

const (
	ExporterTypeHomeAssistantMQTT  ExporterType = "HOMEASSISTANT_MQTT"
	ExporterTypePrometheusExporter ExporterType = "PROMETHEUS_EXPORTER"
)

type Exporter struct {
	Type       ExporterType `validate:"required" yaml:"type"`
	SchemeFile string       `validate:"required" yaml:"scheme_file"`
	Scheme     interface{}  `validate:"required" yaml:",inline"`
	Config     interface{}  `validate:"required" yaml:"config"`
}

type ExporterConfigHomeAssistantMQTT struct {
	Scheme   string `validate:"required" yaml:"scheme" default:"tcp"`
	Broker   string `validate:"required" yaml:"broker"`
	Port     int    `validate:"required" yaml:"port" default:"1883"`
	ClientID string `validate:"required" yaml:"client_id" default:"inverter_exporter_ab4c"`
	Username string `yaml:"username" default:""`
	Password string `yaml:"password" default:""`
}

type ExporterConfigPrometheusExporter struct {
	Host string `yaml:"host" default:"0.0.0.0"`
	Port int    `yaml:"port" default:"2112"`
	Path string `yaml:"path" default:"/metrics"`
}

// NewConfig loads the configuration parameters from the YAML files and returns the Config structure
func NewConfig(configFile string) (*Config, error) {
	config := &Config{}

	log.WithFields(log.Fields{
		"file": configFile,
	})

	if err := config.LoadFile(configFile); err != nil {
		return nil, err
	}

	if err := validate.Struct(config); err != nil {
		return nil, err
	}
	return config, nil
}

// LoadFile loads the given configuration file.
func (c *Config) LoadFile(configFile string) error {
	stat, err := os.Stat(configFile)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", configFile)
	}

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(file, &c); err != nil {
		return err
	}

	return nil
}

type tmpCommunication struct {
	Type         string      `validate:"required" yaml:"type"`
	ProtocolFile string      `validate:"required" yaml:"protocol_file"`
	Config       interface{} `yaml:"config"`
}

// UnmarshalYAML Loads the appropriate communication structure and protocol file.
//
// You can define the type of communication protocol in the configuration. Therefore, the appropriate struct is
// returned based on type.
//
// nolint:funlen
func (c *Communication) UnmarshalYAML(unmarshal func(interface{}) error) error {
	configData := &tmpCommunication{}

	if err := unmarshal(&configData); err != nil {
		return err
	}

	var realResult interface{}

	switch CommunicationType(configData.Type) {
	case CommunicationTypeRS232:
		realResult = &serial.Config{}
	default:
		return fmt.Errorf("unsupported inverter communication protocol: %s", configData.Type)
	}

	tmpByte, err := yaml.Marshal(configData.Config)
	if err != nil {
		return fmt.Errorf("failed to marshal config data (%w)", err)
	}
	if err := yaml.Unmarshal(tmpByte, realResult); err != nil {
		return fmt.Errorf("failed to unmarshal config data (%w)", err)
	}

	c.Type = CommunicationType(configData.Type)
	c.ProtocolFile = configData.ProtocolFile
	c.Config = realResult

	// Load data from protocol file
	tmpProtocol, err := NewProtocol(configData.ProtocolFile)
	if err != nil {
		return err
	}
	c.Protocol = *tmpProtocol

	return nil
}

type tmpExporter struct {
	Type       string                 `validate:"required" yaml:"type"`
	SchemeFile string                 `validate:"required" yaml:"scheme_file"`
	Config     map[string]interface{} `yaml:"config"`
}

// UnmarshalYAML Loads the appropriate exporter structure and configuration.
//
// You can define the type of exporter in the configuration. Therefore, the appropriate struct is returned based
// on type.
//
// nolint:funlen
func (e *Exporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	configData := &tmpExporter{}

	if err := unmarshal(&configData); err != nil {
		return err
	}

	var realResult interface{}

	switch ExporterType(configData.Type) {
	case ExporterTypeHomeAssistantMQTT:
		realResult = &ExporterConfigHomeAssistantMQTT{}
	case ExporterTypePrometheusExporter:
		realResult = &ExporterConfigPrometheusExporter{}
	default:
		return fmt.Errorf("unsupported inverter exporter: %s", configData.Type)
	}

	tmpByte, err := yaml.Marshal(configData.Config)
	if err != nil {
		return fmt.Errorf("failed to marshal config data (%w)", err)
	}
	if err := yaml.Unmarshal(tmpByte, realResult); err != nil {
		return fmt.Errorf("failed to unmarshal config data (%w)", err)
	}

	e.Type = ExporterType(configData.Type)
	e.SchemeFile = configData.SchemeFile
	e.Config = realResult

	// handle exporter config
	stat, err := os.Stat(configData.SchemeFile)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("%s is a directory, not a file", configData.SchemeFile)
	}
	file, err := ioutil.ReadFile(configData.SchemeFile)
	if err != nil {
		return err
	}
	var tmpScheme interface{}

	switch ExporterType(configData.Type) {
	case ExporterTypeHomeAssistantMQTT:
		tmpScheme = &HomeAssistantConfig{}
	case ExporterTypePrometheusExporter:
		tmpScheme = &PrometheusConfig{}
	default:
		return fmt.Errorf("unsupported inverter exporter: %s", configData.Type)
	}

	if err := yaml.Unmarshal(file, tmpScheme); err != nil {
		return err
	}

	e.Scheme = tmpScheme

	return nil
}
