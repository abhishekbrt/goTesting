package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	CountDown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	t.Run("sleep before every point", func(t *testing.T) {
		spySleepPrinter := SpyCountdownOperations{}
		CountDown(&spySleepPrinter, &spySleepPrinter)
		//fmt.Println(spySleepPrinter.Calls)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

}
