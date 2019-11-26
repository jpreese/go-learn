package main

import "testing"

func TestGetHundreds_ByDefault_ReturnsNumberOfHundreds(t *testing.T) {
	expected := 3
	input := []int{1, 100, 2, 100, 100, 3}

	actual := GetHundreds(input)

	if expected != actual {
		t.Errorf("unexpected sum. expected %v actual %v", expected, actual)
	}
}
