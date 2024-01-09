package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wanyansu/holidayinn/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The port where our API server listens to")
	flag.Parse()
	app := fiber.New()
	appv1 := app.Group("/api/v1")

	appv1.Get("/user", api.HandleGetUsers)
	log.Fatal(app.Listen(*listenAddr))
}
