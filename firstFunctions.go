package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func multiply(x int, y int) int{
  return x * y
}

func subtract(x int, y int) int {
  return x - y
}

func main() {
	fmt.Println(add(42, 13))
  fmt.Println(add(40,59))
  fmt.Println(multiply(10,10))
  fmt.Println(subtract(80,20))
}
