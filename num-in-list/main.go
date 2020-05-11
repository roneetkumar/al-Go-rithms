package main

import "fmt"

func main() {

	numInList([]int{1, 2, 3, 4, 5}, 5)
	numInList([]int{3, 3, 3, 3, 4}, 5)
	numInList([]int{3, 5, 6, 7, 3}, 5)
	numInList([]int{11, -10, 33, 54, 25}, -10)

}

func numInList(list []int, num int) {
	f := false
	for _, n := range list {
		if n == num {
			f = true
			break
		}
	}
	if f {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
