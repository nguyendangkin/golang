package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "chin", age: 22})

	fmt.Println(newPerson("kin"))

	s := person{name: "sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{"Rex", true}
	fmt.Println(dog)

	dog2 := person{name: "chin", age: 22}

	fmt.Println(dog2)

}
