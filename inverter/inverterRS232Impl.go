package inverter

import (
	"bencurio/inverter_exporter"
	"bencurio/inverter_exporter/memdb"
	"bencurio/inverter_exporter/tools/serial"
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"time"
)

type inverterRS232Impl struct {
	config inverter_exporter.Config
	serial serial.Serial
	memdb  memdb.MemDB
}

func (i inverterRS232Impl) ReadKeys() (map[string]interface{}, error) {
	return i.memdb.GetAll()
}

func (i inverterRS232Impl) ReadKey(key string) (interface{}, error) {
	return i.memdb.Get(key)
}

//func (i inverterRS232Impl) RawRead() {
//	return i.serial.Read()
//}
//
//func (i inverterRS232Impl) RawWrite(data []byte) {
//
//}

func NewInverterRS232(config inverter_exporter.Config, memdb *memdb.MemDB) (Inverter, error) {
	if config.Inverter.Communication.Type != inverter_exporter.CommunicationTypeRS232 {
		return nil, fmt.Errorf("invalid inverter communication type: %s", config.Inverter.Communication.Type)
	}

	serialConfig := config.Inverter.Communication.Config.(*serial.Config)
	serialPort, _ := serial.NewSerial(*serialConfig)

	if err := serialPort.Open(); err != nil {
		return nil, err
	}

	inv := &inverterRS232Impl{
		config: config,
		memdb:  *memdb,
		serial: serialPort,
	}
	if err := inv.Run(); err != nil {
		return nil, err
	}

	return inv, nil
}

func (i inverterRS232Impl) Run() error {
	go func() {
		// FIXME Hard-coded query time!
		for range time.Tick(time.Second * 10) {
			if err := i.readAllSensors(); err != nil {
				log.Warnf("InverterRS232Impl.readAllSensors: %v", err)
			}
		}
	}()
	return nil
}

func (i inverterRS232Impl) readAllSensors() error {
	for _, sensors := range i.config.Inverter.Communication.Protocol.Protocols {

		if sensors.Type == inverter_exporter.PROTOCOL_TYPE_SET {
			continue
		}

		if _, err := i.serial.Write([]byte(sensors.Command)); err != nil {
			log.Errorf("serial.Write: %v", err)
			continue
		}

		data, err := i.serial.ReadWithTimeout(5)
		if err != nil {
			log.Errorf("serial.ReadWithTimeout: %v", err)
		}

		if err := i.rawHandler(sensors.Command, data); err != nil {
			log.Errorf("InverterRS232Impl.rawHandler: %v", err)
		}
	}
	return nil
}

func (i inverterRS232Impl) rawHandler(command string, data []byte) error {

	// <(><PAYLOAD><CRC><CR> ----> <PAYLOAD>
	data = data[1 : len(data)-3]

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Split(bufio.ScanWords)

	sum := 0
	index := map[int]string{}

	for scanner.Scan() {
		index[sum] = scanner.Text()
		sum++
	}

	protocolSpec := &inverter_exporter.Protocol{}
	fail := true

	// Find relevant protocol structure based on the command
	for _, protocol := range i.config.Inverter.Communication.Protocol.Protocols {
		if protocol.Command == command {
			protocolSpec = protocol
			fail = false
			break
		}
	}

	if fail {
		return fmt.Errorf("the command does not exist: %s", command)
	}

	for _, m := range protocolSpec.Mapping {
		if v, ok := index[m.Index]; ok {
			// Compile regex
			if m.Regex != nil {
				v, _ = i.compileRegex(v, m.Regex)
			}

			if err := i.memdb.Set(m.Key, v); err != nil {
				return fmt.Errorf("unable to write into the memdb: %w", err)
			}
		}
	}

	return nil
}

func (i inverterRS232Impl) compileRegex(value string, data *inverter_exporter.ProtocolMappingRegex) (string, error) {
	regex, err := regexp.Compile(data.Pattern)
	if err != nil {
		return "", fmt.Errorf("unable to compile regex: %w", err)
	}

	var newValue string

	switch data.Type {
	case inverter_exporter.PROTOCOL_REGEX_TYPE_MATCH_STRING:
		if regex.MatchString(value) {
			newValue = "1"
		} else {
			newValue = "0"
		}
	case inverter_exporter.PROTOCOL_REGEX_TYPE_REPLACE_ALL_STRING:
	default:
		newValue = regex.ReplaceAllString(value, data.Replace)
	}

	// TODO: Better support for variable types

	return newValue, nil
}
