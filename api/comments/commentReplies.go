package comments

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"strconv"
	"fmt"
)

type CommentReply struct {
	ClassId string `json:"classId"`
	CommentId int `json:"commentId"`
	CommentReplyId int `json:"commentReplyId"`
	StudentId string `json:"studentId"`
	Comment string `json:"comment"`
	GoodFlg bool `json:"goodFlg"`
	BadFlg bool `json:"badFlg"`
	Good int `json:"good"`
	Bad int `json:"bad"`
}

func CommentReplies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			flg := false
			var tmp string
			commentReply := new(CommentReply)
			result := new(Result)

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
					result.Body = append(result.Bdoy, *commentReply)
				} else {
					result.Result = false
				}
			}
			json, _ := json.Marshal(result)
			w.Write(json)

	}
}

func (commentReply *CommentReply)AddReplyComments() (result bool) {
	num := 0
	statement := `SELECT CASE WHEN COUNT(*) = 0 THEN 0
						ELSE (SELECT commentReplyId
								FROM commentReplies
								WHERE classId=? AND commentId=?
								ORDER BY commentReplyId DESC LIMIT 1) END AS commentIdNum
					FROM commentReplies
					WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	err = stmt.QueryRow(commentReply.ClassId, commentReply.CommentId, commentReply.ClassId).Scan(&num)
	num++

	statement = `INSERT INTO commentReplies (classId, commentId, commentReplyId, comment, studentId) VALUES (?, ?, ?, ?, ?)`
	stmt, err = db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	_, err = stmt.Exec(commentReply.ClassId, commentReply.CommentId, num, commentReply.Comment, commentReply.StudentId)
	if err != nil {
		fmt.Printf("%v\n", err)
		result = false
		return
	}
	commentReply.CommentReplyId = num
	result = true
	return
}