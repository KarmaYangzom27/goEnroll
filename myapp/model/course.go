package model

import "myapp/dataStore/postgres"

type Course struct {
	CId    string `json:"cid"`
	CName  string `json:"cname"`
	Credit int    `json:"credit"`
}

const queryInsertCourse = "INSERT INTO course(cid, cname, credit) VALUES($1, $2, $3) RETURNING cid;"
const queryGetCourse = "SELECT cid, cname, credit FROM course WHERE cid=$1;"
const queryUpdateCourse = "UPDATE course SET cid=$1, cname=$2, credit=$3 WHERE cid=$4 RETURNING cid;"
const queryDeleteCourse = "DELETE FROM course WHERE cid=$1 RETURNING cid;"

func (c *Course) Create() error {
	row := postgres.Db.QueryRow(queryInsertCourse, c.CId, c.CName, c.Credit)
	return row.Scan(&c.CId)
}

func (c *Course) Read() error {
	return postgres.Db.QueryRow(queryGetCourse, c.CId).Scan(&c.CId, &c.CName, &c.Credit)
}

func (c *Course) Update(oldID string) error {
	return postgres.Db.QueryRow(queryUpdateCourse, c.CId, c.CName, c.Credit, oldID).Scan(&c.CId)
}

func (c *Course) Delete() error {
	return postgres.Db.QueryRow(queryDeleteCourse, c.CId).Scan(&c.CId)
}

func GetAllCourses() ([]Course, error) {
	rows, err := postgres.Db.Query("SELECT cid, cname, credit FROM course;")
	if err != nil {
		return nil, err
	}

	courses := []Course{}

	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.CId, &c.CName, &c.Credit)
		if dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}
