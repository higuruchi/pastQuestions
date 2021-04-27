package commentReplies

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type CommentReply struct {
	ClassId        string `json:"classId"`
	CommentId      int    `json:"commentId"`
	CommentReplyId int    `json:"commentReplyId"`
	StudentId      string `json:"studentId"`
	Comment        string `json:"comment"`
	GoodFlg        bool   `json:"goodFlg"`
	BadFlg         bool   `json:"badFlg"`
	Good           int    `json:"good"`
	Bad            int    `json:"bad"`
}

type Result struct {
	Result bool           `json:"result"`
	Body   []CommentReply `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:F_2324@a@tcp(172.28.0.2:3306)/pastQuestion")
	if err != nil {
		panic(err)
	}
}

func (commentReply *CommentReply) AddReplyComments() (result bool) {
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
