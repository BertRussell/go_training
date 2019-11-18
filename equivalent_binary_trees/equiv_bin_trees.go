package main

import "fmt"
import "golang.org/x/tour/tree"

func _walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	_walk(t.Left, ch)
	ch <- t.Value
	_walk(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	chT1 := make(chan int)
	chT2 := make(chan int)
	go Walk(t1, chT1)
	go Walk(t2, chT2)
	for {
		v1, ok1 := <-chT1
		v2, ok2 := <-chT2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}

}

func main() {
	//areSame := Same(tree.New(1), tree.New(2))
	//fmt.Println(areSame)
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
