package designpatterns

import "fmt"

// in heritance uisng structure embedding

type Bird struct {
	name string
}

func (bird *Bird) Fly() {
	fmt.Printf("%v can fly\n", bird.name)
}

func (bird *Bird) Sound() {
	fmt.Printf("Sound undefined\n")
}

type MallardDuck struct {
	Bird
	sound string
}

// method overding
func (mduck *MallardDuck) Sound() {
	fmt.Printf("%v sound %v \n", mduck.name, mduck.sound)
}

func NewMallardDuck(name string, sound string) *MallardDuck {
	return &MallardDuck{Bird: Bird{name: name}, sound: sound}
}
