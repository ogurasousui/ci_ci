package main

import "fmt"

func sayHello() string {
	return "Hello, World!"
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(sayHello())
	fmt.Println("2 + 3 =", add(2, 3))
}
