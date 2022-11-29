package main

import (
	"fmt"
	"net/http"

	"architectures/clean"
	"architectures/myio"
)

func run() error {
	http.HandleFunc("/clean", clean.AddBookHandler)
	http.HandleFunc("/myio", myio.AddBookHandler)

	port := "8080"
	fmt.Println("listening on : ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	return err
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
