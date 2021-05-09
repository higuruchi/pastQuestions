package commentsFront

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"./commentMains"
	"./commentReplies"
	_ "github.com/go-sql-driver/mysql"

	// "regexp"
	"../common"
	// "fmt"
)

func Comments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if flg, studentId := common.CheckLogin(r); !(flg && studentId == r.PostFormValue("studentId")) {
			w.WriteHeader(403)
			return
		}

		flg := false
		// var tmp string
		comment := new(commentMains.Comment)
		result := new(commentMains.Result)

		comment.StudentId, flg = common.CheckInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
		comment.ClassId, flg = common.CheckInput(r.PostFormValue("classId"), `[0-9]{7}`)
		comment.Comment, flg = common.CheckInput(r.PostFormValue("comment"), `.+`)
		// tmp, _ = common.CheckInput(r.PostFormValue("comment"), `.+`)
		// comment.CommentId, _ = strconv.Atoi(tmp)

		if flg {
			result.Result = false
		} else {
			if comment.AddComment() {
				result.Result = true
				result.Body = append(result.Body, *comment)
			} else {
				result.Result = false
			}
		}

		json, _ := json.Marshal(result)
		w.Write(json)
	case "GET":

		flg := false
		var classId string
		var commentId int
		var tmp string
		var err error
		result := new(commentMains.Result)

		classId, flg = common.CheckInput(r.FormValue("classId"), `[0-9]{0,7}`)
		tmp, flg = common.CheckInput(r.FormValue("commentId"), `\d+`)
		fmt.Printf("%v\n", tmp)
		commentId, err = strconv.Atoi(tmp)
		if err != nil {
			fmt.Printf("%v\n", err)
			flg = true
		}

		if flg {
			w.WriteHeader(400)
		} else {
			if ret, flg := commentMains.GetComment(classId, commentId); flg {
				result.Result = true
				result.Body = ret
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				w.WriteHeader(400)
			}
		}

	case "PUT":
		if flg, _ := common.CheckLogin(r); !flg {
			w.WriteHeader(403)
			return
		}
		result := new(commentMains.Result)
		comment := new(commentMains.Comment)
		parsedUri := strings.Split(r.RequestURI, "/")
		changeCommand := parsedUri[3]
		comment.ClassId = parsedUri[5]
		comment.CommentId, _ = strconv.Atoi(parsedUri[6])

		switch changeCommand {
		case "good":
			addOrReduce := parsedUri[4]
			if addOrReduce == "add" {
				result.Result = comment.GoodAndBad(true, true)
			} else if addOrReduce == "reduce" {
				result.Result = comment.GoodAndBad(true, false)
			}
			json, _ := json.Marshal(result)
			w.Write(json)
		case "bad":
			addOrReduce := parsedUri[4]
			if addOrReduce == "add" {
				result.Result = comment.GoodAndBad(false, true)
			} else if addOrReduce == "reduce" {
				result.Result = comment.GoodAndBad(false, false)
			}
			json, _ := json.Marshal(result)
			w.Write(json)

		}
		// case "DELETE":
	}
}

// get /comments/reply/?classId=<classId>&commentId=<commentId>
func CommentReplies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		if flg, studentId := common.CheckLogin(r); !(flg && studentId == r.PostFormValue("studentId")) {
			w.WriteHeader(403)
			return
		}

		flg := false
		var tmp string
		commentReply := new(commentReplies.CommentReply)
		result := new(commentReplies.Result)

		commentReply.ClassId, flg = common.CheckInput(r.PostFormValue("classId"), `[0-9]{0,7}`)
		tmp, _ = common.CheckInput(r.PostFormValue("commentId"), `.+`)
		commentReply.CommentId, _ = strconv.Atoi(tmp)
		commentReply.Comment, flg = common.CheckInput(r.PostFormValue("comment"), `.+`)
		commentReply.StudentId, flg = common.CheckInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)

		if flg {
			w.WriteHeader(400)
		} else {
			if commentReply.AddReplyComments() {
				result.Result = true
				result.Body = append(result.Body, *commentReply)
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				w.WriteHeader(400)
			}
		}
	case "GET":
		var (
			flg       bool
			ok        bool
			tmp       string
			classId   string
			commentId int
		)
		result := new(commentReplies.Result)
		classId, flg = common.CheckInput(r.FormValue("classId"), `[0-9]{0,7}`)
		tmp, flg = common.CheckInput(r.FormValue("commentId"), `[0-9]+`)
		commentId, _ = strconv.Atoi(tmp)

		if flg {
			w.WriteHeader(400)
		} else {
			if result.Body, ok = commentReplies.GetCommentReplies(classId, commentId); ok {
				result.Result = true
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				w.WriteHeader(400)
			}

		}
	}
}
