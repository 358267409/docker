// go-web project main.go
package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		fmt.Println("ListenAndServe: ", er)
	}
}
