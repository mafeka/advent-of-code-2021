package main

import (
	"reflect"
	"testing"
)

func TestGetList(t *testing.T) {
	got := getList("Ich\nbin\nder\nTest")
	want := []string{"Ich", "bin", "der", "Test"}

	result := reflect.DeepEqual(got, want)

	if result != true {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestConvertToInt(t *testing.T) {
	got := convertStringToInt("999")
	want := 999

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIsHigher(t *testing.T) {
	result_1 := isHigher(10, 100)
	if result_1 != 1 {
		t.Errorf("Error: %q should be 1", result_1)
	}

	result_0 := isHigher(100, 10)
	if result_0 != 0 {
		t.Errorf("Error: %q should be 0", result_0)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{10, 10, 10}
	want := 30
	got := addElements(input)

	if got != want {
		t.Errorf("Was expecting %q but got %q", want, got)
	}
}
