
package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	a int;		//8
	b int;		//8
	c float64;	//8
	d int32;	//4
	e int16; 	//2
}

/*

type T struct {
	a int;		//8
	b int;		//8
	c float64;	//8
	d int32;	//4
}

Hello
a=10, b=0
c=10.888000
d=999
unsafe.Sizeof(t)=8
unsafe.Sizeof(*t)=32
unsafe.Sizeof(t.a)=8
unsafe.Sizeof(t.c)=8
unsafe.Sizeof(t.d)=4

*/
func test() {
	//var t = T() //missing argument to conversion to T: T()
	//var t = T //type T is not an expression
	//var t = new T //unexpected T at end of statement
	//var t = new T() //unexpected T at end of statement
	//var t T = new(T) // cannot use new(T) (type *T) as type T in assignment
	var t *T = new (T)
	//var t = new(T) //ok
	
	t.a = 10
	t.c = 10.888
	t.d = 999
	t.e = 22
	
	fmt.Printf("a=%d, b=%d \n", t.a, t.b);
	fmt.Printf("c=%f \n", t.c);
	fmt.Printf("d=%d \n", t.d);
	fmt.Printf("e=%d \n", t.e);
	
	//fmt.Printf("a=%d, b=%d", t->a, t->b); //unexpected >, expecting expression
	
	fmt.Printf("unsafe.Sizeof(t)=%d \n", unsafe.Sizeof(t))
	fmt.Printf("unsafe.Sizeof(*t)=%d \n", unsafe.Sizeof(*t))
	fmt.Printf("unsafe.Sizeof(t.a)=%d \n", unsafe.Sizeof(t.a))
	fmt.Printf("unsafe.Sizeof(t.c)=%d \n", unsafe.Sizeof(t.c))
	fmt.Printf("unsafe.Sizeof(t.d)=%d \n", unsafe.Sizeof(t.d))
	fmt.Printf("unsafe.Sizeof(t.e)=%d \n", unsafe.Sizeof(t.e))
}

func main() {
	fmt.Println("Hello")
	test()
}