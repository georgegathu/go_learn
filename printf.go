package main

import "fmt"

func main () {
	var name string = "Tech Space"
	const pi float64 = 3.14124

	fmt.Println(len(name))
	fmt.Println(name + " is a nice place")

	fmt.Printf("%f \n", pi)
	fmt.Printf("%T \n", pi)
}