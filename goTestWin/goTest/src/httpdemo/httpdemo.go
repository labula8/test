
package main

import (
	"fmt"
	"net/http"

	//"github.com/tabalt/gracehttp"
)

func main() {
	var count int = 0
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello world")
			fmt.Println(count)
			count++;
	})
	
	//err := gracehttp.ListenAndServe(":8088", nil)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println(err)
	}
}