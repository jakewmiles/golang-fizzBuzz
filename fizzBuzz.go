package main

import "fmt"

func fizzBuzz(n, x, y int) {
	for i := 1; i <= n; i++ {
		if i%(x*y) == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%y == 0 {
			fmt.Println("Buzz")
		} else if i%x == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzBuzz(50, 3, 7)
}
