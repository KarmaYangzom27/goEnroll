package main

import (
	"fmt"
	"myapp/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.InitializeRoutes()
	// dynamic route
	router.HandleFunc("/home/{course}", homeHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	course := params["course"]

	_, err := w.Write([]byte("hello world\nThis course is " + course))
	if err != nil {
		fmt.Println("error:", err)
	}
}
