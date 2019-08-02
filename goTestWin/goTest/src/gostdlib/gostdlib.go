package main

import (
	"fmt"
	"io"
	"os"
	//"bufio"
	"strings"
)

func t_ReadFrom(reader io.Reader, num int) ([]byte, error){
	fmt.Printf("t_ReadFrom, reader=%v, num=%d \n", reader, num)
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		fmt.Printf("Read n=%d, err=nil \n", n)
		return p, nil
	} else {
		fmt.Printf("Read n=%d, err=%v \n", n, err)
		return p, err
	}
}

func t_file_open(fileName string) (f *os.File, err error) {
	if fileName == "" {
		fileName = "a.txt"
	}
	
	f, err = os.Open(fileName)
	if err != nil { 
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		fmt.Printf("Program stopping with error %v", err)
		os.Exit(1)
		return f, err
	} else {
		fmt.Printf("Opening the inputfile ok, fileName=%s \n", fileName)
		/*
		iReader := bufio.NewReader(file)
		for {
			//str, err := iReader.ReadString('\n')
			str, _, err := iReader.ReadLine()
			if err != nil {
				fmt.Printf("Program stopping with error %v", err)
				return file, err //error or EOF
			} else {
				fmt.Printf("Read: %s\n", str)
			}
		}
		*/
	}
	//defer file.Close()
	return f, nil
}

func t_ReadFrom_file() ([]byte, error){
	var data []byte
	
	file, err := t_file_open("a.txt")
	if err != nil {
		fmt.Printf("err=%v \n", err)
	} else {
		data, err = t_ReadFrom(file, 10)
		if len(data) > 0 {
			for _, ch := range data {
				fmt.Printf("ch=%d, %c\n", ch, ch)
			}
		}
	}
	defer file.Close()
	return data, err
}

func t_ReadFrom_string() ([]byte, error) {
	s_test := "s-test-string"
	s_reader := strings.NewReader(s_test)
	return t_ReadFrom(s_reader, 10)
}

func t_io_read() {
	
	//var data []byte
	
	// 从标准输入读取
	//data, err := t_ReadFrom(os.Stdin, 10)
	
	// 从文件读取
	//data, err := t_ReadFrom_file()
	
	// 从字符串读取
	data, err := t_ReadFrom_string()
	
	fmt.Printf("Read data=%v, err=%v \n", data, err)
}

func t_test() {
	fmt.Println("==== gostdlib ====\n")
	t_io_read()
}

func main() {
	fmt.Println("==== Begin Main ====\n")
	t_test()
	fmt.Println("==== End Main ====\n")
}