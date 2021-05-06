package memdb

import (
	"bytes"
	"crypto/sha512"
	"encoding/gob"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type MemDBStore map[string]interface{}

type MemDB interface {
	// Get a value from memory by key.
	Get(key string) (interface{}, error)

	// GetAll values from memory
	GetAll() (MemDBStore, error)

	// Set a value in memory based on the specified key
	Set(key string, value interface{}) error

	// SetWithTTL is same as Set but with TTL you can specify in seconds when to delete this key.
	SetWithTTL(key string, value interface{}, ttl int64) error

	// SetTTL to specified time to live value
	SetTTL(key string, ttl int64) error

	// GetTTL time to live value
	GetTTL(key string) (int64, error)

	// DeleteTTL existing time to live value
	DeleteTTL(key string) error

	// HasTTL value?
	HasTTL(key string) bool

	// Delete a value from memory based on the specified key
	Delete(key string) error

	// Clean all values from memory
	Clean() error

	// Exist -> Does the specified key exist?
	Exist(key string) bool

	// Dump values to file
	Dump(filename string) error

	// Load values from file
	Load(filename string) error
}

func NewMemDB(dataVariable *MemDBStore) (MemDB, error) {
	ttlstore := map[string]int64{}

	go func() {
		for now := range time.NewTicker(time.Second).C {
			for key, expire := range ttlstore {
				if now.Unix() >= expire {
					log.Debugf("%s was deleted (expired ttl)", key)
					delete(*dataVariable, key)
					delete(ttlstore, key)
				}
			}
		}
	}()

	return &memdb{
		datastore: *dataVariable,
		ttlstore:  ttlstore,
	}, nil
}

type memdb struct {
	datastore MemDBStore
	ttlstore  map[string]int64
}

type dump struct {
	Datastore MemDBStore
	TTLstore  map[string]int64
	Checksum  [64]byte
}

// Get retrieves a value from memory by key.
func (m *memdb) Get(key string) (interface{}, error) {
	if m.Exist(key) {
		return m.datastore[key], nil
	}
	return nil, fmt.Errorf("%s key is doesn't exsist", key)
}

func (m *memdb) GetAll() (MemDBStore, error) {
	return m.datastore, nil
}

// Set write a value into memory based on a key
func (m *memdb) Set(key string, value interface{}) error {
	m.datastore[key] = value
	return nil
}

// SetWithTTL value
func (m *memdb) SetWithTTL(key string, value interface{}, ttl int64) error {
	if err := m.Set(key, value); err != nil {
		return err
	}
	return m.SetTTL(key, ttl)
}

// SetTTL value
func (m *memdb) SetTTL(key string, ttl int64) error {
	if !m.Exist(key) {
		return fmt.Errorf("%s key is doesn't exsist", key)
	}
	m.ttlstore[key] = time.Now().Unix() + ttl
	return nil
}

// GetTTL value
func (m *memdb) GetTTL(key string) (int64, error) {
	if !m.Exist(key) {
		return 0, fmt.Errorf("%s key is doesn't exsist", key)
	}
	if !m.HasTTL(key) {
		return 0, fmt.Errorf("%s the key has no TTL value or has been deleted", key)
	}
	return m.ttlstore[key], nil
}

// DeleteTTL value
func (m *memdb) DeleteTTL(key string) error {
	if !m.HasTTL(key) {
		return fmt.Errorf("%s the key has no TTL value or has been deleted", key)
	}
	delete(m.ttlstore, key)
	return nil
}

// HasTTL value
func (m *memdb) HasTTL(key string) bool {
	if _, ok := m.ttlstore[key]; ok {
		return true
	}
	return false
}

// Exist checks if such a key exists in memory.
func (m *memdb) Exist(key string) bool {
	if _, ok := m.datastore[key]; ok {
		return true
	}
	return false
}

// Delete the value by key.
func (m *memdb) Delete(key string) error {
	if !m.Exist(key) {
		return fmt.Errorf("%s key is doesn't exsist", key)
	}
	delete(m.datastore, key)
	if m.HasTTL(key) {
		if err := m.DeleteTTL(key); err != nil {
			return err
		}
	}
	return nil
}

// Clean memory
func (m *memdb) Clean() error {
	for key := range m.datastore {
		if err := m.Delete(key); err != nil {
			return err
		}
	}
	return nil
}

// Dump memory database to file
func (m *memdb) Dump(filename string) error {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)

	d := &dump{}
	d.TTLstore = m.ttlstore
	d.Datastore = m.datastore
	d.Checksum = [64]byte{}

	checksum := sha512.Sum512([]byte(fmt.Sprintf("%v", d)))
	d.Checksum = checksum

	if err := encoder.Encode(d); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, w.Bytes(), 0600); err != nil {
		return err
	}

	return nil
}

// Load data from memory database
func (m *memdb) Load(filename string) error {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	r := bytes.NewBuffer(f)
	decoder := gob.NewDecoder(r)

	d := &dump{}
	if err := decoder.Decode(&d); err != nil {
		return err
	}

	exportedChecksum := d.Checksum

	// Delete the Checksum value extracted from the file to make the calculation correct.
	d.Checksum = [64]byte{}

	calculatedChecksum := sha512.Sum512([]byte(fmt.Sprintf("%v", d)))

	if calculatedChecksum != exportedChecksum {
		return fmt.Errorf("invalid checksum: %s", filename)
	}

	m.ttlstore = d.TTLstore
	m.datastore = d.Datastore

	return nil
}
