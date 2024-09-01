package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Printf("\n### All ###\n")
	names := []string{"nahuel", "Celeste", "Ulises"}
	for i, v := range slices.All(names) {
		fmt.Println(i, v)
	}

	fmt.Printf("\n### AppendSeq ###\n")
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i += 2 {
			if !yield(i) {
				return
			}
		}
	}
	sAppendSeq := slices.AppendSeq([]int{1, 2}, seq)
	fmt.Println(sAppendSeq)

	fmt.Printf("\n### Backward ###\n")
	for i, v := range slices.Backward(names) {
		fmt.Println(i, v)
	}

	fmt.Printf("\n### Chunk ###\n")

	type Person struct {
		Name string
		Age  int
	}

	var people = []Person{
		{"Nahuel", 34},
		{"Celeste", 33},
		{"Ulises", 42},
		{"Mora", 46},
		{"Paco", 14},
	}

	chuck := slices.Chunk(people, 1)
	for c := range chuck {
		fmt.Println(c)
	}

	fmt.Printf("\n### Colect ###\n")

	v := slices.Collect(chuck)
	fmt.Println(v)

	fmt.Printf("\n### Repeat ###\n")
	numbers := []int{0, 1, 2, 3}

	vRepeat := slices.Repeat(numbers, 1)
	fmt.Println(vRepeat)

	fmt.Printf("\n### Sorted ###\n")

	seq4 := func(yield func(int) bool) {
		for i := 10; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}

	for v := range seq4 {
		fmt.Println(v)
	}

	vSorted := slices.Sorted(seq4)
	fmt.Println(vSorted)

	fmt.Printf("\n### Values ###\n")

	for v := range slices.Values(names) {
		fmt.Println(v)
	}
}
