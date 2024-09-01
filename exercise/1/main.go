package main

import "fmt"

type User struct {
	Name string
	Age  int
}

var users = []User{
	{"Gopher", 13},
	{"Diego", 20},
	{"Celeste", 14},
	{"Ulises", 24},
	{"Paco", 15},
}

func iterate(values []User) func(func(User) bool) {
	return func(yield func(User) bool) {

		for _, v := range values {

			if v.Age < 18 {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	for v := range iterate(users) {
		fmt.Println(v)
	}
}
