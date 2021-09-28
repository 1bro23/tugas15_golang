package main

import "fmt"
import "net/http"

func main(){
  http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
    for i:=0;i<=100;i++{
      fmt.Fprint(w,i,", ")
    }
  })
  fmt.Println("Server start at localhost:8080")
  http.ListenAndServe(":8080",nil)
}
