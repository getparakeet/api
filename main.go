package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type KeyData struct {
	Title string `json:"projectTitle"`
	Key   string `json:"key"`
}

func main() {
	godotenv.Load()
	app := fiber.New()
	app.Post("/v1/verify/key", func(res *fiber.Ctx) {
		fields := new(KeyData)
		if err := res.BodyParser(fields); err != nil {
			panic(err)
		}
		query := `SELECT "name", "key" FROM "keydata"`
		psqlconn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("PSQL_USER"), os.Getenv("PSQL_PWD"), os.Getenv("PSQL_HOST"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_DB"))
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		resp, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		}
		if strings.Contains(query, "SELECT") {
			for resp.Next() {
				var name string
				var key string

				err := resp.Scan(&name, &key)
				if err != nil {
					panic(err)
				}
				if fields.Key != key {
					res.Status(400).Send("Invalid key")
				} else if fields.Key == key {
					res.Status(200).Send("Valid key")
				}
			}
		}
		defer resp.Close()
	})
	app.Listen(":3000")
}
