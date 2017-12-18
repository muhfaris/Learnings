package main

import "fmt"

func fibonanci(x int) int {
	if x <= 3 {
		return 1
	} else {
		return fibonanci(x-1) + fibonanci(x-2)
	}
}

func main() {
	value := 6
	x := fibonanci(value)
	fmt.Println("Fibonanci", value, "is ", x)
}
