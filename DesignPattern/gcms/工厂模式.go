package main

import "fmt"

/*
简单工厂 优化到复合开不原则
*/

// 水果的抽象接口
type Fruit interface {
	Show()
}

// 工厂抽象接口
type Factory interface {
	CreateFactory() Fruit
}

// 对象
type Apple struct{}

func (a Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct{}

func (b Banana) Show() {
	fmt.Println("我是香蕉")
}

// 每个对象都有一个生产该对象的工厂
type AppleFactory struct{}

func (af AppleFactory) CreateFactory() Fruit {
	a := new(Apple)
	return a
}

type BananaFactory struct{}

func (bf BananaFactory) CreateFactory() Fruit {
	b := new(Banana)
	return b
}

func main() {
	a := new(AppleFactory)
	f := a.CreateFactory()
	f.Show()

}
