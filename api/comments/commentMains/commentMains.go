package commentMains

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	ClassId   string `json:"classId"`
	CommentId int    `json:"commentId"`
	Comment   string `json:"comment"`
	StudentId string `json:"studentId"`
	GoodFlg   bool   `json:"goodflg"`
	BadFlg    bool   `json:"badflg"`
	Good      int    `json:"good"`
	Bad       int    `json:"bad"`
}
type Result struct {
	Result bool      `json:"result"`
	Body   []Comment `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:F_2324@a@tcp(172.29.0.2:3306)/pastQuestion")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) AddComment() (result bool) {
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

func GetComment(classId string, commentId int) (comments []Comment, result bool) {
	statement := `SELECT classId, commentId, comment, studentId, good, bad FROM comments
					WHERE classId=? AND commentId>=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}

	rows, _ := stmt.Query(classId, commentId)
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

func (comment *Comment) GoodAndBad(goodOrBad bool, addOrReduce bool) (result bool) {
	num := 0

	statement := `SELECT CASE WHEN COUNT(*) = 0 THEN 0
						ELSE 1 END
					FROM comments WHERE classId=? AND commentId=?`
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
		// addの場合
		if addOrReduce {
			statement = `UPDATE comments
						SET good = (
							SELECT good
							FROM (
								SELECT good+1 AS good
								FROM comments
								WHERE classId=? AND commentId=?
							) TMP
						)
						WHERE classId=? AND commentId=?`
			stmt, err = db.Prepare(statement)
			defer stmt.Close()
			if err != nil {
				return
			}
			_, err = stmt.Exec(comment.ClassId, comment.CommentId, comment.ClassId, comment.CommentId)
			if err != nil {
				return
			}
			result = true
			return

			// reduceの場合
		} else {
			// good--
			// if good < 0 {
			// 	return
			// }
			statement = `UPDATE comments
						SET good = (
							SELECT good
							FROM (
								SELECT (CASE 
										WHEN good = 0 THEN 0
										WHEN good > 0 THEN good - 1
										ELSE good END) AS good 
								FROM comments
								WHERE classId=? AND commentId=?
							) TMP
						)
						WHERE classId=? AND commentId=?`
			stmt, err = db.Prepare(statement)
			defer stmt.Close()
			if err != nil {
				return
			}
			_, err = stmt.Exec(comment.ClassId, comment.CommentId, comment.ClassId, comment.CommentId)
			if err != nil {
				return
			}
			result = true
			return
		}

		// badの場合
	} else {

		// addの場合
		if addOrReduce {
			statement = `UPDATE comments
						SET bad = (
							SELECT bad
							FROM (
								SELECT bad+1 AS bad
								FROM comments
								WHERE classId=? AND commentId=?
							) TMP
						)
						WHERE classId=? AND commentId=?`
			stmt, err = db.Prepare(statement)
			defer stmt.Close()
			if err != nil {
				return
			}
			_, err = stmt.Exec(comment.ClassId, comment.CommentId, comment.ClassId, comment.CommentId)
			if err != nil {

				return
			}
			result = true
			return

			// reduceの場合
		} else {
			statement = `UPDATE comments
						SET bad = (
							SELECT bad
							FROM (
								SELECT (CASE 
										WHEN bad = 0 THEN 0
										WHEN bad > 0 THEN bad - 1
										ELSE bad END) AS bad 
								FROM comments
								WHERE classId=? AND commentId=?
							) TMP
						)
						WHERE classId=? AND commentId=?`
			stmt, err = db.Prepare(statement)
			defer stmt.Close()
			if err != nil {
				return
			}
			_, err = stmt.Exec(comment.ClassId, comment.CommentId, comment.ClassId, comment.CommentId)
			if err != nil {
				return
			}
			result = true
			return
		}

	}

}
