package main

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
)

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

func main() {
	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(b.Age, a.Age)
	})

	next, stop := iter.Pull(slices.Values(users))
	for {
		u, ok := next()
		if !ok {
			break
		}

		if u.Age < 18 {
			stop()
			continue
		}
		fmt.Println(u)
	}
}
