package studentsFront

import (
	"net/http"
	// "database/sql"
	"encoding/json"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	// "regexp"
	"../common"
	"./studentsObj"
	// "golang.org/x/crypto/bcrypt"
)

func Students(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	switch r.Method {
	case "POST":
		student := new(studentsObj.Student)
		result := new(studentsObj.Result)
		var (
			flg      bool
			password string
		)

		student.StudentId, flg = common.CheckInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
		student.Name, flg = common.CheckInput(r.PostFormValue("name"), ``)
		password, flg = common.CheckInput(r.PostFormValue("password"), `[0-9a-zA-Z]{4,}`)

		if flg {
			w.WriteHeader(400)
		} else {
			if student.AddStudent(password) {
				result.Result = true
				result.Body = append(result.Body, *student)
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				result.Result = false
				json, _ := json.Marshal(result)
				w.Write(json)
			}
		}

	case "DELETE":
		student := new(studentsObj.Student)
		result := new(studentsObj.Result)
		var flg bool
		parsedUri := strings.Split(r.RequestURI, "/")
		student.StudentId, flg = common.CheckInput(parsedUri[2], `[0-9]{2}[A-Z][0-9]{3}`)

		if flg {
			w.WriteHeader(400)
		} else {
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
		}
	case "PUT":
		student := new(studentsObj.Student)
		result := new(studentsObj.Result)
		var (
			flg      bool
			password string
			key      string
			val      string
		)

		parseUri := strings.Split(r.RequestURI, "/")
		student.StudentId, flg = common.CheckInput(parseUri[2], `[0-9]{2}[A-Z][0-9]{3}`)
		password, flg = common.CheckInput(r.FormValue("password"), `[0-9a-zA-Z]{4,}`)
		key, flg = common.CheckInput(parseUri[3], `.*`)
		val, flg = common.CheckInput(r.FormValue(key), `.*`)

		if flg {
			w.WriteHeader(400)
		} else {
			if student.ModifyStudent(key, val, password) {
				result.Result = true
				result.Body = append(result.Body, *student)
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				result.Result = false
				json, _ := json.Marshal(result)
				w.Write(json)
			}
		}
	case "GET":
		result := new(studentsObj.Result)
		studentId, flg := common.CheckInput(r.FormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
		if flg {
			w.WriteHeader(400)
		} else {
			// 失敗した場合のエラー処理のことなどは後回しにする
			result.ShowStudnet(studentId)
			result.Result = true
			json, _ := json.Marshal(result)
			w.Write(json)
		}

	}
}
