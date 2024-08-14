package main

import "fmt"

/*
有一个抽象接口

实现抽象接口的各个实例

一个工厂对象.包含内部调用的方法

再函数汇中实现工厂对象 根据不同需求来调用给工厂对象传入不同的实例

*/

type Fruit interface {
	Show()
}

type Apple struct{}

func (a Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct{}

func (b Banana) Show() {
	fmt.Println("我是香蕉")
}

type Factory struct {
}

// 将复杂的业务逻辑放在工厂这个方法中
func (f Factory) CreateFactory(name string) Fruit {
	var fruit Fruit
	if name == "apple" {
		fruit = new(Apple)
	}
	if name == "banana" {
		fruit = new(Banana)
	}
	return fruit
}
func main() {
	factory := new(Factory)
	f := factory.CreateFactory("apple")
	f.Show()
}
