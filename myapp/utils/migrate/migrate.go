package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db_info := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		"dpg-d7tb6fvavr4c738ivmn0-a.singapore-postgres.render.com",
		5432,
		"postgres_admin",
		"rresBEvRCgnwcH8qVIGLIBjftnasvWH0",
		"my_db_qrnl",
	)

	db, err := sql.Open("postgres", db_info)
	if err != nil {
		panic(err)
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS student (StdId INT NOT NULL, FirstName VARCHAR(45) NOT NULL, LastName VARCHAR(45) DEFAULT NULL, Email VARCHAR(45) NOT NULL, PRIMARY KEY (StdId));`,
		`CREATE TABLE IF NOT EXISTS admin (FirstName VARCHAR(45) NOT NULL, LastName VARCHAR(45) DEFAULT NULL, Email VARCHAR(45) NOT NULL, Password VARCHAR(45) NOT NULL, PRIMARY KEY (Email));`,
		`CREATE TABLE IF NOT EXISTS course (cid VARCHAR(45) NOT NULL, cname VARCHAR(45) NOT NULL, credit INT DEFAULT NULL, PRIMARY KEY (cid));`,
		`CREATE TABLE IF NOT EXISTS enroll (std_id INT NOT NULL, course_id VARCHAR(45) NOT NULL, date_enrolled VARCHAR(45) DEFAULT NULL, PRIMARY KEY (std_id, course_id), CONSTRAINT course_fk FOREIGN KEY (course_id) REFERENCES course (cid) ON DELETE CASCADE, CONSTRAINT std_fk FOREIGN KEY (std_id) REFERENCES student (StdId) ON DELETE CASCADE);`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			log.Println("Error:", err)
		} else {
			log.Println("Table created successfully")
		}
	}
}
