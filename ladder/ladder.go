package main

import "fmt"

func main() {
	var a = 10
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			if j+i+1 > a {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
