
package main

import (
	"fmt"
	"math/rand"
	"math"
)

func add(a int, b int) int {
	return a + b
}

func swap(s1 string, s2 string) (string, string) {
	return s2, s1
} 

func test() {
	var n int = 100
	var n_rand int = rand.Intn(n)
	var f_sqrt float64 = math.Sqrt(7.0)
	fmt.Println("rand num: ", n_rand)
	fmt.Println("sqrt num: ", f_sqrt)
	fmt.Println("math pi: ", math.Pi)
	//var c int = add(3, 4)
	c := add(3, 4)
	fmt.Println("c is: ", c)
}

func test2() {
	s1 := "hello"
	s2 := "world"
	s3, s4 := swap(s1, s2)
	fmt.Println(s3)
	fmt.Println(s3, s4)
}

func t_type() {
	var flag bool = true
	fmt.Println(flag)
	
	if (flag) {
		fmt.Println("flag is: ", flag)
	}
	
	var f_num float32 = 3.11
	fmt.Println("f_num = ", f_num)
	
	var d_n1, d_n2 float64 = 3.18, 3.88
	fmt.Println(d_n1, d_n2)
}

func t_const() {
	const pi float32 = 3.14 //const前后没有关键字 var
	fmt.Println("pi ", pi)
	//pi = 3.88 //cannot assign to pi
}

func main() {
	//test()
	//test2()
	//t_type()
	t_const()
}