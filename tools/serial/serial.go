package serial

import (
	"bufio"
	"fmt"
	jserial "github.com/jacobsa/go-serial/serial"
	log "github.com/sirupsen/logrus"
	"github.com/snksoft/crc"
	"io"
	"time"
)

type Config struct {
	PortName        string             `validate:"required" yaml:"port_name" default:"/dev/ttyUSB0"`
	BraudRate       uint               `validate:"required,numeric" yaml:"braud_rate" default:"2400"`
	DataBits        uint               `validate:"required,numeric" yaml:"data_bits" default:"8"`
	StopBits        uint               `validate:"required,numeric" yaml:"stop_bits" default:"1"`
	ParityMode      jserial.ParityMode `yaml:"parity_mode" default:"0"`
	MinimumReadSize uint               `validate:"required,numeric" yaml:"minimum_read_size" default:"3"`
	CRC             CRC                `yaml:"crc"`
}

type CRC struct {
	Table     string `yaml:"table"`
	FirstByte byte   `yaml:"first_byte" default:"0x28"`
	EOLByte   byte   `yaml:"eol_byte" default:"0x0D"`
}

type Serial interface {
	Open() error
	Close() error
	Read(length int) ([]byte, error)
	ReadWithTimeout(length int, timeout time.Duration) ([]byte, error)
	Write(data []byte) (int, error)
	CheckCRC(data []byte) error
	AppendCRC(data []byte) ([]byte, error)
}

type serial struct {
	config Config
	port   io.ReadWriteCloser
}

// Read data from serial
func (s serial) Read(length int) ([]byte, error) {
	r := bufio.NewReader(s.port)
	lvl := log.GetLevel()
	var data []byte
	sum := 0
	for {
		sum++
		tkn, err := r.ReadByte()
		if err != nil {
			break
		}
		if lvl == log.DebugLevel {
			log.Debug(string(tkn))
			log.Debug(tkn)
		}
		data = append(data, tkn)
		if sum == length {
			break
		}
	}

	log.Debugf("serial.Read: %s", string(data))
	log.Debugf("serial.Read: %v", data)
	log.Debugf("serial.Read: %d bytes", len(data))

	// If this value is set in the configuration, a CRC check is performed.
	if len(s.config.CRC.Table) != 0 {
		if err := s.CheckCRC(data); err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (s serial) ReadWithTimeout(length int, timeout time.Duration) ([]byte, error) {
	var buffer []byte
	done := make(chan error)

	go func() {
		var err error
		buffer, err = s.Read(length)
		done <- err
	}()

	select {
	case err := <-done:
		return buffer, err
	case <-time.After(time.Second * timeout):
		return nil, fmt.Errorf("timeout")
	}
}

// Write data to serial
func (s *serial) Write(data []byte) (int, error) {
	// If this value is set in the configuration, a CRC check is performed.
	if len(s.config.CRC.Table) != 0 {
		var err error
		data, err = s.AppendCRC(data)
		if err != nil {
			return 0, fmt.Errorf("AppendCRC: %w", err)
		}
	}

	log.Debugf("serial.Write: %s", string(data))
	log.Debugf("serial.Write: %v", data)

	n, err := s.port.Write(data)
	if err != nil {
		return 0, fmt.Errorf("port.Write: %w", err)
	}

	log.Debugf("serial.Write: %d bytes", n)

	return n, nil
}

func (s serial) AppendCRC(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data")
	}

	switch s.config.CRC.Table {
	case "XMODEM":
		crcTable := crc.NewTable(crc.XMODEM)
		crcCalc := crcTable.CalculateCRC(data)
		return append(
			data,
			byte(0xff&(crcCalc>>8)),
			byte(0xff&(crcCalc>>0)),
			s.config.CRC.EOLByte,
		), nil
	default:
		return nil, fmt.Errorf("unsupported crc algorithm: %s", s.config.CRC.Table)
	}
}

func (s serial) CheckCRC(data []byte) error {
	dataLength := len(data)

	if dataLength == 0 {
		return fmt.Errorf("empty data")
	}

	if s.config.CRC.FirstByte != 0 && data[0] != s.config.CRC.FirstByte {
		return fmt.Errorf("invalid first byte")
	}

	if s.config.CRC.EOLByte != 0 && data[dataLength-1] != s.config.CRC.EOLByte {
		return fmt.Errorf("invalid eol byte")
	}

	switch s.config.CRC.Table {
	case "XMODEM":
		if dataLength < 3 {
			return fmt.Errorf("data is too short")
		}

		var crcCheck uint64
		crcCheck <<= 8
		crcCheck += uint64(data[dataLength-3])
		crcCheck <<= 8
		crcCheck += uint64(data[dataLength-2])

		// Payload: [<FirstByte>]<PAYLOAD>[<CRC><CR>]
		crcTable := crc.NewTable(crc.XMODEM)
		crcCalc := crcTable.CalculateCRC(data[:dataLength-3])
		if crcCheck != crcCalc {
			return fmt.Errorf("crc mismatch (%d != %d)", crcCheck, crcCalc)
		}
		return nil
	default:
		return fmt.Errorf("unsupported crc algorithm: %s", s.config.CRC.Table)
	}
}

// Open creates an io.ReadWriteCloser based on the supplied options struct.
func (s *serial) Open() error {
	options := jserial.OpenOptions{
		PortName:        s.config.PortName,
		BaudRate:        s.config.BraudRate,
		DataBits:        s.config.DataBits,
		StopBits:        s.config.StopBits,
		ParityMode:      s.config.ParityMode,
		MinimumReadSize: s.config.MinimumReadSize,
	}

	if len(s.config.CRC.Table) == 0 {
		log.Warnf("this serial communication is used without CRC checking")
	}

	var err error
	s.port, err = jserial.Open(options)

	if err != nil {
		return fmt.Errorf("serial.Open: %w", err)
	}

	return nil
}

func (s serial) Close() error {
	return s.port.Close()
}

func NewSerial(config Config) (Serial, error) {
	return &serial{
		config: config,
	}, nil
}
