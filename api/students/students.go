package students

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
	"strings"
	// "golang.org/x/crypto/bcrypt"
)

type Student struct {
	StudentId string `json:"studentId"`
	Name string `json:"name"`
}

type Result struct {
	Result bool `json:"result"`
	Body []Student `json:"body"`
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:Fumiya_0324@/pastQuestions")
	if err != nil {
		panic(err)
	}
}

func Students(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
		case "POST":
			student := new(Student)
			result := new(Result)
			student.StudentId = r.PostFormValue("studentId")
			student.Name = r.PostFormValue("name")
			password := r.PostFormValue("password")
			if student.AddStudent(password) {
				result.Result = true
				result.Body = append(result.Body, *student)
				json, _ := json.Marshal(result)
				w.Header().Set("Content-Type", "application/json; charset=utf8")
				w.Write(json)
			} else {
				result.Result = false
				json, _ := json.Marshal(result)
				w.Header().Set("Content-Type", "application/json; charset=utf8")
				w.Write(json)
			}
		case "DELETE":
			student := new(Student)
			result := new(Result)
			parsedUri := strings.Split(r.RequestURI, "/")
			student.StudentId = parsedUri[2]
			fmt.Printf("%v\n", student.StudentId)
			if student.DeleteStudent() {
				result.Result = true
				result.Body = append(result.Body, *student)
				json, _ := json.Marshal(result)
				w.Header().Set("Content-Type", "application/json; charset=utf8")
				w.Write(json)
			} else {
				result.Result = false
				json, _ := json.Marshal(result)
				w.Header().Set("Content-Type", "application/json; charset=utf8")
				w.Write(json)
			}
	}
}

func (student *Student)AddStudent(password string) (result bool) {
	num := 0
	result = false
	statement := `SELECT COUNT(*) FROM students WHERE studentId=?`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(student.StudentId).Scan(&num)

	if num == 0 {
		statement = `INSERT INTO students (studentId, name, password) VALUES (?, ?, ?)`
		stmt, err = Db.Prepare(statement)
		if err != nil {
			return
		}
		_, err = stmt.Exec(student.StudentId, student.Name, password)
		if err !=nil {
			return
		}
		result = true
	}
	return
}

func (student *Student)DeleteStudent() (result bool) {
	result = false
	num := 0

	statement := `SELECT studentId, name, COUNT(*) OVER() FROM students WHERE studentId=?`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(student.StudentId).Scan(&student.StudentId, &student.Name, &num)
	fmt.Printf("%v\n", num)
	if num != 0 {
		statement = `DELETE FROM students WHERE studentId=?`
		stmt, err = Db.Prepare(statement)
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

// func (student *Student)ModifyStudent() (result bool) {

// }

// func(student *Student)ShowStudnet() {

// }