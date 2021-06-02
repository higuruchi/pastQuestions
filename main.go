package main

import (
	"fmt"
	"net/http"

	classesFront "./api/classes"
	commentsFront "./api/comments"
	loginLogoutFront "./api/loginLogout"
	pastQuestionsFront "./api/pastQuestions"
	questionBoardRepliesFront "./api/questionBoardReplies"
	questionBoardsFront "./api/questionBoards"
	studentsFront "./api/students"
	// "github.com/gorilla/sessions"
)

// var sessionName = "gsid"
// var store = sessions.NewCookieStore([]byte(sessionName))

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("./public"))
	pastQuestionFiles := http.FileServer(http.Dir("./pastQuestions"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/pastQuestions/", http.StripPrefix("/pastQuestions/", pastQuestionFiles))
	mux.HandleFunc("/students/", studentsFront.Students)
	mux.HandleFunc("/classes/", classesFront.Classes)
	mux.HandleFunc("/comments/main/", commentsFront.Comments)
	mux.HandleFunc("/comments/reply/", commentsFront.CommentReplies)
	mux.HandleFunc("/login", loginLogoutFront.Login)
	mux.HandleFunc("/pastQuestion/", pastQuestionsFront.PastQuestions)
	mux.HandleFunc("/questionBoards/", questionBoardsFront.QuestionBoards)
	mux.HandleFunc("/questionBoardsReply/", questionBoardsFront.QuestionBoards)
	mux.HandleFunc("/questionBoardRepliesFront/", questionBoardRepliesFront.QuestionBoardReplies)

	server := &http.Server{
		Addr:    "172.29.0.3:8080",
		Handler: mux,
	}
	fmt.Println("8080番ポートにて過去問データベースが稼働しています")
	server.ListenAndServe()
}
