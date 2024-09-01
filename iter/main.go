package main

import (
	"fmt"
	"iter"
)

func Backward(mySlice []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := len(mySlice) - 1; i >= 0; i-- {

			if !yield(mySlice[i]) {
				fmt.Println("break")
				return
			}
		}
	}
}

func Backward2[E comparable](mySlice []E, value E) iter.Seq2[string, E] {
	return func(yield func(string, E) bool) {
		for i := len(mySlice) - 1; i >= 0; i-- {

			if !yield(fmt.Sprintf("M%d", i), mySlice[i]) {
				fmt.Println("break")
				return
			}
		}
	}
}

func main() {

	mySlice := []int{1, 2, 3, 4}

	fmt.Println("Backward")
	for s := range Backward(mySlice) {
		fmt.Println(s)
	}

	fmt.Println("Backward2")
	for i, v := range Backward2(mySlice, 5) {
		fmt.Println(i, v)
	}

	fmt.Println("Pull")
	next, stop := iter.Pull(Backward(mySlice))
	for {
		v, ok := next()
		fmt.Println(v, ok)
		if v <= 2 {
			stop()
		}
		if !ok {
			break
		}
		fmt.Println(v)
	}

	fmt.Println("Pull2")
	next2, _ := iter.Pull2(Backward2(mySlice, 3))
	for {
		v1, v2, ok := next2()
		fmt.Println(v1, v2, ok)
		if !ok {
			break
		}
	}

}
