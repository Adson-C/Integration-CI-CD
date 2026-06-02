package main

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(15, 5)
	expected := 20

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}