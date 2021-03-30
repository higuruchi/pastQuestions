package classes

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"regexp"
	"fmt"
	"strings"
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

func checkInput(val string, check string) (ret string, err bool) {
	r := regexp.MustCompile(check)
	if r.MatchString(val) {
		err = false
		ret = val
		return
	} else {
		err = true
		return
	}
}

func Classes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	switch r.Method {
		case "POST":
			flg := false
			class := new(Class)
			result := new(Result)

			class.ClassId, flg = checkInput(r.PostFormValue("classId"), `[0-9]{0,7}`)
			class.ClassName, flg = checkInput(r.PostFormValue("className"), `.+`)
			if flg {
				result.Result = false
			} else {
				if class.AddClass() {
					result.Result = true
					result.Body = append(result.Body, *class)
				} else {
					result.Result = false
				}
			}
			json, _ := json.Marshal(result)
			w.Write(json)
		case "PUT":
			flg := false
			parsedUri := strings.Split(r.RequestURI, "/")
			class := new(Class)
			result := new(Result)
			class.ClassId, flg = checkInput(parsedUri[2], `[0-9]{0,7}`)
			class.ClassName, flg = checkInput(r.FormValue("className"), `.+`)

			if flg {
				result.Result = false
			} else {
				if class.ModifyClass() {
					result.Result = true
					result.Body = append(result.Body, *class)
				} else {
					result.Result = false
				}
			}
			json, _ := json.Marshal(result)
			w.Write(json)
		case "DELETE":
			flg := false
			parseUri := strings.Split(r.RequestURI, "/")
			class := new(Class)
			result := new(Result)
			class.ClassId, flg = checkInput(parseUri[2], `[0-9]{0,7}`)

			if flg {
				result.Result = false	
			} else {
				if class.DeleteClass() {
					result.Result = true
					result.Body = append(result.Body, *class)
				} else {
					result.Result = false
				}
			}
			json, _ := json.Marshal(result)
			w.Write(json)
			
	}
}

func (class *Class)AddClass()(result bool) {
	num := 0
	statement := `SELECT COUNT(*) FROM classes WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	err = stmt.QueryRow(class.ClassId).Scan(&num)

	if num == 0 {
		statement = `INSERT INTO classes (classId, className) VALUES (?, ?)`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			result = false
			return
		}
		_, err = stmt.Exec(class.ClassId, class.ClassName)
		if err != nil {
			result = false
			return
		}
		result = true
		return
	}
	result = false
	return
}

func (class *Class)ModifyClass()(result bool) {
	num := 0
	statement := "SELECT COUNT(*) FROM classes WHERE classId=?"
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Printf("%v", err)
		result = false
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(class.ClassId).Scan(&num)
	if err != nil {
		result = false
		return
	}

	if num != 0 {
		statement = "UPDATE classes SET className=? where classId=?"
		stmt, err = db.Prepare(statement)
		if err != nil {
			result = false
			return
		}

		_, err = stmt.Exec(class.ClassName, class.ClassId)
		if err != nil {
			result = false
			return
		}

		result = true
		return
	}
	result = false
	return
}

func (class *Class)DeleteClass()(result bool){
	num := 0
	statement := "SELECT COUNT(*) FROM classes WHERE classId=?"
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}

	err = stmt.QueryRow(class.ClassId).Scan(&num)
	if err != nil {
		result = false
		return
	}

	if num == 1 {
		statement = "DELETE FROM classes WHERE classId=?"
		stmt, err := db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			result = false
			return
		}
		_, err = stmt.Exec(class.ClassId)
		if err != nil {
			result = false
			return
		}
		result = true 
		return
	}
	result = false
	return
}

// func (class *Class)GetClass()(classes []Class) {

// }