package commentReplies

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	// "fmt"
	// "strings"
	"regexp"
	"strconv"
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

type Result struct {
	Result bool `json:"result"`
	Body []CommentReply `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Fumiya_0324@/pastQuestions")
	if err != nil {
		panic(err)
	}
}

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

func CommentReplies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			flg := false
			var tmp string
			commentReply := new(CommentReply)
			result := new(Result)

			commentReply.ClassId, flg = checkInput(r.PostFormValue("classId"), `[0-9]{0,7}`)
			tmp, _ = checkInput(r.PostFormValue("comment"), `.+`)
			commentReply.CommentId, _ = strconv.Atoi(tmp)

			if flg {
				result.Result = false

			} else {
				if commentReply.AddReplyComments() {
					result.Result = true
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
	statement := `SELECT COUNT(*) FROM commentReplies WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	err = stmt.QueryRow(commentReply.ClassId).Scan(&num)

	if num == 0 {
		num = 1
	} else {
		statement = `SELECT commentId FROM commentReplies WHERE classId=? AND commentId=? ORDER BY commentReplyId DESC LIMIT 1`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			result = false
			return
		}
		err = stmt.QueryRow(commentReply.ClassId, commentReply.CommentId).Scan(&num)
		if err != nil {
			result = false
			return
		}
		num++
	}

	statement = `INSERT INTO commentReplies (classId, commentId, commentReplyId, comment, studentId) VALUES (?, ?, ?, ?, ?)`
	stmt, err = db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	_, err = stmt.Exec(commentReply.ClassId, commentReply.CommentId, num, commentReply.Comment, commentReply.StudentId)
	if err != nil {
		result = false
		return
	}
	result = true
	return
}