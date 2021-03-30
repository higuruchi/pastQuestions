package comments

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	// "fmt"
	// "strings"
	"regexp"
)

type Comment struct {
	ClassId string
	CommentId int
	Comment string
	StudentId string
	GoodFlg bool
	BadFlg bool
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
			comment := new(Comment)
			result := new(Result)

			comment.ClassId, flg = checkInput(r.PostFormValue("classId"), `[0-9]{7}`)
			comment.Comment, flg = checkInput(r.PostFormValue("comment"), `.+`)

			if flg == false {
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
		// case "GET":

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

	statement = `INSERT INTO comments (classId, commentId, comment) VALUES (?, ?, ?)`
	stmt, err = db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	_, err = stmt.Exec(comment.ClassId, num, comment.Comment)
	if err != nil {
		result = false
		return
	}

	comment.CommentId = num
	result = true
	return
}
