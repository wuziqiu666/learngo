package clockface_test

import (
	"testing"
	"time"
)

func TestSecondInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}

}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func 