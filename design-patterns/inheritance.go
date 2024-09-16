package designpatterns

import "fmt"

// in heritance uisng structure embedding

type Bird struct {
	name string
}

func (bird *Bird) Fly() {
	fmt.Printf("%v can fly\n", bird.name)
}

type MallardDuck struct {
	Bird
	sound string
}

func (mduck *MallardDuck) Sound() {
	fmt.Printf("%v sound %v \n", mduck.name, mduck.sound)
}

func NewMallardDuck(name string, sound string) *MallardDuck {
	return &MallardDuck{Bird: Bird{name: name}, sound: sound}
}
