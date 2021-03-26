package main

import (
	"net/http"
	"fmt"
	"pastQuestions/api/loginLogout"
	"pastQuestions/api/students"
	"pastQuestions/api/classes"
	"github.com/gorilla/sessions"
)

var sessionName = "gsid"

var store = sessions.NewCookieStore([]byte(sessionName))

func main() {
	mux := http.NewServeMux()
	
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/students/", students.Students)
	mux.HandleFunc("/classes/", classes.Classes)
	mux.HandleFunc("/login", login)
	
	server := &http.Server{
		Addr: "192.168.33.10:8080",
		Handler: mux,
	}
	fmt.Println("8080番ポートにてサーバが稼働しています")
	server.ListenAndServe()
}

func login(w http.ResponseWriter, r *http.Request) {
	studentId := r.PostFormValue("studentId")
	password := r.PostFormValue("password")
	student, _ := loginLogout.Login(studentId, password)

	if student.StudentId != ""{
		session, _ := store.Get(r, sessionName)
		session.Values["StudentId"] = student.StudentId
		session.Values["Name"] = student.Name
		session.Values["Email"] = student.Email
		err := session.Save(r, w)
		if err != nil {
			fmt.Println(w, err)
		}
		fmt.Printf("%v", session)
	} else {
		fmt.Println("bad")
	}
}