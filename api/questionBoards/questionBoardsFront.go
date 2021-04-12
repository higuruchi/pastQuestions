package questionBoardsFront

import (
	"../common"
	"net/http"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
	"strings"
	"strconv"
	"./questionBoardsObj"
)

func QuestionBoards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	// /questionBoard/classId/studentId/
	switch r.Method {
		case "POST":
			parsedUri := strings.Split(r.RequestURI, "/")
			questionBoard := new(questionBoardsObj.QuestionBoard)
			result := new(questionBoardsObj.Result);
			flg := false
			var tmp string

			questionBoard.ClassId, flg = common.CheckInput(parsedUri[2], `[0-9]{0,7}`)
			tmp, flg = common.CheckInput(parsedUri[3], `\d{4}`)
			questionBoard.Year, _ = strconv.Atoi(tmp)
			questionBoard.StudentId, flg = common.CheckInput(parsedUri[4], `[0-9]{2}[A-Z][0-9]{3}`)
			questionBoard.Question, flg = common.CheckInput(r.PostFormValue("question"), `.+`)

			fmt.Printf("%v\n", questionBoard)

			if flg {
				w.WriteHeader(400)
			} else {
				result.Result = questionBoard.AddQuestionBoard()
				result.Body = append(result.Body, *questionBoard)
				json, _ := json.Marshal(result)
				w.Write(json)
			}
			

	}
}

