package rand

import (
	"reflect"
	"testing"
)

func TestRandInt(t *testing.T) {
	actual := make([]int, 5)
	noExpected := make([]int, 5)
	GenerateInt(actual)

	if reflect.DeepEqual(actual, noExpected) {
		t.Errorf("Could not generate random integer array")
	}
}

func TestRandString(t *testing.T) {
	actual := make([]string, 5)
	noExpected := make([]string, 5)
	GenerateString(actual)

	if reflect.DeepEqual(actual, noExpected) {
		t.Errorf("Could not generate random character string")
	}
}
