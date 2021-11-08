package seeds

import (
	"github.com/bxcodec/faker/v3"
	"log"
	"math/rand"
)

var cities = [5]string{"Moscow", "NewYork", "Paris", "Dallas", "Milan"}
var err error

// TestSeed
func (s Seed) DeptSeed() {
	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(cities))

		_, err = s.db.Exec(
			`INSERT INTO dept(dname, loc) VALUES ($1, $2)`, faker.Word(), cities[randomIndex])
		if err != nil {
			log.Fatalf("error seeding test_seeds: %v", err)
		}
	}
}
