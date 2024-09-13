package main

import (
	"fmt"
	designpatterns "go-advanced/design-patterns"
)

/**
* Playgroud
 */
func main() {
	/** strategy design patterns **/
	ducks := []designpatterns.Duck{}
	ducks = append(ducks, *designpatterns.NewDuck("duck-1", &designpatterns.NoFly{}))
	ducks = append(ducks, *designpatterns.NewDuck("duck-2", &designpatterns.NormalFly{}))

	// function is performed dynamically
	for _, duck := range ducks {
		fmt.Println("----------------")
		duck.PrintName()
		duck.PerformFly()
		fmt.Println("----------------")
	}
	/** strategy design patterns **/
}
