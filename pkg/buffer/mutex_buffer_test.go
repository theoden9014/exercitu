package utils

import (
	"sync"
	"testing"
)

func TestMutexBuffer(t *testing.T) {
	var buf MutexBuffer
	wg := &sync.WaitGroup{}
	parallels := 10
	writeString := "exercitu"
	writeCount := 100

	for i := 0; i < parallels; i++ {
		wg.Add(1)

		go func() {
			for n := 0; n < writeCount; n++ {
				buf.Write([]byte(writeString))
			}
			wg.Done()
		}()
	}

	wg.Wait()

	actual := len(buf.String())
	expected := parallels * writeCount * len(writeString)
	if actual != expected {
		t.Errorf("Failed to write due to data conflict")
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
}
