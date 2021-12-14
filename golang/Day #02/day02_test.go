package main

import (
	"testing"
)

func TestGetAxisAndDirection(t *testing.T) {
	got := getAxisAndDirection("up 5")
	want := Position{x: 0, y: -5}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = getAxisAndDirection("down 6")
	want = Position{x: 0, y: 6}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = getAxisAndDirection("forward 10")
	want = Position{x: 10, y: 0}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func TestGetMultiplier(t *testing.T) {

	got := getMultiplier("up")
	want := Position{x: 0, y: -1}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = getMultiplier("down")
	want = Position{x: 0, y: 1}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got = getMultiplier("forward")
	want = Position{x: 1, y: 0}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
