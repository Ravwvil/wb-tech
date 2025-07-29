package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// justString = v[:100]
	// Проблема - срез строки не создает копию, а просто ссылается на исходную строку
	// Чтобы избежать проблем с памятью, нужно создать копию строки

	justString = string([]byte(v[:100])) // Корректный вариант, здесь мы создаем копию первых 100 байт строки
}

func main() {
	someFunc()
}
