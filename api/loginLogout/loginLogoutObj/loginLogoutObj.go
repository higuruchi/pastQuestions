package loginLogoutObj

import (
	"database/sql"

	"../../common"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	StudentId   string `json:"studentId"`
	StudentName string `json:"studentName	"`
}
type Result struct {
	Result bool    `json:"result"`
	Body   Student `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:F_2324@a@tcp(172.28.0.2:3306)/pastQuestion")
	if err != nil {
		panic(err)
	}
}

func (student *Student) Login(password string) (result bool) {

	statement := `SELECT CASE
					WHEN (COUNT(*) OVER()) = 0 THEN "nothing"
					ELSE name END
					FROM students
					WHERE studentId=? AND password=?`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	password = common.EncriptPassword(password)
	err = stmt.QueryRow(student.StudentId, password).Scan(&student.StudentName)
	if err != nil || student.StudentName == "nothing" {
		return
	}
	result = true
	return
}
