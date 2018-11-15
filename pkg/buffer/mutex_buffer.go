package buffer

import (
	"bytes"
	"sync"
)

type MutexBuffer struct {
	buffer bytes.Buffer
	mutex  sync.Mutex
}

func (b *MutexBuffer) Write(p []byte) (n int, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.buffer.Write(p)
}

func (b *MutexBuffer) String() string {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.buffer.String()
}
