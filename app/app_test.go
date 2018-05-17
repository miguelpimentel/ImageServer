package main

import "testing"

func TestReturn10(t *testing.T) {

	var number = return10()

	if number != 10 {
		t.Error("Expected 10, got ", number)
	}
}
