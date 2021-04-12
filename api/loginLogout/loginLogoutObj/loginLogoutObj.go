package loginLogoutObj

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "crypto/md5"
	// "io"
	"fmt"
	"../../common"
)

type Student struct {
	StudentId string
	StudentName string
}
type Result struct {
	Result bool `json:"result"`
	Body Student `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Fumiya_0324@/pastQuestions")
	if err != nil {
		panic(err)
	}
}

func (student *Student)Login(password string)(result bool){

	statement := `SELECT CASE
					WHEN (COUNT(*) OVER()) = 0 THEN "nothing"
					ELSE name END
					FROM students
					WHERE studentId=? AND password=?`
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Printf("%v\n", err)
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