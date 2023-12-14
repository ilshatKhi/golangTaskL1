package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) getName() string {
	return h.name
}

func (h Human) getAge() int {
	return h.age
}

type Action struct {
	Human
}

func (a Action) run() {
	fmt.Println(a.Human.name, "run")
}

func (a Action) walk() {
	fmt.Println(a.Human.name, "walk")
}
