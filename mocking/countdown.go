package mocking

import (
	"fmt"
	"io"
	"time"
	//"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// struct for spying write and sleep operation
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"
const start = 3
const lastword = "Go!"

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := start; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
		//sleeper.Sleep()
	}
	fmt.Fprint(writer, lastword)
}
