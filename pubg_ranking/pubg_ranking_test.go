package main

import "testing"

type test struct {
	id             int
	currentLeaders []int
	myScore        []int
	expectedOutput []int
}

var tests = []test{
	{
		0,
		[]int{110, 110, 80, 60, 60, 30, 25},
		[]int{3, 25, 50, 80, 90, 110, 120},
		[]int{6, 5, 4, 2, 2, 1, 1}},
	{
		1,
		[]int{110},
		[]int{3},
		[]int{2}},
	{
		2,
		[]int{110},
		[]int{111},
		[]int{1}},
	{
		3,
		[]int{110},
		[]int{110},
		[]int{1}},
	{
		4,
		[]int{},
		[]int{110},
		[]int{1}},
}

func TestLadderRanking(t *testing.T) {
	for _, test := range tests {
		out := ladderRanking(test.currentLeaders, test.myScore)
		if !slicesAreEqual(out, test.expectedOutput) {
			t.Errorf("TestLadderRanking Error on Test id = %d", test.id)
		}
	}
}

func slicesAreEqual(a []int, b []int) bool {
	if !(len(a) == len(b)) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
