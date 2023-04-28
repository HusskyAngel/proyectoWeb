package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/HusskyAngel/routes"
) 

func main() {

    fmt.Println("hola")
    r:=mux.NewRouter()
    r.HandleFunc("/", routes.HomeHandler  )
    http.ListenAndServe(":5000",r) 
}
