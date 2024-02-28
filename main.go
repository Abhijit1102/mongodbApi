package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Abhijit1102/mongodbApi/router"
)

func main() {
	fmt.Println("Mongodb API ")
	r := router.Router()
	fmt.Println("Serve is getting Started ...")
	log.Fatal(http.ListenAndServe{":4000", r})
	fmt.Println("Listing at Port 4000 ...")

	
}