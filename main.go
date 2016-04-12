package main

import (
	"fmt"
)

func main() {
	a, b := 2, 0
	for {
		tag := false
		for i := 2; i < a; i++ {
			if a%i == 0 {
				tag = true
				break
			}
		}
		if tag == false {
			b += 1
			if b == 10001 {
				fmt.Println(a)
				break
			}

		}
		a += 1
	}
}

/*
Project Euler Q7:

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

Answer: 104743
*/
