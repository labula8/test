package main

import (
	"fmt"
	"time"
	"reflect"
)

func boring(msg string) {
	for i:=0; i<10000; i++ {
		fmt.Println(msg, i)
		var d = time.Second
		time.Sleep(d)
	}
}

func test() {
	boring("hello")
}

func test_reflect() {
	var n int32 = 20
	var n_ref = reflect.TypeOf(n)
	fmt.Println("type:", reflect.TypeOf(n))
	//fmt.Println("n_ref:", reflect.TypeOf(n_ref)) //n_ref: *reflect.rtype
	var s_ref = n_ref.String()
	fmt.Printf("s_ref:%s", s_ref)
}

func get_type(v interface {}) string {
	return reflect.TypeOf(v).String()
}

func test_get_type() {
	var type_s = get_type("hello")
	fmt.Printf("type_s:%v \n", type_s)
	var type_n = get_type(10)
	fmt.Printf("type_n:%v \n", type_n)
}

func main() {
	fmt.Println("Hello")
	//test()
	//test_reflect()
	test_get_type()
}

