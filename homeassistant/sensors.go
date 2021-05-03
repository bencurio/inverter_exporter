package homeassistant

import (
	sensor2 "bencurio/inverter_exporter/homeassistant/sensor"
	"errors"
)

type Sensor string

const (
	ALARM_CONTROL_PANEL Sensor = "alarm_control_panel"
	BINARY_SENSOR       Sensor = "binary_sensor"
	CAMERA              Sensor = "camera"
	CLIMATE             Sensor = "climate"
	COVER               Sensor = "cover"
	DEVICE_TRACKER      Sensor = "device_tracker"
	DEVICE_TRIGGER      Sensor = "device_trigger"
	FAN                 Sensor = "fan"
	LIGHT               Sensor = "light"
	LOCK                Sensor = "lock"
	SCENE               Sensor = "scene"
	SENSOR              Sensor = "sensor"
	SWITCH              Sensor = "switch"
	TAG                 Sensor = "tag"
)

func GetSensorStructByTypeName(name Sensor) (interface{}, error) {
	switch name {
	case ALARM_CONTROL_PANEL:
		return &sensor2.AlarmControlPanel{}, nil
	case BINARY_SENSOR:
		return &sensor2.BinarySensor{}, nil
	case CAMERA:
		return &sensor2.Camera{}, nil
	case CLIMATE:
		return &sensor2.Climate{}, nil
	case COVER:
		return &sensor2.Cover{}, nil
	case DEVICE_TRACKER:
		return &sensor2.DeviceTracker{}, nil
	case DEVICE_TRIGGER:
		return &sensor2.DeviceTrigger{}, nil
	case FAN:
		return &sensor2.Fan{}, nil
	case LIGHT:
		return &sensor2.Light{}, nil
	case LOCK:
		return &sensor2.Lock{}, nil
	case SCENE:
		return &sensor2.Scene{}, nil
	case SENSOR:
		return &sensor2.Sensor{}, nil
	case SWITCH:
		return &sensor2.Switch{}, nil
	case TAG:
		return &sensor2.Tag{}, nil
	default:
		return nil, errors.New("unsupported sensor type")
	}
}
