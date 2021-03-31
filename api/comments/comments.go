package comments

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

type Comment struct {
	ClassId string `json:"classId"`
	CommentId int `json:"commentId"`
	Comment string `json:"comment"`
	StudentId string `json:"studentId`
	GoodFlg bool `json:"goodflg"`
	BadFlg bool `json:"badflg"`
	Good int `json:"good"`
	Bad int `json:"bad"`
}
type Result struct {
	Result bool `json:"result"`
	Body []Comment `json:"body"`
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

func Comments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			flg := false
			// var tmp string
			comment := new(Comment)
			result := new(Result)

			comment.StudentId, flg = checkInput(r.PostFormValue("studentId"), `[0-9]{2}[A-Z][0-9]{3}`)
			comment.ClassId, flg = checkInput(r.PostFormValue("classId"), `[0-9]{7}`)
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
			comment := new(Comment)
			result := new(Result)

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


		// case "PUT":
		// case "DELETE":
	}
}

func (comment *Comment)AddComment()(result bool) {
	num := 0
	statement := `SELECT COUNT(*) FROM comments WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	err = stmt.QueryRow(comment.ClassId).Scan(&num)
	if err != nil {
		result = false
		return
	}
	if num == 0 {
		num = 1
	} else {
		statement = `SELECT commentId FROM comments WHERE classId=? ORDER BY commentId DESC LIMIT 1`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			result = false
			return
		}
		err = stmt.QueryRow(comment.ClassId).Scan(&num)
		if err != nil {
			result = false
			return
		}
		num++
	}

	statement = `INSERT INTO comments (classId, commentId, comment, studentId) VALUES (?, ?, ?, ?)`
	stmt, err = db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	_, err = stmt.Exec(comment.ClassId, num, comment.Comment, comment.StudentId)
	if err != nil {
		result = false
		return
	}

	comment.CommentId = num
	result = true
	return
}

func (comment *Comment)GetComment()(comments []Comment, result bool) {
	statement := `SELECT classId, commentId, comment, studentId, good, bad FROM comments WHERE classId=? AND commentId>=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}

	rows, _ := stmt.Query(comment.ClassId, comment.CommentId)
	for rows.Next() {
		resultComment := new(Comment)
		if err := rows.Scan(&resultComment.ClassId, &resultComment.CommentId, &resultComment.Comment, &resultComment.StudentId, &resultComment.Good, &resultComment.Bad); err != nil {
			result = false
			return
		}
		comments = append(comments, *resultComment)
	}
	result = true
	return
}