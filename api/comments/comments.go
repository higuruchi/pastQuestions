package comments

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	// "fmt"
	"regexp"
	"strings"
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

		case "PUT":
			result := new(Result)
			comment := new(Comment)
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
					addOrReduce, _ := strconv.ParseBool(parsedUri[3]) 
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

func (comment *Comment)AddComment()(result bool) {
	num := 0
	statement := `SELECT COUNT(*) 
					FROM comments 
					WHERE classId=?`
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

// goodOrBad : true->good, false->bad  addOrReduce : true->add, false->reduce 
func (comment *Comment)GoodAndBad(goodOrBad bool, addOrReduce bool) (result bool) {
	num := 0

	statement := `SELECT COUNT(*) FROM comments WHERE classId=? AND commentId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	err = stmt.QueryRow(comment.ClassId, comment.CommentId).Scan(&num)
	if err != nil {
		return
	}
	if num == 0 {
		return
	}
	// goodの場合
	if goodOrBad {
		var good int
		statement = `SELECT good FROM comments WHERE classId=? AND commentId=?`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err!= nil {
			return
		}

		err = stmt.QueryRow(comment.ClassId, comment.CommentId).Scan(&good)
		if err != nil {
			return
		}
		// addの場合
		if addOrReduce {
			good++
			statement = `UPDATE comments SET good=? WHERE classId=? AND commentId=?`
			stmt, err = db.Prepare(statement)
			defer stmt.Close()
			if err!= nil {
				return
			}
			_, err = stmt.Exec(good, comment.ClassId, comment.CommentId)
			if err != nil {
				return
			}
			result = true
			return

		// reduceの場合
		} else {
			good--
			if good < 0 {
				return
			}
			statement = `UPDATE comments SET good=? WHERE classId=? AND commentId=?`
			stmt, err = db.Prepare(statement)
			defer stmt.Close()
			if err!= nil {
				return
			}
			_, err = stmt.Exec(good, comment.ClassId, comment.CommentId)
			if err != nil {
				return
			}
			result = true
			return
		}

	// badの場合
	} else {
		var bad int
		statement = `SELECT bad FROM comments WHERE classId=? AND commentId=?`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err!= nil {
			return
		}

		err = stmt.QueryRow(comment.ClassId, comment.CommentId).Scan(&bad)
		if err != nil {
			return
		}
		// addの場合
		if addOrReduce {
			bad++
			statement = `UPDATE comments SET bad=? WHERE classId=? AND commentId=?`
			stmt, err = db.Prepare(statement)
			defer stmt.Close()
			if err!= nil {
				return
			}
			_, err = stmt.Exec(bad, comment.ClassId, comment.CommentId)
			if err != nil {
				return
			}
			result = true
			return

		// reduceの場合
		} else {
			bad--
			if bad < 0 {
				return
			}
			statement = `UPDATE comments SET bad=? WHERE classId=? AND commentId=?`
			stmt, err = db.Prepare(statement)
			defer stmt.Close()
			if err!= nil {
				return
			}
			_, err = stmt.Exec(bad, comment.ClassId, comment.CommentId)
			if err != nil {
				return
			}
			result = true
			return
		}

	}

}