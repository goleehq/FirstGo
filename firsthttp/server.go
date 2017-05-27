package main

import (
	"net/http"
	"log"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()//解析参数
	fmt.Println(r.Form)
	fmt.Fprintf(w, "hello go!")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("LinsenAndServe", err)
	}
}

