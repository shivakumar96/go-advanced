package main

import (
	"fmt"
	designpatterns "go-advanced/design-patterns"
	goroutines "go-advanced/explore-concurrency"
	customcontainer "go-advanced/explore-containers"
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
	fmt.Println("----------------")
	/** Genrics End **/

	/** Go routines **/
	//go routine
	goroutines.Mainroutine()

	//Anonymous function
	goroutines.Anonymousroutine()

	/** Go routines End **/
	fmt.Println("----------------")
	/** channels **/
	goroutines.PassNumberUsingChannel()

	goroutines.PassNumberUsingBufferedChannel()

	fmt.Println("----------------")

	cache := customcontainer.NewLRUCache[int, int](5)
	cache.Put(2, 3)
	cache.Put(3, 4)
	cache.PrintList()
	cache.Put(3, 5)
	cache.PrintList()
	cache.Get(2)
	cache.PrintList()
	cache.Put(6, 9)
	cache.Put(7, 8)
	cache.Put(8, 16)
	cache.PrintList()
	cache.Put(9, 12)
	cache.PrintList()

	/**
	* Inheritance using struct embedding
	 */

	fmt.Println("----------------")

	mduck := designpatterns.NewMallardDuck("duck-1", "quack quack")
	mduck.Fly()
	mduck.Sound()

	fmt.Println("----------------")

}
