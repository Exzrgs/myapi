package middlewares

import (
	"sync"
)

var (
	logNo = 1
	mu    sync.Mutex
)

func newTraceID() int {
	mu.Lock()
	no := logNo
	logNo += 1
	mu.Unlock()

	return no
}
