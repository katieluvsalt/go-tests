package main

import (
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "github.com/katieluvsalt/API/stuff"
	"log"
	"github.com/katieluvsalt/microservicesCloud/dal"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()
	dal.InitDB();

	log.Fatal(http.ListenAndServe(":8080", router))

}
