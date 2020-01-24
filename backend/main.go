package main

import (
	"backend/controller"
	"fmt"
	"net/http"
)

const (
	port string = "8000"
)

func main() {
	start()
}

func start() {
	// fmt.Println("Start webserver...")
	// router := mux.NewRouter()
	// //API routes

	// router.HandleFunc("/api/save", controller.Save).Methods("POST")
	// router.HandleFunc("api/get", controller.Get).Methods("POST")

	// // Routes to serve the webpage
	// router.PathPrefix("/").Handler(http.HandlerFunc(controller.Serve))

	// // Use the JWT middle ware
	// // Set the cors
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
	// 	AllowCredentials: true,
	// })
	// handler := c.Handler(router)
	// fmt.Printf("Serving frontend on http://127.0.0.1:%s\n", port)
	// fmt.Printf("Api end routes exposed on port %s\n", port)
	// err := http.ListenAndServe(":"+port, handler) //Launch the app, visit localhost:5555/api
	// if err != nil {
	// 	fmt.Print(err)
	// }

	http.HandleFunc("/", controller.Serve)
	http.HandleFunc("/api/save", controller.Save)
	http.HandleFunc("/api/get", controller.Get)
	fmt.Println("Webserver is on http://127.0.0.1:8000")
	http.ListenAndServe(":"+port, nil)
}