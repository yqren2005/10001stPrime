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
