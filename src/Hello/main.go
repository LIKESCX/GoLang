package main

import (
	"fmt"
)

func main(){

	fmt.Println("Hello, World!")
	var a int = 4
	switch a {
		case 0:
			fmt.Println("0")
		case 1:
			fmt.Println("1")
		case 2:
			fallthrough;
		case 3:
			fmt.Println("3")
		case 4, 5, 6:
			fmt.Println("4, 5, 6")
		default:
			fmt.Println("Default")
	}
}