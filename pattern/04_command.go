package pattern

import (
	"fmt"
	"time"
)

/*
Паттерн Command относится к поведенческим паттернам уровня объекта.
Паттерн Command позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект. Такие запросы, например, можно ставить в очередь, отменять или возобновлять.
Этот паттерн удобен для оборачивания операций в отдельные обьекты.Отсутствие прямой зависимости, засчет усложнения кода программы
*/

type Button struct {
	command Command
}

type Device interface {
	On()
	Off()
}

type PC struct {
	status bool
}

type On struct {
	device Device
}

type Off struct {
	device Device
}

type Command interface {
	Execute()
}

func (pc *PC) On() {
	pc.status = true
	fmt.Println("ПК включен")
}

func (pc *PC) Off() {
	pc.status = false
	fmt.Println("ПК выключен")
}

func (b *Button) Exec() {
	b.command.Execute()
}

func (on *On) Execute() {
	on.device.On()
}

func (off *Off) Execute() {
	off.device.Off()
}

func main() {
	pc := &PC{}
	on := &On{
		device: pc,
	}
	off := &Off{
		device: pc,
	}
	btnOn := &Button{
		command: on,
	}
	btnOff := &Button{
		command: off,
	}
	btnOn.Exec()
	time.Sleep(1 * time.Second)
	btnOff.Exec()
}
