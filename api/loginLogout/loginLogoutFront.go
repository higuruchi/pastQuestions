package loginLogoutFront

import (
	"net/http"
	"github.com/gorilla/sessions"
	"./loginLogoutObj"
	"encoding/json"
	"regexp"
	// "fmt"
)

var sessionName = "gsid"
var store = sessions.NewCookieStore([]byte(sessionName))

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


func Login(w http.ResponseWriter, r *http.Request) {
	var (
		flg bool
		studentId string
		password string
	)
	student := new(loginLogoutObj.Student)
	result := new(loginLogoutObj.Result)

	studentId, flg = checkInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
	password, flg = checkInput(r.PostFormValue("password"), `.+`)

	if flg {
		json, _ := json.Marshal(result)
		w.Write(json)
	}

	student.StudentId = studentId
	
	if student.Login(password) {

		session, _ := store.Get(r, sessionName)
		session.Values["login"] = true
		session.Values["StudentId"] = student.StudentId
		session.Values["studentName"] = student.StudentName
		err := session.Save(r, w)
		if err != nil {
			return
		}
		result.Result = true
		result.Body = *student
	}
	json, _ := json.Marshal(result)
	w.Write(json)
}