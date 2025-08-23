package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-sudo/webkelaspplg2/config"
	homecontroller "github.com/your-sudo/webkelaspplg2/controllers/homeController"
)

func main(){
  config.DBconnection()

  http.HandleFunc("/", homecontroller.Welcome)

  fmt.Println("Starting server on :8080")
  http.ListenAndServe(":8080", nil)
  log.Fatal(http.ListenAndServe(":8080", nil))
}


