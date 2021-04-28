package questionBoardsFront

import (
	"net/http"

	"../common"

	// "database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"

	// "fmt"
	"strconv"
	"strings"

	"./questionBoardsObj"
)

func QuestionBoards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	switch r.Method {
	// /questionBoards/classId/year/studentId/
	// /questionBoardReply/classId/questionBoardId/studentId/
	case "POST":
		parsedUri := strings.Split(r.RequestURI, "/")

		if parsedUri[1] == "questionBoards" {
			questionBoard := new(questionBoardsObj.QuestionBoard)
			result := new(questionBoardsObj.Result)
			flg := false
			var tmp string

			questionBoard.ClassId, flg = common.CheckInput(parsedUri[2], `[0-9]{0,7}`)
			tmp, flg = common.CheckInput(parsedUri[3], `\d{4}`)
			questionBoard.Year, _ = strconv.Atoi(tmp)
			questionBoard.StudentId, flg = common.CheckInput(parsedUri[4], `[0-9]{2}[A-Z][0-9]{3}`)
			questionBoard.Question, flg = common.CheckInput(r.PostFormValue("question"), `.+`)

			if flg {
				w.WriteHeader(400)
			} else {
				result.Result = questionBoard.AddQuestionBoard()
				result.Body = append(result.Body, *questionBoard)
				json, _ := json.Marshal(result)
				w.Write(json)
			}
		} else if parsedUri[1] == "questionBoardsReply" {
			result := new(questionBoardsObj.Result)
			flg := false
			var tmp string
			questionBoardReply := new(questionBoardsObj.QuestoinBoardReply)

			questionBoardReply.ClassId, flg = common.CheckInput(parsedUri[2], `[0-9]{0,7}`)
			tmp, flg = common.CheckInput(parsedUri[3], `\d{1,}`)
			questionBoardReply.QuestionBoardId, _ = strconv.Atoi(tmp)
			// tmp, flg = common.CheckInput(parsedUri[4], `\d{4}`)
			// questionBoardReply.Year, _ = strconv.Atoi(tmp)
			questionBoardReply.StudentId, flg = common.CheckInput(parsedUri[4], `[0-9]{2}[A-Z][0-9]{3}`)
			questionBoardReply.Reply, flg = common.CheckInput(r.PostFormValue("reply"), `.+`)

			if flg {
				w.WriteHeader(400)
			} else {
				result.Result = questionBoardReply.AddQuestionBoardReply()
				json, _ := json.Marshal(result)
				w.Write(json)
			}
		}

	//  /questionBoard/classId/
	case "GET":
		parsedUri := strings.Split(r.RequestURI, "/")
		questionBoard := new(questionBoardsObj.QuestionBoard)
		result := new(questionBoardsObj.Result)
		flg := false

		questionBoard.ClassId, flg = common.CheckInput(parsedUri[2], `[0-9]{0,7}`)
		if flg {
			w.WriteHeader(400)
		} else {
			result.Body, result.Result = questionBoard.GetQuestionBoard()
			json, _ := json.Marshal(result)
			w.Write(json)
		}
	}
}
