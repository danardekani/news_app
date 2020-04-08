package main

import (
	"net/http"
	"os"
)

func webHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

//Create web server
func main()  {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
}