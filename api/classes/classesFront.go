package classesFront

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	// "regexp"
	"strings"

	"../common"
	"./classesObj"
)

func Classes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	switch r.Method {
	case "POST":
		flg := false
		class := new(classesObj.Class)
		result := new(classesObj.Result)

		class.ClassId, flg = common.CheckInput(r.PostFormValue("classId"), `[0-9]{0,7}`)
		class.ClassName, flg = common.CheckInput(r.PostFormValue("className"), `.+`)
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
		class.ClassId, flg = common.CheckInput(parsedUri[2], `[0-9]{0,7}`)
		class.ClassName, flg = common.CheckInput(r.FormValue("className"), `.+`)

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
		class.ClassId, flg = common.CheckInput(parseUri[2], `[0-9]{0,7}`)

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

	case "GET":
		result := new(classesObj.Result)
		className, flg := common.CheckInput(r.FormValue("className"), `.+`)

		if flg {
			if classes, ok := classesObj.GetHomeClass(); ok {
				result.Body = classes
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				w.WriteHeader(400)
			}
		} else {
			if classes, ok := classesObj.GetClass(className); ok {
				result.Body = classes
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				w.WriteHeader(400)
			}
		}
	}
}
