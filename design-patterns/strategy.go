package designpatterns

import "fmt"

/**
* Strategy design pattern
 */

type IFly interface {
	Fly()
}

type Duck struct {
	name string // privte field accessed through the
	fly  IFly
}

func NewDuck(name string, fly IFly) *Duck {
	duck := &Duck{name: name, fly: fly}
	return duck
}

func (d *Duck) PrintName() {
	fmt.Println(d.name)
}

func (d *Duck) PerformFly() {
	d.fly.Fly()
}

type NormalFly struct{}
type NoFly struct{}

func (f *NoFly) Fly() {
	fmt.Println("Soory, I cannot fly")
}

func (f *NormalFly) Fly() {
	fmt.Println("Hey, I can fly")
}

// usage
// duck1 := NewDuck("duck-1",Nofly{})
// duck1 := NewDuck("duck-2",Normalfly{})
// duck1.PrintName()
// duck1.PerformFly()
// -- change of fuction depending on the behaviour --
// duck1.PrintName()
// duck1.PerformFly()
