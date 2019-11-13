package main

import "testing"

type test struct {
	id          int
	inputA      []int
	inputB      []int
	expectedOut int
}

var tests = []test{
	{
		inputA:      []int{8, 7, 23, 9, 2, 0},
		inputB:      []int{0, 23, 8, 9, 2},
		expectedOut: 7}}

func TestSmin_smax(t *testing.T) {
	//var tests []test
	var out int

	for _, test := range tests {
		out = FindMissingElement(test.inputA, test.inputB)
		if out != test.expectedOut {
			t.Errorf("FindMissingElement() Error on Test id = %d", test.id)
		}
	}
}
