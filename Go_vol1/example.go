package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")

		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data is %d ", d)
	})

	http.HandleFunc("/goodBye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good Bye World")
	})
	http.ListenAndServe(":9090", nil)
}
