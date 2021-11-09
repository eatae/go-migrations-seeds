package seeds

import (
	"github.com/bxcodec/faker/v3"
	"log"
	"math/rand"
	"time"
)

var jobs = [6]string{"Clerk", "Manager", "Salesman", "Analyst", "Hr", "Cleaner"}

// Emp Seed
func (s Seed) EmpSeed() {
	var err error
	var mgrNo interface{}

	deptsNo := s.getDeptsNo()

	for i := 1; i <= 100; i++ {
		mgrNo = nil
		if i%5 == 0 {
			mgrNo = s.getRandomManagerNo()
		}
		jobsRandIdx := rand.Intn(len(jobs))
		depsNoRandIdx := rand.Intn(len(deptsNo))

		// Insert random data
		_, err = s.db.Exec(
			"INSERT INTO emp(ename, job, mgr, hiredate, sal, comm, deptno) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			faker.Name(),
			jobs[jobsRandIdx],
			mgrNo,
			s.getRandomDate(),
			s.getRandomSalary(),
			s.getRandomComm(),
			deptsNo[depsNoRandIdx])
		if err != nil {
			log.Fatalf("error seeding test_seeds: %v", err)
		}

	}
}

// getDeptsNo
func (s Seed) getDeptsNo() []int {
	var result []int
	var deptno int

	rows, err := s.db.Query("SELECT deptno FROM dept")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&deptno)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, deptno)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// Get random Manager number
func (s Seed) getRandomManagerNo() int {
	var empNo int
	err := s.db.QueryRow("SELECT empno FROM emp ORDER BY random() LIMIT 1").Scan(&empNo)
	if err != nil {
		log.Fatal(err)
	}
	return empNo
}

// Get random date
func (s Seed) getRandomDate() string {
	min := time.Date(1975, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2002, 12, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	t := time.Unix(sec, 0)

	return t.Format("2006-01-02")
}

// Get random salary
func (s Seed) getRandomSalary() int {
	var salaries = []int{1000, 1200, 1300, 1400, 1500, 1600, 1700, 2000, 2200, 2500}
	randIdx := rand.Intn(len(jobs))

	return salaries[randIdx]
}

// Get random comm
func (s Seed) getRandomComm() int {
	var salaries = []int{0, 0, 130, 100, 500, 60, 700, 0, 0, 0}
	randIdx := rand.Intn(len(jobs))

	return salaries[randIdx]
}
