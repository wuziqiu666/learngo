package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Printer 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOpreation{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOpreation{}
		Countdown(spySleepPrinter, spySleepPrinter)
		got := spySleepPrinter.Calls
		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want calls %v got %v", want, got)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spyTime.Sleep,
	}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have seplt for %v but seplt for %v", sleepTime, spyTime.durationSlept)
	}
}

type SpyCountdownOpreation struct {
	Calls []string
}

func (s *SpyCountdownOpreation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOpreation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
