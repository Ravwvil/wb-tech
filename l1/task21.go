package main

import (
	"fmt"
)

type SomeInterface interface {
	SomeMethod() string
}

type SomeInterfaceOld interface {
	SomeOldMethod() int
}

type ImplementationOld struct{}

func (old *ImplementationOld) SomeOldMethod() int {
	return 52
}

type Adapter struct {
	Adaptee SomeInterfaceOld
}

func (a *Adapter) SomeMethod() string {
	value := a.Adaptee.SomeOldMethod()
	return fmt.Sprintf("Adapted value: %d", value)
}

func main() {
	old := &ImplementationOld{}
	var new SomeInterface = &Adapter{Adaptee: old}
	fmt.Println(new.SomeMethod())
	/*
		Паттерн "Адаптер" позволяет объектам с несовместимыми интерфейсами работать вместе
		без изменения исходного коде

		Плюс паттерна "Адаптер":
		- Позволяет использовать существующие классы с несовместимыми интерфейсами
		- Упрощает интеграцию старых и новых систем
		Минус паттерна "Адаптер":
		- Может привести к усложнению структуры кода, если используется слишком много адаптеров

		Реальные примеры использования:
		Адаптер часто применяется при работе с HTTP (например, для адаптации хендлеров и роутеров),
		а также при интеграции с внешними API, где нужно преобразовывать данные из одного формата в другой.
	*/
}
