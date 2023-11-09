package main

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Mongodb API")
	r := router.Router()
	fmt.Println("Server is getting ready ...")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening on port : 4000!")

}
