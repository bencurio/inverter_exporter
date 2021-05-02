package inverter

type Inverter interface {
	ReadKeys() (map[string]interface{}, error)
	ReadKey(key string) (interface{}, error)
	//RawRead()
	//RawWrite(data []byte)
}
