package main

import "fmt"

func main (){
	x := 44

	valueChange(&x) // & points to X
	fmt.Println(x) 
}
func valueChange (x *int){
	*x = 69
}