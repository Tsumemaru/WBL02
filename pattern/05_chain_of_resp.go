package pattern

import "fmt"

/*
Паттерн Chain Of Responsibility относится к поведенческим паттернам уровня объекта.
Паттерн Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса, при этом давая шанс обработать этот запрос нескольким объектам. Получатели связываются в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.
Хорошо, что нет стандартной привязки отправитель-получатель,
Но плохо, то что запрос будет потерян, если не пройдет ни одного обработчика
*/

type request struct {
	name                 string
	concreteHandler1Done bool
	concreteHandler2Done bool
	concreteHandler3Done bool
	concreteHandler4Done bool
}

type mainHandler interface {
	execute(*request)
	setNext(mainHandler)
}
type concreteHandler1 struct {
	next mainHandler
}
type concreteHandler2 struct {
	next mainHandler
}
type concreteHandler3 struct {
	next mainHandler
}
type concreteHandler4 struct {
	next mainHandler
}

func (ch1 *concreteHandler1) execute(r *request) {
	if r.concreteHandler1Done {
		ch1.next.execute(r)
		return
	}
	fmt.Println("Запрос прошел первый обработчик")
	r.concreteHandler1Done = true
	ch1.next.execute(r)
}

func (ch1 *concreteHandler1) setNext(next mainHandler) {
	ch1.next = next
}

func (ch2 *concreteHandler2) execute(r *request) {
	if r.concreteHandler2Done {
		ch2.next.execute(r)
		return
	}
	fmt.Println("Запрос прошел второй обработчик")
	r.concreteHandler2Done = true
	ch2.next.execute(r)
}

func (ch2 *concreteHandler2) setNext(next mainHandler) {
	ch2.next = next
}

func (ch3 *concreteHandler3) execute(r *request) {
	if r.concreteHandler3Done {
		ch3.next.execute(r)
		return
	}
	fmt.Println("Запрос прошел третий обработчик")
	r.concreteHandler3Done = true
	ch3.next.execute(r)
}

func (ch3 *concreteHandler3) setNext(next mainHandler) {
	ch3.next = next
}

func (ch4 *concreteHandler4) execute(r *request) {
	if r.concreteHandler4Done {
		return
	}
	fmt.Println("Запрос прошел четвертый обработчик, финальный")
	r.concreteHandler4Done = true
}

func (ch4 *concreteHandler4) setNext(next mainHandler) {
	ch4.next = next
}

func main() {
	ch4 := &concreteHandler4{}

	ch3 := &concreteHandler3{}
	ch3.setNext(ch4)

	ch2 := &concreteHandler2{}
	ch2.setNext(ch3)

	ch1 := &concreteHandler1{}
	ch1.setNext(ch2)

	request := &request{
		name: "Имя"}
	ch1.execute(request)
}
