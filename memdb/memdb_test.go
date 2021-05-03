package memdb

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

type tmpTest struct {
	key      string
	value    interface{}
	excepted string
}

func TestMemDBExsist(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.Set("test", "testing"); err != nil {
		t.Error(err)
	}
	if !db.Exist("test") {
		t.Error("the key has not been created")
	}
	if db.Exist("notValidKey") {
		t.Error("the key does not actually exist")
	}
}

func TestMemDBInsert(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	test := map[string]tmpTest{
		"test1": {key: "a", value: "string", excepted: "string"},
		"test2": {key: "b", value: 0, excepted: "int"},
		"test3": {key: "c", value: false, excepted: "bool"},
		"test4": {key: "d", value: 100.55, excepted: "float64"},
	}

	for _, value := range test {
		if err := db.Set(value.key, value.value); err != nil {
			t.Error(err)
		}

		result, err := db.Get(value.key)
		if err != nil {
			t.Error(err)
		}

		if result != value.value {
			t.Errorf("value mismatch (%v != %v)", result, value)
		}

		resultType := reflect.TypeOf(result).String()
		if resultType != value.excepted {
			t.Errorf("type mismatch (%v != %v)", resultType, value.excepted)
		}
	}
}

func TestMemDbInsertWithTTL(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.SetWithTTL("test", "testing", 1); err != nil {
		t.Error(err)
	}
	if !db.Exist("test") {
		t.Error("the key has not been created")
	}

	time.Sleep(time.Second * 2)

	if db.Exist("test") {
		t.Error("the key still exists even though the TTL has expired.")
	}
}

func TestMemDBSetTTL(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.Set("test", "testing"); err != nil {
		t.Error(err)
	}
	if !db.Exist("test") {
		t.Error("the key has not been created")
	}

	if err := db.SetTTL("test", 1); err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second * 2)

	if db.Exist("test") {
		t.Error("the key still exists even though the TTL has expired.")
	}
}

func TestMemDBHasTTL(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.Set("test", "testing"); err != nil {
		t.Error(err)
	}
	if !db.Exist("test") {
		t.Error("the key has not been created")
	}
	if db.HasTTL("test") {
		t.Error("The TTL value has not yet been set, but it seems to exist.")
	}
	if err := db.SetTTL("test", 3); err != nil {
		t.Error(err)
	}
	if !db.HasTTL("test") {
		t.Error("The TTL value is set, but it does not seem to exist.")
	}
}

func TestMemDBGetTTL(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.Set("test", "testing"); err != nil {
		t.Error(err)
	}
	if !db.Exist("test") {
		t.Error("the key has not been created")
	}

	now := time.Now().Unix()
	if err := db.SetTTL("test", 3); err != nil {
		t.Error(err)
	}
	if get, err := db.GetTTL("test"); err != nil {
		t.Error(err)
	} else if get != (now + 3) {
		t.Errorf("the set TTL value and the retrieved TTL value differ (%d != %d)", get, now+3)
	}
}

func TestMemDBDeleteTTL(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.Set("test", "testing"); err != nil {
		t.Error(err)
	}
	if !db.Exist("test") {
		t.Error("the key has not been created")
	}

	if err := db.SetTTL("test", 3); err != nil {
		t.Error(err)
	}
	if err := db.DeleteTTL("test"); err != nil {
		t.Error(err)
	}
	if db.HasTTL("test") {
		t.Errorf("TTL has previously been cleared, but there is still a TTL value")
	}
}

func TestMemDBUpdate(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	firstValue := "hello"
	secondValue := "world"

	if err := db.Set("test5", firstValue); err != nil {
		t.Error(err)
	}

	if result, err := db.Get("test5"); err != nil {
		t.Error(err)
	} else if result != firstValue {
		t.Errorf("value mismatch (%v != %v)", result, firstValue)
	}

	if err := db.Set("test5", secondValue); err != nil {
		t.Error(err)
	}

	if result, err := db.Get("test5"); err != nil {
		t.Error(err)
	} else if result != secondValue {
		t.Errorf("value mismatch (%v != %v)", result, secondValue)
	}
}

func TestMemDBDelete(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.Set("test6", "Please delete me!"); err != nil {
		t.Error(err)
	}

	if err := db.Delete("test6"); err != nil {
		t.Error(err)
	}

	if db.Exist("test6") {
		t.Errorf("the value exists even though it has been deleted (%v)", "test6")
	}
}

func TestMemDBClean(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	sum := 0
	for i := 0; i < 1000; i++ {
		if err := db.Set(fmt.Sprintf("test%d", sum), "Test!"); err != nil {
			t.Error(err)
		}
		sum++
	}

	if err := db.Clean(); err != nil {
		t.Error(err)
	}

	sum = 0
	for i := 0; i < 1000; i++ {
		if db.Exist(fmt.Sprintf("test%d", sum)) {
			t.Errorf("not all values have been deleted (%v)", fmt.Sprintf("test%d", sum))
		}
		sum++
	}
}

func TestMemdb_Dump(t *testing.T) {
	datastore := map[string]interface{}{
		"key1": "value1",
		"key2": 1,
		"key3": true,
		"key4": false,
		"key5": 5.5551,
	}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.Dump("TestMemdb_Dump.bin"); err != nil {
		t.Error(err)
	}

	// Corrupted byte: 000
	binaryData := []byte{61, 255, 129, 3, 1, 1, 4, 100, 117, 109, 112, 1, 255, 130, 0, 1, 3, 1, 9, 68, 97, 116, 97, 115, 116, 111, 114, 101, 1, 255, 132, 0, 1, 8, 84, 84, 76, 115, 116, 111, 114, 101, 1, 255, 134, 0, 1, 8, 67, 104, 101, 99, 107, 115, 117, 109, 1, 255, 136, 0, 0, 0, 39, 255, 131, 4, 1, 1, 23, 109, 97, 112, 91, 115, 116, 114, 105, 110, 103, 93, 105, 110, 116, 101, 114, 102, 97, 99, 101, 32, 123, 125, 1, 255, 132, 0, 1, 12, 1, 16, 0, 0, 32, 255, 133, 4, 1, 1, 16, 109, 97, 112, 91, 115, 116, 114, 105, 110, 103, 93, 105, 110, 116, 54, 52, 1, 255, 134, 0, 1, 12, 1, 4, 0, 0, 26, 255, 135, 1, 1, 1, 9, 91, 54, 52, 93, 117, 105, 110, 116, 56, 1, 255, 136, 0, 1, 6, 1, 255, 128, 0, 0, 255, 194, 255, 130, 1, 5, 4, 107, 101, 121, 50, 3, 105, 110, 116, 4, 2, 0, 2, 4, 107, 101, 121, 51, 4, 98, 111, 111, 108, 2, 2, 0, 1, 4, 107, 101, 121, 52, 4, 98, 111, 111, 108, 2, 2, 0, 0, 4, 107, 101, 121, 53, 7, 102, 108, 111, 97, 116, 54, 52, 8, 10, 0, 248, 213, 9, 104, 34, 108, 56, 22, 64, 4, 107, 101, 121, 49, 6, 115, 116, 114, 105, 110, 103, 12, 8, 0, 6, 118, 97, 108, 117, 101, 49, 1, 0, 1, 64, 255, 152, 255, 130, 66, 255, 147, 19, 255, 147, 255, 144, 255, 153, 0, 103, 10, 255, 144, 255, 146, 71, 255, 169, 16, 255, 202, 255, 139, 63, 20, 255, 225, 11, 255, 145, 115, 255, 238, 117, 255, 200, 95, 49, 255, 180, 255, 195, 118, 255, 204, 104, 255, 192, 255, 154, 255, 223, 107, 10, 255, 195, 15, 64, 255, 195, 255, 193, 44, 255, 244, 7, 13, 255, 196, 120, 105, 255, 246, 25, 26, 108, 255, 158, 255, 000, 255, 136, 255, 218, 255, 207, 36, 10, 23, 255, 152, 0}
	if err := ioutil.WriteFile("TestMemdb_Dump_Corrupted.bin", binaryData, 0644); err != nil {
		t.Error(err)
	}
}

func TestMemdb_Load(t *testing.T) {
	datastore := map[string]interface{}{}

	db, err := NewMemDB(&datastore)
	if err != nil {
		t.Error(err)
	}

	if err := db.Load("TestMemdb_Dump.bin"); err != nil {
		t.Error(err)
	}

	tests := []struct {
		key    string
		value  interface{}
		typeof string
	}{
		{key: "key1", value: "value1", typeof: "string"},
		{key: "key2", value: 1, typeof: "int"},
		{key: "key3", value: true, typeof: "bool"},
		{key: "key4", value: false, typeof: "bool"},
		{key: "key5", value: 5.5551, typeof: "float64"},
	}

	for _, test := range tests {
		value, err := db.Get(test.key)
		if err != nil {
			t.Error(err)
		}
		typeof := reflect.TypeOf(value).String()
		if value != test.value {
			t.Errorf("invalid value, given: %s, excepted: %s", value, test.value)
		}
		if typeof != test.typeof {
			t.Errorf("invalid value type, given: %s, excepted: %s", typeof, test.typeof)
		}
	}

	if err := db.Load("TestMemdb_Dump_Corrupted.bin"); err == nil {
		t.Errorf("checksum passed, but data is corrupted: %w", err)
	}

	// Clean up
	if err := os.Remove("TestMemdb_Dump.bin"); err != nil {
		t.Error(err)
	}
	if err := os.Remove("TestMemdb_Dump_Corrupted.bin"); err != nil {
		t.Error(err)
	}
}
