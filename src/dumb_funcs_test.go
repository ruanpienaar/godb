package godb

import "testing"

// func TestXxx(*testing.T){}

func TestDoubleNum(t *testing.T) {
	var input int = 123
	output := DoubleNum(input)
	if output != 246 {
		t.Error("Expected 123, got, ", output)
	}
}