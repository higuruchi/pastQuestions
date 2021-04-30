package questionBoardRepliesFront

import (
	"encoding/json"
	"net/http"

	// "fmt"
	"strconv"

	"../common"
	"./questionBoardRepliesObj"
)

func QuestionBoardReplies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	switch r.Method {

	// //questionBoardRepliesFront/
	// <name="classId">
	// <name="questionBoardId">
	// <name="studentId">
	// <name="reply">
	case "POST":
		result := new(questionBoardRepliesObj.Result)
		questionBoardReply := new(questionBoardRepliesObj.QuestionBoardReply)
		var (
			flg bool
			tmp string
			err error
		)

		questionBoardReply.ClassId, flg = common.CheckInput(r.PostFormValue("classId"), `[0-9]{0,7}`)
		questionBoardReply.Year = 2020
		questionBoardReply.StudentId, flg = common.CheckInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
		tmp, flg = common.CheckInput(r.PostFormValue("questionBoardId"), `.*`)
		questionBoardReply.QuestionBoardId, err = strconv.Atoi(tmp)
		if err != nil {
			return
		}
		questionBoardReply.Reply, flg = common.CheckInput(r.PostFormValue("reply"), `.*`)

		if flg {
			w.WriteHeader(400)
		} else {
			ok := questionBoardReply.AddQuestionBoardReply()
			if ok {
				result.Body = append(result.Body, *questionBoardReply)
				json, _ := json.Marshal(result)
				w.Write(json)
			} else {
				w.WriteHeader(400)
			}
		}
	case "GET":
		result := new(questionBoardRepliesObj.Result)
		var (
			flg             bool
			classId         string
			tmp             string
			questionBoardId int
			err             error
		)
		classId, flg = common.CheckInput(r.FormValue("classId"), `[0-9]{0,7}`)
		tmp, flg = common.CheckInput(r.FormValue("questionBoardId"), `\d+`)
		questionBoardId, err = strconv.Atoi(tmp)
		if err != nil || flg {
			return
		}
		result.Body, flg = questionBoardRepliesObj.GetQuestionBoardReply(classId, questionBoardId)
		result.Result = true
		return
	}
}
