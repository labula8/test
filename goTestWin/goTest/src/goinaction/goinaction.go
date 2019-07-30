package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
	"sync/atomic"
)

//全局公共变量counter
var (
	counter int64
	wg sync.WaitGroup
	mutex sync.Mutex
)

func show_a_z() {
	for count := 0; count < 3; count++ {
		for ch := 'a'; ch <= 'z'; ch ++ {
			//fmt.Printf("ch=%c ", ch)
			fmt.Printf("%c ", ch)
		}
		time.Sleep(time.Duration(1)*time.Second)
		fmt.Printf("#%d# \n", count)
	}
}

func show_A_Z() {
	for count := 0; count < 3; count++ {
		for ch := 'A'; ch <= 'Z'; ch ++ {
			fmt.Printf("%c ", ch)
		}
		time.Sleep(time.Duration(1)*time.Second)
		fmt.Printf("#Big %d# \n", count)
	}
}

func t_Hello() {
	time.Sleep(2*250 * time.Millisecond)
	fmt.Println("t_Hello() Hello World")
}

func t_grouting_1() {
	fmt.Printf("Start t_grouting_1() \n")
	
	//设置并行，cpu 数量
	//runtime.GOMAXPROCS(1) //大小写分开
	runtime.GOMAXPROCS(2) // 大小写会混合
	
	var wg sync.WaitGroup
	wg.Add(2)
	
	//show_a_z()
	//go show_a_z()
	//go t_Hello()
	
	go func() {
		defer wg.Done()
		show_a_z()
	}()
	
	go func() {
		defer wg.Done()
		show_A_Z()
	}()

	fmt.Printf("Waiting to finish \n")
	wg.Wait()
	
	time.Sleep(2*time.Second)
	fmt.Printf("End t_grouting_1() \n")
}

//go build -race goinaction.go
func add_counter(id int) {
	defer wg.Done()
	
	for i:=0; i<2; i++ {
		//fmt.Printf("in 11 add_counter(%d) counter=%d \n", id, counter)
		value := counter
		//grouting 退出，放回队列？
		runtime.Gosched()
		value ++
		counter = value
		
		fmt.Printf("in 22 add_counter(%d) counter=%d \n", id, counter)
	}
	
	/*
	for i:=0; i<2; i++ {
		fmt.Printf("in 11 add_counter(%d) counter=%d \n", id, counter)
		//grouting 退出，放回队列？
		runtime.Gosched()
		counter++
		
		fmt.Printf("in 22 add_counter(%d) counter=%d \n", id, counter)
	}
	*/
}

func t_grouting_2() {
	fmt.Printf("in t_grouting_2() counter=%d \n", counter)
	wg.Add(2)
	
	go add_counter(1)
	go add_counter(2)
	
	wg.Wait()
	fmt.Printf("out t_grouting_2() counter=%d \n", counter)
}

func add_counter_atomic(id int) {
	defer wg.Done()

	for i:=0; i<2; i++ {
		//fmt.Printf("in 11 add_counter(%d) counter=%d \n", id, counter)
		//atomic.AddInt64(&counter, 1) //cannot use &counter (type *int) as type *int64 in argument to "sync/atomic".AddInt64
		//atomic.AddInt64(&int64(counter), 1)// cannot take the address of int64(counter)
		atomic.AddInt64(&counter, 1)
		//grouting 退出，放回队列？
		runtime.Gosched()
		fmt.Printf("in 22 add_counter_atomic(%d) counter=%d \n", id, counter)
	}
}

func add_counter_mutex(id int) {
	defer wg.Done()

	for i:=0; i<2; i++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value ++
			counter = value
		}
		mutex.Unlock()

		fmt.Printf("in 22 add_counter_mutex(%d) counter=%d \n", id, counter)
	}
}

//atomic 原子锁，解决共享变量，冲突问题
func t_grouting_3() {
	fmt.Printf("in t_grouting_3() counter=%d \n", counter)
	wg.Add(2)
	
	go add_counter_atomic(1)
	go add_counter_atomic(2)
	
	wg.Wait()
	fmt.Printf("out t_grouting_3() counter=%d \n", counter)
}

//mutex 临界区锁，解决共享变量，冲突问题
func t_grouting_4() {
	fmt.Printf("in t_grouting_4() counter=%d \n", counter)
	wg.Add(2)
	
	go add_counter_mutex(1)
	go add_counter_mutex(2)
	
	wg.Wait()
	fmt.Printf("out t_grouting_4() counter=%d \n", counter)
}

func t_test() {
	fmt.Print("t_test()\n")
	//t_grouting_1()
	//t_grouting_2()
	t_grouting_3()
	//t_grouting_4()
}

func main() {
	fmt.Println("==== Begin Main ====")
	t_test()
	fmt.Println("==== End Main ====")
}