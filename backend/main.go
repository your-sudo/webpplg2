package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-sudo/webkelaspplg2/config"
	"github.com/your-sudo/webkelaspplg2/controllers/homeController"
	logincontroller "github.com/your-sudo/webkelaspplg2/controllers/loginController"
)

func main() {
	config.DBconnection()

	http.HandleFunc("/", homecontroller.Welcome) 
	  http.HandleFunc("/login", logincontroller.Login) 

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
