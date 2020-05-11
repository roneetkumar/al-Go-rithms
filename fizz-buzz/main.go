package main

import (
	"fmt"
	"strconv"
)

func main() {

	fizz(5)
}

func fizz(n int) {

	s := ""

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			s += "Fizz"
		} else if i%5 == 0 {
			s += "Buzz"
		} else if i%5 == 0 && i%3 == 0 {
			s += "Fizz, Buzz"
		} else {
			s += strconv.Itoa(i)
		}

		if i != n {
			s += ", "
		}
	}
	fmt.Println(s)
}
