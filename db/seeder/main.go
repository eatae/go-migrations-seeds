package main

import (
	"fmt"
	_ "github.com/joho/godotenv"
	"os"
	// postgres driver.
	_ "github.com/lib/pq"
)

func main() {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DOCKER_DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(url)
	os.Exit(0)

	//db, err := sqlx.Open("postgres", url)
	//if err != nil {
	//	log.Fatalf("error opening a connection with the database %s\n", err)
	//}
	//
	//s := seeds.NewSeed(db)
	//
	//if err := seeder.Execute(s); err != nil {
	//	log.Fatalf("error seeding the db %s\n", err)
	//}
}
