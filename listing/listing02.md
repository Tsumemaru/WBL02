
```package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
test() выведет 2, anotherTest() выведет 1
Перед завершением отработывают defer'ы, если их несколько, то первым будет тот, который в коде написан первым.
А дальше последующие.
Поэтому в первой функции значение x = 1, но перед завершение в defer инкрементируем x и x = 2
(Так как вывод конкретной переменной в функции определен, то анонимная функция может вернуть результат инкрементации)
Во второй функции есть просто вывод, но анонимная горутина не возвращает результат