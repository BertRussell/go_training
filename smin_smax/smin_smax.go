package main

import "fmt"

func main() {
	var smin, smax int
	a := []int{5, 1, 4, 1, 4, 54, 22, 53}

	smin, smax = smin_smax_sorted(a)
	fmt.Printf("smin:  %v smax: %v \n", smin, smax)
	smin, smax = smin_smax_nsorted(a)
	fmt.Printf("smin:  %v smax: %v \n", smin, smax)
}

func smin_smax_sorted(a []int) (smin int, smax int) {
	if len(a) >= 1 {
		smin = sum(a[:len(a)-1])
		smax = sum(a[1:])
	}
	return
}

func smin_smax_nsorted(a []int) (smin int, smax int) {
	var max_v, min_v int
	var s int = sum(a)
	if len(a) >= 1 {
		_, max_v = max(a)
		_, min_v = min(a)
		smin = s - max_v
		smax = s - min_v
	}
	return
}

func max(a []int) (max_idx int, max_value int) {
	if len(a) >= 1 {
		max_value = a[0]
		max_idx = 0
		for i, v := range a {
			if v > max_value {
				max_value = v
				max_idx = i
			}
		}
	}
	return
}

func min(a []int) (min_idx int, min_value int) {
	if len(a) >= 1 {
		min_value = a[0]
		min_idx = 0
		for i, v := range a {
			if v < min_value {
				min_value = v
				min_idx = i
			}
		}
	}
	return
}

func sum(a []int) (sum int) {
	sum = 0
	for _, v := range a {
		sum += v
	}
	return
}
