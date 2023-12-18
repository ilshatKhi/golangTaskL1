package main

import "fmt"

type Human struct { // создали структуру
	name string
	age  int
}

func (h Human) getName() string { // геттер структуры Human возвращает имя
	return h.name
}

func (h Human) getAge() int { // геттер структуры Human возвращает возраст
	return h.age
}

type Action struct { // создаем структуру Action
	Human // встраивавем структуру Human
}

func (a Action) run() { // создали метод run структуры
	fmt.Println(a.Human.name, "run")
}

func (a Action) walk() {
	fmt.Println(a.Human.name, "walk")
}

/*func main() {
	hAction := Action{Human{ // создали экземпляр Action
		name: "Ivan",
		age:  25,
	}}

	hAction.walk() // вызываем методы структуры Action
	hAction.run()
}
*/
