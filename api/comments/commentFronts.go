package commentsFront

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"strings"
	"strconv"
	"./commentMains"
	"./commentReplies"
	"regexp"
	// "fmt"
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

func Comments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			flg := false
			// var tmp string
			comment := new(commentMains.Comment)
			result := new(commentMains.Result)

			comment.StudentId, flg = checkInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
			comment.ClassId, flg = checkInput(r.PostFormValue("classId"), `[0-9]{7}`)
			comment.Comment, flg = checkInput(r.PostFormValue("comment"), `.+`)
			// tmp, _ = checkInput(r.PostFormValue("comment"), `.+`)
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
			var tmp string
			comment := new(commentMains.Comment)
			result := new(commentMains.Result)

			comment.ClassId, flg = checkInput(r.FormValue("classId"), `[0-9]{0,7}`)
			tmp, _ = checkInput(r.FormValue("commentId"), ``)
			comment.CommentId, _ = strconv.Atoi(tmp)

			if flg {
				result.Result = false
			} else {
				if ret, flg := comment.GetComment(); flg {
					result.Result = true
					result.Body = ret
				} else {
					result.Result = false
				}
			}

			json, _ := json.Marshal(result)
			w.Write(json)

		case "PUT":
			result := new(commentMains.Result)
			comment := new(commentMains.Comment)
			parsedUri := strings.Split(r.RequestURI, "/")
			changeCommand := parsedUri[3]
			comment.ClassId = parsedUri[5]
			comment.CommentId, _ = strconv.Atoi(parsedUri[6])

			
			switch changeCommand {
				case "good":
					addOrReduce, _ := strconv.ParseBool(parsedUri[4]) 
					if addOrReduce {
						result.Result = comment.GoodAndBad(true, true)
					} else {
						result.Result = comment.GoodAndBad(true, false)
					}
					json, _ := json.Marshal(result)
					w.Write(json)
				case "bad":
					addOrReduce, _ := strconv.ParseBool(parsedUri[4]) 
					if addOrReduce {
						result.Result = comment.GoodAndBad(false, true)
					} else {
						result.Result = comment.GoodAndBad(false, false)
					}
					json, _ := json.Marshal(result)
					w.Write(json)

			}
		// case "DELETE":
	}
}

func CommentReplies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			flg := false
			var tmp string
			commentReply := new(commentReplies.CommentReply)
			result := new(commentReplies.Result)

			commentReply.ClassId, flg = checkInput(r.PostFormValue("classId"), `[0-9]{0,7}`)
			tmp, _ = checkInput(r.PostFormValue("commentId"), `.+`)
			commentReply.CommentId, _ = strconv.Atoi(tmp)
			commentReply.Comment, flg = checkInput(r.PostFormValue("comment"), `.+`)
			commentReply.StudentId, flg = checkInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)

			if flg {
				result.Result = false

			} else {
				if commentReply.AddReplyComments() {
					result.Result = true
					result.Body = append(result.Body, *commentReply)
				} else {
					result.Result = false
				}
			}
			json, _ := json.Marshal(result)
			w.Write(json)

	}
}