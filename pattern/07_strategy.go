package pattern

import "fmt"

/*
Паттерн Strategy относится к поведенческим паттернам уровня объекта.
Паттерн Strategy определяет набор алгоритмов схожих по роду деятельности, инкапсулирует их в отдельный класс и делает их подменяемыми.
Паттерн Strategy позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.
Из-за большого количества структур усложняет код и пользоваться можно только при понимании подходящей статегии
*/

type Eviction interface {
	strategyResult(obj *Object)
}

type Strategy1 struct{}
type Strategy2 struct{}
type Strategy3 struct{}

type Object struct {
	Info     string
	Strategy Eviction
}

func (obj *Object) initStrategy(info string, strategy Eviction) {
	obj.Info = info
	obj.Strategy = strategy
}

func (obj *Object) changeStrategy(strategy Eviction) {
	obj.Strategy = strategy
}

func (str *Strategy1) strategyResult(obj *Object) {
	fmt.Println(obj.Info)
}

func (str *Strategy2) strategyResult(obj *Object) {
	fmt.Println(obj.Info + obj.Info)
}

func (str *Strategy3) strategyResult(obj *Object) {
	fmt.Println(obj.Info + obj.Info + obj.Info)
}

func (obj *Object) showInfo() {
	obj.Strategy.strategyResult(obj)
}

func main() {
	obj := &Object{}
	str1 := &Strategy1{}
	str2 := &Strategy2{}
	str3 := &Strategy3{}
	obj.initStrategy("something", str1)
	obj.showInfo()
	obj.changeStrategy(str2)
	obj.showInfo()
	obj.changeStrategy(str3)
	obj.showInfo()
}
