package loginLogoutObj

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"crypto/md5"
	"io"
	"fmt"
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

func encriptPassword(password string) (string) {
	h := md5.New()
	io.WriteString(h, password)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	salt1 := "@#$%"
	salt2 := "^&*()"
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)
	password = fmt.Sprintf("%x", h.Sum(nil))
	return password
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
	password = encriptPassword(password)
	err = stmt.QueryRow(student.StudentId, password).Scan(&student.StudentName)
	if err != nil || student.StudentName == "nothing" {

		return
	}
	result = true
	return
}