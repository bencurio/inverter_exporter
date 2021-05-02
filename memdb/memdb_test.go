package memdb

import (
	"fmt"
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
