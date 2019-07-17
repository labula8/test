
package main

import (
	"fmt"
	"flag"
	"os"
)

func test() {
	//user_name := "aaaa"
	user_name := flag.String("user", "", "Input your name:")
	count := flag.Int("c", 0, "Input count:")
	is_flag := flag.Bool("is", false, "Input is_flag:")
	
	flag.Parse()
	fmt.Println("Hello ", *user_name)
	fmt.Println("count ", *count)
	fmt.Println("is_flag ", *is_flag)
}

func test2()  {
	var user_name = "test"
	fmt.Println("Hello ", user_name)
	flag.StringVar(&user_name, "user", "", "Input your name:")
	
	flag.Parse() //MUST Parse
	
	fmt.Printf("Hello %s \n", user_name)
}

//$ ./flagdemo.exe -n=false xxx n bbb  cccc
func test3() {
	var is_newline = flag.Bool("n", false, "is new line")
	const (
		Space = " "
		NewLine = "\n"
	)
	
	var s string = ""
	
	flag.Parse()
	//打印，参数列表
	for i:=0; i<flag.NArg(); i++ {
		var s_temp string = flag.Arg(i)
		fmt.Printf("i=%d, s_temp=%s\n", i, s_temp)
		
		if s_temp == "#" {
			break
		}
		
		if i>0 {
			s += Space
		}
		s += s_temp
		
		if !(*is_newline) {
			s += NewLine
		}
	}
	fmt.Printf("is_newline=%t\n", *is_newline)
	fmt.Printf("s=%s\n", s)
	
	os.Stdout.WriteString(s)
}

func main() {
	//test()
	//test2()
	test3()
}