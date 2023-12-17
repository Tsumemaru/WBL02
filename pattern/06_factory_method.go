package pattern

import "fmt"

/*
Паттерн Factory Method относится к порождающим паттернам уровня класса и сфокусирован только на отношениях между классами.
Паттерн Factory Method полезен, когда система должна оставаться легко расширяемой путем добавления объектов новых типов.
Этот паттерн является основой для всех порождающих паттернов и может легко трансформироваться под нужды системы.
По этому, если перед разработчиком стоят не четкие требования для продукта или не ясен способ организации взаимодействия между продуктами, то для начала можно воспользоваться паттерном Factory Method, пока полностью не сформируются все требования.
*/

type Factory interface {
	setName(name string)
	setInfo(info string)
	getName() string
	getInfo() string
}

type Object struct {
	name string
	info string
}

type ConcreteObject1 struct {
	Object
}

type ConcreteObject2 struct {
	Object
}

func (obj *Object) setName(name string) {
	obj.name = name
}

func (obj *Object) setInfo(info string) {
	obj.info = info
}

func (obj *Object) getName() string {
	return obj.name
}

func (obj *Object) getInfo() string {
	return obj.info
}

func newConcreteObject1() Factory {
	return &ConcreteObject1{
		Object: Object{
			name: "Первый обьект",
			info: "Cоздали первый обьект",
		},
	}
}

func newConcreteObject2() Factory {
	return &ConcreteObject2{
		Object: Object{
			name: "Второй обьект",
			info: "Cоздали второй обьект",
		},
	}
}

func getObj(number int) Factory {
	if number == 1 {
		return newConcreteObject1()
	}
	if number == 2 {
		return newConcreteObject2()
	}
	return nil
}

func main() {
	obj1 := getObj(1)
	obj2 := getObj(2)
	fmt.Println(obj1.getName())
	obj1.setInfo("Обновили информацию о первом")
	fmt.Println(obj1.getInfo())
	fmt.Println(obj2.getInfo())
	fmt.Println(obj2.getName())
}
