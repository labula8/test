package main 

import (
    "flag"
    "fmt"
    "time"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func test() {
	var t = time.Now()
	fmt.Printf("t=%v \n", t)
	
	var t_unix = t.Unix()
	var t_unix_nano = t.UnixNano()
	var t_nano = t.Nanosecond()
	var t_unix_nano_sec = t_unix_nano / 1e9
	c := time.Unix(t_unix_nano_sec, 0)
	
	fmt.Printf("t_unix=%v \n", t_unix)
	fmt.Printf("t_unix_nano=%v \n", t_unix_nano)
	fmt.Printf("t_nano=%v \n", t_nano)
	fmt.Printf("t_unix_nano_sec=%v \n", t_unix_nano_sec)
	fmt.Printf("c=%v \n", c)
}

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    
    test()
}

