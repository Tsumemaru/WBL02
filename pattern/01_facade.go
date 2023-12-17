package pattern

import "fmt"

/*
Паттерн Facade относится к структурным паттернам уровня объекта.
Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору взаимосвязанных классов или объектов некоторой подсистемы, что облегчает ее использование.
Будет ли его реализация больше плюсом или минусом зависит от того, нужно ли пользователю тестирование, дополнение, изменинение в целом этого интерфейса
*/

type Facade struct {
	first  *first
	second *second
	third  *third
}

type first struct {
	msg string
}
type second struct {
	msg string
}
type third struct {
	msg string
}

func newFacade() *Facade {
	facade := &Facade{
		first:  newFirst("Первый сегмент имеется"),
		second: newSecond("Второй сегмент имеется"),
		third:  newThird("Третий сегмент имеется"),
	}
	fmt.Println("Создание фасада")
	return facade
}

func newFirst(msg string) *first {
	return &first{
		msg: msg,
	}
}

func (f *Facade) checkFirst() {
	fmt.Print(f.first.msg)
}

func newSecond(msg string) *second {
	return &second{
		msg: msg,
	}
}
func (f *Facade) checkSecond() {
	fmt.Print(f.second.msg)
}
func newThird(msg string) *third {
	return &third{
		msg: msg,
	}
}
func (f *Facade) checkThird() {
	fmt.Print(f.third.msg)
}

func main() {
	facade := newFacade()
	facade.checkSecond()
}
