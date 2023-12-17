package pattern

import "fmt"

/*
Паттерн State относится к поведенческим паттернам уровня объекта.
Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния и является объектно-ориентированной реализацией конечного автомата.
Поведение объекта изменяется настолько, что создается впечатление, будто изменился класс объекта.
Концентрирует код в одном месте, но если мало состояний или редко меняются, то может быть паттерн излишним
*/

type State interface {
	checkState()
}

type Object struct {
	ObjectState1 State
	ObjectState2 State
	ObjectState3 State
	CurrentState State
}
type State1 struct{}
type State2 struct{}
type State3 struct{}

func (obj *Object) checkState() {
	switch obj.CurrentState {
	case obj.ObjectState1:
		obj.ObjectState1.checkState()
		obj.setState(obj.ObjectState2)
	case obj.ObjectState2:
		obj.ObjectState2.checkState()
		obj.setState(obj.ObjectState3)
	case obj.ObjectState3:
		obj.ObjectState3.checkState()
		obj.setState(obj.ObjectState1)
	}
}

func NewObject() *Object {
	return &Object{
		ObjectState1: &State1{},
		ObjectState2: &State2{},
		ObjectState3: &State3{},
	}
}

func (obj *Object) setState(s State) {
	obj.CurrentState = s
}

func (st *State1) checkState() {
	fmt.Println("Обьект в состоянии 1")
}
func (st *State2) checkState() {
	fmt.Println("Обьект в состоянии 2")
}
func (st *State3) checkState() {
	fmt.Println("Обьект в состоянии 3")
}

func main() {
	obj := NewObject()
	obj.setState(obj.ObjectState1)
	for i := 0; i < 7; i++ {
		obj.checkState()
	}
}
