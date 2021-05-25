package loginLogoutFront

import (
	"encoding/json"
	"net/http"

	"./loginLogoutObj"
	"github.com/gorilla/sessions"

	// "regexp"
	"../common"
	// "fmt"
)

var sessionName = "gsid"
var store = sessions.NewCookieStore([]byte(sessionName))

func Login(w http.ResponseWriter, r *http.Request) {
	var (
		flg       bool
		studentId string
		password  string
	)
	student := new(loginLogoutObj.Student)
	result := new(loginLogoutObj.Result)

	studentId, flg = common.CheckInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
	password, flg = common.CheckInput(r.PostFormValue("password"), `.+`)

	if flg {
		w.WriteHeader(400)
	}

	student.StudentId = studentId

	if student.Login(password) {
		store.Options = &sessions.Options{
			Domain:   "localhost",
			Path:     "/",
			MaxAge:   2592000,
			Secure:   false,
			HttpOnly: true,
		}
		session, _ := store.Get(r, sessionName)
		session.Values["login"] = true
		session.Values["studentId"] = student.StudentId
		session.Values["studentName"] = student.StudentName
		err := session.Save(r, w)
		if err != nil {
			return
		}
		result.Result = true
		result.Body = *student
	}
	json, _ := json.Marshal(result)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(json)
}
