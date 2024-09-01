package main

import (
	"fmt"
)

func Backward(mySlice []int, value int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := len(mySlice) - 1; i >= 0; i-- {

			if value > mySlice[i] {
				continue
			}

			if !yield(mySlice[i]) {
				fmt.Println("break")
				return
			}
		}
	}
}

func Backward2[E comparable](mySlice []E, value E) func(func(string, E) bool) {
	return func(yield func(string, E) bool) {
		for i := len(mySlice) - 1; i >= 0; i-- {

			if !yield(fmt.Sprintf("M%d", i), mySlice[i]) {
				fmt.Println("break")
				return
			}
		}
	}
}

func GetStrings() []string {
	return []string{"a", "b", "c", "d", "e"}
}

func GetInteger() []int {
	return []int{1, 2, 3, 4, 5}
}

func iterate[V any](f func() []V) func(func(V) bool) {
	return func(yield func(V) bool) {
		values := f()

		for _, v := range values {
			if !yield(v) {
				fmt.Println("break")
				return
			}
		}
	}
}

func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Backward")
	for s := range Backward(mySlice, 5) {
		fmt.Println(s)
	}

	fmt.Println("Backward2")
	for i, v := range Backward2(mySlice, 5) {
		fmt.Println(i, v)
	}

	fmt.Println("GetStrings iterate")
	for v := range iterate(GetStrings) {
		fmt.Println(v)
	}

	fmt.Println("GetInteger iterate")
	for v := range iterate(GetInteger) {
		fmt.Println(v)
	}
}
