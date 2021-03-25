package students

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
	"strings"
	"regexp"
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

func Students(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	
	switch r.Method {
		case "POST":
			student := new(Student)
			result := new(Result)
			var (
				flg bool = false
				password string
			)

			student.StudentId, flg = checkInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
			student.Name, flg = checkInput(r.PostFormValue("name"), ``)
			password, flg = checkInput(r.PostFormValue("password"), ``)

			if flg {
				result.Result = false
				json, _ := json.Marshal(result)
				w.Write(json)
				return
			}
			
			if student.AddStudent(password) {
				result.Result = true
				result.Body = append(result.Body, *student)
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				result.Result = false
				json, _ := json.Marshal(result)
				// w.Header().Set("Content-Type", "application/json; charset=utf8")
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
				w.Write(json)
			} else {
				result.Result = false
				json, _ := json.Marshal(result)
				w.Write(json)
			}
		case "PUT":
			student := new(Student)
			result := new(Result)
			parseUri := strings.Split(r.RequestURI, "/")
			student.StudentId = parseUri[2]
			password := r.FormValue("password")
			key := parseUri[3]
			val := r.FormValue(key)
			if student.ModifyStudent(key, val, password) {
				result.Result = true
				result.Body = append(result.Body, *student)
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				result.Result = false
				json, _ := json.Marshal(result)
				// w.Header().Set("Content-Type", "application/json; charset=utf8")
				w.Write(json)
			}
		case "GET":
			condition := make(map[string]string)
			result := new(Result)
			query := r.URL.Query()
			
			for key, val := range query {
				condition[key] = val[0]
			}

			// 失敗した場合のエラー処理のことなどは後回しにする
			result.ShowStudnet(condition)
			result.Result = true
			json, _ := json.Marshal(result)
			w.Write(json)

	}
}

func (student *Student)AddStudent(password string) (result bool) {
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
		statement = `INSERT INTO students (studentId, name, password) VALUES (?, ?, ?)`
		stmt, err = db.Prepare(statement)
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
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Printf("%v\n", err)
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

func (student *Student)ModifyStudent(key string, val string, password string) (result bool) {
	result = false
	num := 0

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
		switch key{
			case "eMail":
				statement = `UPDATE students SET eMail=? WHERE studentId=?`
			case "name":
				statement = `UPDATE students SET name=? WHERE studentId=?`
			case "password":
				statement = `UPDATE students SET password=? WHERE studentId=?`
		}
		stmt, err = db.Prepare(statement)
		if err != nil {
			fmt.Printf("%v\n", err)

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

func (result *Result)ShowStudnet(condition map[string]string) {

	statement := `SELECT studentId, name FROM students `
	fmt.Printf("%v\n", condition)
	if len(condition) > 0 {
		keyStack := make([]string, 0)
		valStack := make([]string, 0)

		statement = statement + "WHERE "
		for key, val := range condition {
			keyStack = append(keyStack, key)
			valStack = append(valStack, val)
		}

		fmt.Printf("%v\n%v\n", keyStack, valStack)

		// SQLインジェクションの脆弱性がある
		// ブレースホルダでSQLを実行したいのだが、方法がいまいちわからない。要件等
	
		statement += keyStack[0] + " = \"" + valStack[0] + "\" "
		for i := 1; i < len(keyStack); i++ {
			statement += "AND " + keyStack[i] + "= \"" + valStack[i] + "\" "
		}

	}
	fmt.Printf("%v\n", statement)
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		student := new(Student)
		_ = rows.Scan(&student.StudentId, &student.Name)
		result.Body = append(result.Body, *student)
	}
}