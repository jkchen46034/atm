package handler

import (
	"testing"
	"time"
)

func Test_Timer(t *testing.T) {

	timer := NewTimer()
	timer.Reset()
	time.Sleep(100 * time.Millisecond)
	timer.Close()
}
