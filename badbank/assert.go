package main

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
