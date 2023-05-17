package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Get("/", indexHandler)
	app.Post("/", postHandler)
	app.Put("/update", putHandler)
	app.Delete("/delete", deleteHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("index")
}

func postHandler(c *fiber.Ctx) error {
	return c.SendString("post")
}

func putHandler(c *fiber.Ctx) error {
	return c.SendString("put")
}

func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("delete")
}
