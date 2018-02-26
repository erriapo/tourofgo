package main

import "golang.org/x/tour/tree"
import "fmt"
import "time"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		Walk(t1, c1)
	}()

	go func() {
		Walk(t2, c2)
	}()

	var v1 int = 0
	var v2 int = 0
	var ex1 bool = false
	var ex2 bool
	for {

		select {
		case v1 = <-c1:
			fmt.Println(v1)
			ex1 = false
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 1")
			ex1 = true
		}

		select {
		case v2 = <-c2:
			fmt.Println(v2)
			ex2 = false
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 2")
			ex2 = true
		}

		if !ex1 && !ex2 {
			if v1 != v2 {
				return false
			}
		}

		if ex1 || ex2 {
			break
		}
	}
	return true
}

func main() {
	var r1 bool = Same(tree.New(1), tree.New(1))
	fmt.Println(r1)

	var r2 bool = Same(tree.New(1), tree.New(2))
	fmt.Println(r2)
}
