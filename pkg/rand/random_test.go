package rand_test

import (
	"reflect"
	"testing"

	"github.com/theoden9014/exercitu/pkg/rand"
)

func TestRandInt(t *testing.T) {
	actual := make([]int, 5)
	noExpected := make([]int, 5)
	rand.GenerateInt(actual)

	if reflect.DeepEqual(actual, noExpected) {
		t.Errorf("Could not generate random integer array")
	}
}

func TestRandString(t *testing.T) {
	actual := make([]string, 5)
	noExpected := make([]string, 5)
	rand.GenerateString(actual)

	if reflect.DeepEqual(actual, noExpected) {
		t.Errorf("Could not generate random character string")
	}
}
