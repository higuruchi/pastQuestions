package students

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
	// "golang.org/x/crypto/bcrypt"
)

type Student struct {
	StudentId string
	Name string
	Email string
	Password string
}

type Result struct {
	Result bool `json:"result"`
	StudentId string `json:"studentId"`
	Name string `json:"name"`
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
			student.StudentId = r.PostFormValue("studentId")
			student.Name = r.PostFormValue("name")
			student.Password = r.PostFormValue("password")
			if student.AddStudent() {
				fmt.Println("hello wold bad !!!")
				result := new(Result)
				result.Result = true
				result.StudentId = student.StudentId
				result.Name = student.Name
				json, _ := json.Marshal(result)
				w.Header().Set("Content-Type", "application/json; charset=utf8")
				w.Write(json)
			}
		// case "DELETE":

	}
}

func (student *Student)AddStudent() (result bool) {
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
		_, err = stmt.Exec(student.StudentId, student.Name, student.Password)
		if err !=nil {
			fmt.Printf("%v", err)
			return
		}
		result = true
	}
	return
}

// func (student *Student)DeleteStudent() (result bool) {

// }

// func (student *Student)ModifyStudent() (result bool) {

// }

// func(student *Student)ShowStudnet() {

// }