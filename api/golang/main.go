package main

import (
	"fmt"
	"log"
	"net/http"

	"app/config"
	"app/controllers"

	"github.com/gorilla/mux"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World !!")
}

func main() {
	r := mux.NewRouter()

	// :/
	r.HandleFunc("/", helloHandle)
	controllers.Serve(r)

	// Run Server
	log.Printf("Running at :%d\n", config.Config.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Server.Port), r)
	if err != nil {
		panic(err)
	}
}
