package main

import "fmt"

func main() {

	rev("hello world")
	rev("roop")
	rev("davinder")
}

func rev(str string) {

	// s := ""
	// for i := len(str) - 1; i >= 0; i-- {
	// 	s += string(str[i])
	// }
	// fmt.Println(s)

	// var s strings.Builder

	// for i := len(str) - 1; i >= 0; i-- {
	// 	s.WriteByte(str[i])
	// }

	// fmt.Println(s.String())

	s := ""
	for _, n := range str {
		s = string(n) + s
	}
	fmt.Println(s)

}
