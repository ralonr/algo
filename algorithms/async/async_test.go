package async

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestDispatcher(t *testing.T) {
	tests := []struct {
		name       string
		dispatcher Dispatcher
		isRunning  bool
		number     int
	}{
		{
			name:       "test Forget",
			dispatcher: Forget(),
			isRunning:  true,
			number:     0,
		},
		{
			name:       "test Wait",
			dispatcher: Wait(),
			isRunning:  false,
			number:     10,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			running := false
			i := 0
			var m sync.RWMutex
			test.dispatcher.Dispatch(func() {
				m.Lock()
				running = true
				m.Unlock()
				time.Sleep(1 * time.Second)
				m.Lock()
				i = 10
				running = false
				m.Unlock()
			})
			time.Sleep(10 * time.Millisecond)
			m.RLock()
			assert.True(t, running == test.isRunning)
			assert.Equal(t, test.number, i)
			m.RUnlock()
		})
	}
}
