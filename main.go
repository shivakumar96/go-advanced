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

	/** Genrics **/
	// all possible maps
	gmap1 := designpatterns.NewGenericMap[int, string](2)
	gmap2 := designpatterns.NewGenericMap[int, int](2)
	gmap3 := designpatterns.NewGenericMap[string, string](2)
	gmap4 := designpatterns.NewGenericMap[string, int](2)

	gmap1[1] = "s"
	fmt.Println(gmap1)

	gmap2[2] = 2
	fmt.Println(gmap2)

	gmap3["a"] = "a"
	fmt.Println(gmap3)

	gmap4["str"] = 123
	fmt.Println(gmap4)

	/** Genrics **/
}
