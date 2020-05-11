package main

import "fmt"

func main() {

	sum([]int{1, 2, 3, 4, 5})
	sum([]int{3, 3, 3, 3, 4})
	sum([]int{})
	sum([]int{11, -10, 33, 54, 25})
}

func sum(list []int) {
	t := 0
	for _, n := range list {
		t += n
	}
	fmt.Println(t)
}
