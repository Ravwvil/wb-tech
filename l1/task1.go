package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (h Human) Speak() string {
	return fmt.Sprintf("Hello, my name is %s", h.Name)
}

func (h Human) AgeInYears() int {
	return h.Age
}

type Action struct {
	Human
}

func main() {
	someAction := Action{
		Human{
			Name:   "John",
			Age:    30,
			Gender: "male",
		},
	}
	fmt.Println(someAction.Speak())
	fmt.Println("Age in years:", someAction.AgeInYears())
}
