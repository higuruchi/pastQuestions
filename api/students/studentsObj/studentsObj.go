package studentsObj

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	// "golang.org/x/crypto/bcrypt"
	// "crypto/md5"
	// "io"
	"../../common"
)

type Student struct {
	StudentId string `json:"studentId"`
	Name      string `json:"name"`
}

type Result struct {
	Result bool      `json:"result"`
	Body   []Student `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:F_2324@a@tcp(172.28.0.2:3306)/pastQuestion")
	if err != nil {
		panic(err)
	}
}

func (student *Student) AddStudent(password string) (result bool) {
	num := 0
	result = false
	statement := `SELECT COUNT(*) FROM students WHERE studentId=?`
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(student.StudentId).Scan(&num)

	if num == 0 {
		password = common.EncriptPassword(password)
		statement = `INSERT INTO students (studentId, name, password) VALUES (?, ?, ?)`
		stmt, err = db.Prepare(statement)
		if err != nil {
			return
		}
		_, err = stmt.Exec(student.StudentId, student.Name, password)
		if err != nil {
			return
		}
		result = true
	}
	return
}

func (student *Student) DeleteStudent() (result bool) {
	result = false
	num := 0

	statement := `SELECT studentId, name, COUNT(*) OVER() FROM students WHERE studentId=?`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(student.StudentId).Scan(&student.StudentId, &student.Name, &num)
	if err != nil {
		return
	}
	if num != 0 {
		statement = `DELETE FROM students WHERE studentId=?`
		stmt, err = db.Prepare(statement)
		if err != nil {
			return
		}
		_, err = stmt.Exec(student.StudentId)
		if err != nil {
			return
		}
		result = true
	}
	return
}

func (student *Student) ModifyStudent(key string, val string, password string) (result bool) {
	result = false
	num := 0
	password = common.EncriptPassword(password)

	statement := `SELECT COUNT(*) OVER() FROM students WHERE studentId=? AND password=?`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(student.StudentId, password).Scan(&num)
	if err != nil {
		return
	}
	if num != 0 {
		switch key {
		case "eMail":
			statement = `UPDATE students SET eMail=? WHERE studentId=?`
		case "name":
			statement = `UPDATE students SET name=? WHERE studentId=?`
		case "password":
			statement = `UPDATE students SET password=? WHERE studentId=?`
		}
		stmt, err = db.Prepare(statement)
		if err != nil {
			return
		}
		_, err = stmt.Exec(val, student.StudentId)
		if err != nil {
			return
		}
		result = true
	}
	return

}

func (result *Result) ShowStudnet(studentId string) {
	student := new(Student)

	statement := `SELECT studentId, name FROM students WHERE studentId=?`
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(studentId).Scan(&student.StudentId, &student.Name)

	result.Body = append(result.Body, *student)
}
