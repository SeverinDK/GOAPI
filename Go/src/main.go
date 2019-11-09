package main

import (
	"log"
	"net/http"
	"server/config"
)

func main() {
	c, err := config.NewConnection("mysql", "docker:docker@tcp(db:3306)/test_db")

	if err != nil {
		log.Panic(err)
	}

	s := &config.Server{
		Env: &config.Env{
			Connection: c,
		},
	}

	r := config.InitializeRouter(s)
	listenAndServe(r)
}

func listenAndServe(r *config.Router) {
	log.Fatal(http.ListenAndServe(":8080", r.Router))
}
