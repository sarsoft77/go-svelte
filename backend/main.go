package main

import (
	"log"

	"github.com/example/myapp/db"
	"github.com/example/myapp/router"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := router.Setup()
	r.Run(":8080")
}
