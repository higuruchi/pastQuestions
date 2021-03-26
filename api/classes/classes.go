package classes

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	// "fmt"
	// "strings"
)

type Class struct {
	ClassId string `json:"classId"`
	ClassName string `json:"className"`
}
type Result struct {
	Result bool `json:"result"`
	Body []Class `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Fumiya_0324@/pastQuestions")
	if err != nil {
		panic(err)
	}
}

func Classes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	switch r.Method {
		case "POST":
			class := new(Class)
			result := new(Result)
			class.ClassId = r.PostFormValue("classId")
			class.ClassName = r.PostFormValue("className")

			if class.AddClass() {
				result.Result = true
				result.Body = append(result.Body, *class)
				json, _ := json.Marshal(result)
				w.Write(json)
			}
	}
}

func (class *Class)AddClass()(ret bool) {
	num := 0
	statement := `SELECT COUNT(*) FROM classes WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		ret=false
		return
	}
	err = stmt.QueryRow(class.ClassId).Scan(&num)

	if num == 0 {
		statement = `INSERT INTO classes (classId, className) VALUES (?, ?)`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			ret=false
			return
		}
		_, err = stmt.Exec(class.ClassId, class.ClassName)
		if err != nil {
			ret=false
			return
		}
		ret=true
		return
	}
	ret=false
	return
}