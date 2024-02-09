package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main () {
	fmt.Println("Hello world")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment variables")
	}

	router := chi.NewRouter()

	srv := &http.Server {
		Handler: router,
		Addr: ":" + portString,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}