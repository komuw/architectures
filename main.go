package main

import (
	"fmt"
	"net/http"

	"github.com.komuw/architectures/clean"
	"github.com.komuw/architectures/myio"
)

func run() error {
	http.HandleFunc("/clean", clean.AddBookHandler)
	http.HandleFunc("/myio", myio.AddBookHandler)

	port := "8080"
	fmt.Println("listening on : ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	return err
}

/*
curl -vkL localhost:8080/clean
curl -vkL localhost:8080/myio
*/
func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
