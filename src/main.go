package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"github.com/ruanpienaar/godb"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(7*24*60*60*time.Second)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST" {
			fmt.Println("Received POST")
		} else if r.Method == "GET" {
			fmt.Println("Received GET")
			// w.Write( []byte )
		}
	})
	log.Fatal(http.ListenAndServe(":8801", nil))
}