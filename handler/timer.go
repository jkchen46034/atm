package handler

import (
	"fmt"
	"time"

	"takeoff.com/atm/model"
)

type Timer struct {
	ticker   *time.Ticker
	commands chan string
}

func NewTimer() *Timer {
	timer := &Timer{}

	timer.ticker = time.NewTicker(time.Duration(600 * time.Second))
	timer.commands = make(chan string)

	go timer.GoRoutine()

	return timer
}

func (t *Timer) GoRoutine() {
	for {
		select {
		case ti := <-t.ticker.C:
			fmt.Println("timer fired", ti)
			model.Log.Logout()
		case command := <-t.commands:
			t.exec(command)
			if command == "CLOSE_TIMER" {
				return
			}
		}
	}
}

func (t *Timer) Reset() {
	t.commands <- "RESET_TIMER"
}

func (t *Timer) Stop() {
	t.commands <- "STOP_TIMER"
}

func (t *Timer) Close() {
	t.commands <- "CLOSE_TIMER"
}

func (t *Timer) exec(command string) {
	switch command {
	case "RESET_TIMER":
		t.ticker.Stop()
		t.ticker = time.NewTicker(time.Duration(10 * time.Second))
		return
	case "STOP_TIMER":
		t.ticker.Stop()
		return
	case "CLOSE_TIMER":
		t.ticker.Stop()
		return
	}
}
