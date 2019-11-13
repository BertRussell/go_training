package main

import (
	"fmt"
	"sort"
)

func main() {
	var z, z2, z3 int
	a := []int{8, 7, 23, 9, 2, 0}
	b := []int{0, 23, 8, 9, 2}
	z = FindMissingElement(a, b)
	z2 = FindMissingElement2(b, a)
	z3 = FindMissingElement3(b, a)
	fmt.Printf("z = %d \n", z)
	fmt.Printf("z2 = %d \n", z2)
	fmt.Printf("z3 = %d \n", z3)
}

// O(nlogn + nlogn + n) = O(nlogn)
func FindMissingElement(a []int, b []int) (x int) {
	sort.Ints(a)
	sort.Ints(b)

	var long, short []int

	if len(a) > len(b) {
		long = a
		short = b
	} else {
		long = b
		short = a
	}

	for i, v := range long {
		x = v
		if i < len(short) && long[i] != short[i] {
			break
		}
	}

	return
}

//O(n^2)
func FindMissingElement2(a []int, b []int) (x int) {

	var long, short []int

	if len(a) > len(b) {
		long = a
		short = b
	} else {
		long = b
		short = a
	}
	var found bool

	for _, vl := range long {
		found = false
		x = vl
		for j, vs := range short {
			if vl == vs {
				found = true
				continue
			}
			if !found && j == len(short)-1 {
				return
			}
		}
	}
	return
}

//O(n^2)
func FindMissingElement3(a []int, b []int) (x int) {
	for _, va := range a {
		for j, vb := range b {
			if va == vb {
				b[j] = b[len(b)-1]
				b = b[:len(b)-1]
				continue
			}
		}
	}

	if len(b) > 0 {
		x = b[0]
	} else {
		x = a[len(a)-1]
	}

	return
}
