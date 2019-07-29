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

func t_Hello() {
	time.Sleep(2*250 * time.Millisecond)
	fmt.Println("t_Hello() Hello World")
}

func t_DelayPrint() {
	for i:=0; i<5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("t_DelayPrint %d\n", i)
	}
}

func t_groutine() {
	//t_Hello()
	go t_DelayPrint()
	go t_Hello()
	//t_DelayPrint()
	//t_Hello()
	
	time.Sleep(2*time.Second)
}

var done chan bool

func t_channel_Hello() {
	fmt.Println("t_channel_Hello() Hello World")
	time.Sleep(2*time.Second)
	done <- true
}

func t_channel() {
	/*
	var ch1 chan int
	//ch1 := make(chan int)
	var value int = 5
	ch1 <- value
	var value2 = <-ch1
	fmt.Printf("t_channel value2=%d\n", value2)
	close(ch1)
	*/
	done = make(chan bool)
	go t_channel_Hello()
	
	//<-done
	fmt.Printf("t_channel() <-done=%v\n", <-done)
	close(done)

	//fmt.Printf("t_channel() <-done=%v\n", <-done)
	//close(done)
}

func main() {
	fmt.Println("main() Begin")
	//test()
	//test_reflect()
	//test_get_type()
	t_groutine()
	//t_channel();
	fmt.Println("main() End")
}

