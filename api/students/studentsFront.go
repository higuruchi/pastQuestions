package studentsFront

import (
	"net/http"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
	"strings"
	// "regexp"
	"./studentsObj"
	"../common"
	// "golang.org/x/crypto/bcrypt"
)

func Students(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	
	switch r.Method {
		case "POST":
			student := new(studentsObj.Student)
			result := new(studentsObj.Result)
			var (
				flg bool = false
				password string
			)

			student.StudentId, flg = common.CheckInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
			student.Name, flg = common.CheckInput(r.PostFormValue("name"), ``)
			password, flg = common.CheckInput(r.PostFormValue("password"), `[0-9a-zA-Z]{4,}`)
			// hash,_ := bcrypt.GenerateFromPassword([]byte(password),12)
			// fmt.Printf("%v\n", hash)

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
			student := new(studentsObj.Student)
			result := new(studentsObj.Result)
			parsedUri := strings.Split(r.RequestURI, "/")
			student.StudentId = parsedUri[2]

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
			student := new(studentsObj.Student)
			result := new(studentsObj.Result)
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
			// condition := make(map[string]string)
			result := new(studentsObj.Result)
			// query := r.URL.Query()
			studentId, flg := common.CheckInput(r.FormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
			fmt.Printf("%v\n", flg)
			if flg {
				result.Result = false
				json, _ := json.Marshal(result)
				w.Write(json)
			}
			
			// for key, val := range query {
			// 	condition[key] = val[0]
			// }

			// 失敗した場合のエラー処理のことなどは後回しにする
			result.ShowStudnet(studentId)
			result.Result = true
			json, _ := json.Marshal(result)
			w.Write(json)

	}
}