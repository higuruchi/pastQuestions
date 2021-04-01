package classesFront

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"regexp"
	"strings"
	"./classesObj"
)

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
			class := new(classesObj.Class)
			result := new(classesObj.Result)

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
			class := new(classesObj.Class)
			result := new(classesObj.Result)
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
			class := new(classesObj.Class)
			result := new(classesObj.Result)
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
