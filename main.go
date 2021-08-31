package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func db() {
	psqlconn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("PSQL_USER"), os.Getenv("PSQL_PWD"), os.Getenv("PSQL_HOST"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_DB"))
	fmt.Println(psqlconn)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
}
func main() {
	godotenv.Load()
	db()
	app := fiber.New()
	app.Get("/v1/verify/key", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
}
