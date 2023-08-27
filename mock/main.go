package main

import (
	"os"
	"time"
)

func main() {
	sleeper := &ConfigurableSleeper{
		duration: 5 * time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}
