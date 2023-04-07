package store

import (
	"fmt"
	"testing"
)

func TestSetValue(t *testing.T) {

	store := NewCaptchaBucket()

	for i := 0; i < 5; i++ {
		i := i
		go func() {
			err := store.Set(fmt.Sprintf("item-%v", i), "1")
			if err != nil {
				t.Error(err)
			}
		}()
	}

}

func TestGet(t *testing.T) {
	store := NewCaptchaBucket()
	err := store.Set("id-1", "1")
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 5; i++ {
		go func() {
			value := store.Get("id-1", false)
			if value != "1" {
				t.Errorf("value is : %s", value)
			}
		}()
	}

}

func TestCaptchaBucket_Delete(t *testing.T) {
	store := NewCaptchaBucket()
	err := store.Set("id-1", "1")
	err = store.Set("id-2", "2")
	if err != nil {
		t.Error(err)
	}

	value := store.Get("id-1", false)
	if value != "1" {
		t.Errorf("value is : %s", value)
	}

	store.Delete("id-2")

	ok := store.Verify("id-2", "2", false)

	if ok != false {
		t.Errorf("Failure validation failed - %v", ok)
	}

}
