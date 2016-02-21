package main

import (
	"RestExample/rest"
	"net/http"
)

func main() {
  serve := rest.GetServe()
	http.HandleFunc("/", serve.ServeJson)
	http.ListenAndServe("localhost:1337", nil)
}
