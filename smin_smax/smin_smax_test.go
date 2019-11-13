package main

import "testing"

type test struct {
	input         []int
	expected_smin int
	expected_smax int
}

func TestSmin_smax(t *testing.T) {
	var tests []test
	var smin, smax int

	tests = append(tests, test{input: []int{5, 1, 4, 1, 4, 54, 22, 53}, expected_smax: 143, expected_smin: 90})
	tests = append(tests, test{input: []int{}, expected_smax: 0, expected_smin: 0})
	tests = append(tests, test{input: []int{-1}, expected_smax: 0, expected_smin: 0})
	tests = append(tests, test{input: []int{-1, -2}, expected_smax: -1, expected_smin: -2})
	tests = append(tests, test{input: []int{-1, -2, 3}, expected_smax: 2, expected_smin: -3})

	for _, test := range tests {
		smin, smax = smin_smax_nsorted(test.input)
		if smin != test.expected_smin || smax != test.expected_smax {
			t.Errorf("smin_smax_nsorted() Error. Expected smin: %d Got: %d Expected smax: %d Got: %d", test.expected_smin, smin, test.expected_smax, smax)
		}
	}
}
