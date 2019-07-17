
package main

import (
	"fmt"
)

func show_arry(arry []int, size int) {
	for i:=0; i<size; i++ {
		fmt.Printf("%d, %d\n", i, arry[i]);
	}
	
	fmt.Printf("%v \n", arry)
}

func test() {
	//var arry1[3] int {1,3,8} //error
	var arry1 = []int{1,3,8}
	for i:=0; i<3; i++ {
		fmt.Printf("%d, %d\n", i, arry1[i]);
	}
	
	//test show_arry
	show_arry(arry1, 3)
	var n_sum = sum(arry1, 3)
	fmt.Printf("n_sum=%d", n_sum)
}

func sum(arry []int, size int) int{
	var sum = 0
	for i:=0; i<size; i++ {
		sum += arry[i]
	}
	return sum
}

func main() {
	fmt.Println("Hello")
	test()
}