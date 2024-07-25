package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("countdown", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)
		got := buffer.String()
		want := "3\n2\n1\ngo!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if sleeper.Calls != 3 {
			t.Errorf("not enough calls to sleepper, wanted 3, got %q", sleeper.Calls)
		}
	})
}
