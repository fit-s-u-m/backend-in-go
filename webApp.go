package main

import (
	"fmt"
	"net/http"
)


func WebApp(){
  http.HandleFunc("/", handleFunc)
  fmt.Println("starting the server at port 8080")
  http.ListenAndServe(":8080", nil)
}
func handleFunc(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w,"Hello World")
}
