package commentMains

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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


func (comment *Comment)AddComment()(result bool) {
	num := 0
	statement := `SELECT CASE WHEN COUNT(*) = 0 THEN 0
						ELSE (SELECT commentId
								FROM comments
								WHERE classId=? 
								ORDER BY commentId DESC LIMIT 1) END AS commentIdNum
					FROM comments 
					WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	err = stmt.QueryRow(comment.ClassId, comment.ClassId).Scan(&num)
	if err != nil {
		result = false
		return
	}
	num++

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
	statement := `SELECT classId, commentId, comment, studentId, good, bad FROM comments
					WHERE classId=? AND commentId>=?`
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