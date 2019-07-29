
package main

import (
	"fmt"
	"os"
	"bufio"
	"reflect"
	"strings"
)

func t_file_countlines(f *os.File, map_count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Printf("input.Text()=%s \n", input.Text())
		if input.Text() == "#" {
			break
		} else {
			map_count[input.Text()]++
		}
	}
	
	fmt.Printf("map_count=%v \n", map_count)
}

func t_file_count(fileName string) {
	//var fileName string = "a.txt"
	if file, err := os.Open(fileName); err != nil {
		fmt.Printf("Open file fail, fileName=%s, error=%v\n", fileName, err)
	} else {
		defer file.Close()
		fmt.Printf("Open file ok, fileName=%s, error=%v\n", fileName, err)
		//var map_count map[string]int
		map_count := make(map[string]int)
		t_file_countlines(file, map_count)
		fmt.Printf("map_count=%v \n", map_count)
	}
}

func t_file_open(file_in string) bool {
	var fileName = file_in
	if 0 == len(fileName) {
		fileName = "a.txt"
		fmt.Printf("Input fileNmae is NULL, Dfault fileName=%s\n", fileName)
	} else {
		fmt.Printf("Input fileNmae=%s\n", fileName)
	}
	
	file, err := os.Open(fileName)
	if err != nil { 
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		fmt.Printf("Program stopping with error %v", err)
		os.Exit(1)
		return false
	} else {
		iReader := bufio.NewReader(file)
		for {
			//str, err := iReader.ReadString('\n')
			str, _, err := iReader.ReadLine()
			if err != nil {
				fmt.Printf("Program stopping with error %v", err)
				return false//error or EOF
			} else {
				fmt.Printf("Read: %s\n", str)
			}
		}
	}
	defer file.Close()
	return true
}

func t_file() {
	//var fileName = "b.txt"
	//t_file_open(fileName)
	//t_file_open("")
	t_file_count("a.txt")
}

func t_args() {
	args := os.Args
	fmt.Printf("args=%v \n", args)
	fmt.Printf("TypeOf(args)=%v \n", reflect.TypeOf(args))
	fmt.Printf("len(args)=%v \n", len(args))
	if len(args) > 0 {
		fmt.Printf("args[1]=%v,os.Args[1]=%v \n", args[1], os.Args[1])
	}
	
	args_paras := args[1:]
	fmt.Printf("args_paras=%v \n", args_paras)
	fmt.Printf("TypeOf(args_paras)=%v \n", reflect.TypeOf(args_paras))
	
	for pos, arg := range args_paras {
		fmt.Printf("pos=%d, arg=%s \n",  pos, arg)
	}
	
	for _, arg := range args[1:] {
		fmt.Printf("arg=%s \n", arg)
		t_file_count(arg)
	}
}

func t_strings_join() {
	//var s1 string = "hello"
	//s2 := strings.Join(s1, "#") //cannot use s1 (type string) as type []string in argument to strings.Join
	
	//map1 := make([]string) //missing len argument to make([]string)
	map1 := []string {"aaa", "hello", "ggggo"}
	s2 := strings.Join(map1, "#")
	fmt.Printf("map1=%s, s2=%s \n", map1, s2)
}

func t_test() {
	//t_file()
	//t_args()
	t_strings_join()
}

func main() {
	fmt.Println("==== Begin Main ====\n")
	t_test()
	fmt.Println("==== End Main ====\n")
}