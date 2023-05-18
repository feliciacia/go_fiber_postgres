package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	connStr := "postgresql://<postgres>:<postgres>@<localhost>/todos?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c)
	})
	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c)
	})
	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos_rows []string
	query := "SELECT * FROM todos"
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos_rows = append(todos_rows, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos_rows,
	})
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
