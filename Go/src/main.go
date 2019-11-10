package main

import (
	"log"
	"server/config"
	"server/controllers"
)

func main() {
	c, err := config.NewConnection("mysql", "backend_api:securepassword@tcp(db:3306)/backend_api_db")

	if err != nil {
		log.Panic(err)
	}

	s := &config.Server{
		Env: &config.Env{
			Connection: c,
		},
	}

	controllers.InitializeRouter(s)
}
