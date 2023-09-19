package clockface_test

import (
	"testing"
	"time"

	"github.com/wuziqiu666/learngo/math/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestSecondHandAt30Second(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
