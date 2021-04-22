package main

import (
	"net/http"
	"fmt"
	"./api/classes"
	"./api/students"
	"./api/comments"
	"./api/pastQuestions"
	"./api/loginLogout"
	"./api/questionBoards"
	// "github.com/gorilla/sessions"
)

// var sessionName = "gsid"
// var store = sessions.NewCookieStore([]byte(sessionName))

func main() {
	mux := http.NewServeMux()
	
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/students/", studentsFront.Students)
	mux.HandleFunc("/classes/", classesFront.Classes)
	mux.HandleFunc("/comments/main/", commentsFront.Comments)
	mux.HandleFunc("/comments/reply/", commentsFront.CommentReplies)
	mux.HandleFunc("/login", loginLogoutFront.Login)
	mux.HandleFunc("/pastQuestion/", pastQuestionsFront.PastQuestions)
	mux.HandleFunc("/questionBoards/", questionBoardsFront.QuestionBoards)
	mux.HandleFunc("/questionBoardsReply/", questionBoardsFront.QuestionBoards)

	
	server := &http.Server{
		Addr: "192.168.33.10:8080",
		Handler: mux,
	}
	fmt.Println("8080番ポートにて過去問データベースが稼働しています")
	server.ListenAndServe()
}

// push test second
