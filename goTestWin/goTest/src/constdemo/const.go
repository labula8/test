package main 

import (
	"fmt"
)

func test() {
	var a int64 = 1
	fmt.Printf("a=%d \n", a);
	
	//a = 1 << (64-1) //warn: constant 9223372036854775808 overflows int64
	a = 1 << (64-2)
	fmt.Printf("a=%d \n", a);
	a = 1 << (32 -1)
	fmt.Printf("a=%d \n", a);
	
	a = 1000
	fmt.Printf("a=%d \n", a);
	
	var a64 uint64 = 1
	a64 = 1 << (64-1)
	fmt.Printf("a64=%d \n", a64);
	
	//const var b int = 1000 //unexpected var, expecting name
	//var const b int = 100 //unexpected const, expecting name
	//b int := 10 //unexpected int at end of statement
	b := 101 //普通变量
	//const b int = 10 //ok, 常量定义
	b = 10	//cannot assign to b
	fmt.Printf("b=%d \n", b);
	
	b = 0x1234
	fmt.Printf("b=%d \n", b);
	
	b = 1e6
	fmt.Printf("b=%d \n", b);
	
	f := 1e-6 //ok
	//f := 1e(-6) //cannot call non-function 0 (type float64)
	fmt.Printf("f=%f \n", f);
	
	//f = 2/3 //TODO f=0.000000
	//f = 2/3.0
	//f = 2.0/3
	//f = 2./3
	f = 2/3.
	fmt.Printf("f=%f \n", f);
}

func main() {
	fmt.Println("Hello")
	test()
}

