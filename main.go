package main

import "fmt"

func add(x, y int) {
	fmt.Println(x + y)
}

func subtract(x, y int) {
	fmt.Println(x - y)
}

func multiply(x, y int) {
	fmt.Println(x * y)
}

//---BEFORE

//func main() {
//
//	method := "somar"
//
//	if method == "somar" {
//		add(1, 2)
//	}
//	if method == "subtrair" {
//		subtract(1, 2)
//	}
//	if method == "multiplicar" {
//		multiply(1, 2)
//	}
//
//	switch method {
//	case "somar":
//		add(1, 2)
//	case "subtrair":
//		subtract(1, 2)
//	case "multiplicar":
//		multiply(1, 2)
//	default:
//		fmt.Println("this method don't exists!")
//		return
//	}
//
//}

//---AFTER

var functions = map[string]func(y, z int){
	"somar":       add,
	"subtrair":    subtract,
	"multiplicar": multiply,
}

func main() {

	env := "somar"

	fn, ok := functions[env]
	if !ok {
		fmt.Printf("you called %s this method don't exists on map, try again!", env)
		return
	}

	fn(1, 2)
}
