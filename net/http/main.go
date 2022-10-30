package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	file, _ := os.ReadFile("./hello.txt")
	fmt.Fprintf(w, string(file))
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}
}
