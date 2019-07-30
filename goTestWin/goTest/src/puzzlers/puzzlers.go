package main

import (
	"fmt"
)

//container := []string {"zero", "one", "two"} //syntax error: non-declaration statement outside function body
var container = []string {"zero", "one", "two"}

func t_test() {
	fmt.Printf("container[1]=%v \n", container[1])
	container := map[int]string {3:"aaa", 2:"bbb", 1:"cccc"}
	fmt.Printf("container[1]=%v \n", container[1])
}

func main() {
	fmt.Println("==== Begin Main ====\n")
	t_test()
	fmt.Println("==== End Main ====\n")
}