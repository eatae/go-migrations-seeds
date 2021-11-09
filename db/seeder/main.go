package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go_migrations_seeds/db/seeder/seeds"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file")
	}
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DOCKER_DB_HOST"),
		os.Getenv("DOCKER_DB_PORT"),
		os.Getenv("DOCKER_DB_USER"),
		os.Getenv("DOCKER_DB_PASSWORD"),
		os.Getenv("DOCKER_DB_NAME"),
	)

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatalf("error opening a connection with the database %s\n", err)
	}

	s := seeds.NewSeed(db)
	s.DeptSeed()
	s.EmpSeed()

	//s := seeds.NewSeed(db)
	//
	//if err := seeder.Execute(s); err != nil {
	//	log.Fatalf("error seeding the db %s\n", err)
	//}
}
