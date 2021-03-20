package loginLogout

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	StudentId string
	Name string
	Email string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:Fumiya_0324@/pastQuestions")
	if err != nil {
		panic(err)
	}
}

func Login(studentId string, password string) (student *Student, err error) {
	student = &Student{}
	statement := "SELECT studentId, name, email FROM students WHERE studentId=$1 AND password=$2"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(studentId, password).Scan(&student.StudentId, &student.Name, &student.Email)
	
		
	if err != nil {
		return
	}
	return
}