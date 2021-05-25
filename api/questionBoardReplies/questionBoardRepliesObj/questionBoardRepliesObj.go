package questionBoardRepliesObj

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	// "golang.org/x/crypto/bcrypt"
	// "crypto/md5"
	// "io"
)

type QuestionBoardReply struct {
	ClassId              string `json:"classId"`
	Year                 int    `json:"year"`
	QuestionBoardId      int    `json:"questionBoardId"`
	QuestionBoardReplyId int    `json:"questionBoardReplyId"`
	StudentId            string `json:"studentId"`
	Reply                string `json:"reply"`
}

type Result struct {
	Result bool                 `json:"result"`
	Body   []QuestionBoardReply `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:F_2324@a@tcp(172.25.0.2:3306)/pastQuestion")
	if err != nil {
		panic(err)
	}
}

func (questionBoardReply *QuestionBoardReply) AddQuestionBoardReply() (result bool) {
	var nextQuestionBoardReplyId int
	statement := `SELECT CASE WHEN COUNT(*) = 0 THEN 0
						ELSE (SELECT commentReplyId+1
								FROM questionBoardReplies
								WHERE classId=? AND questionBoardId=?
								ORDER BY questionBoardReplyId DESC LIMIT 1) END AS questionBoardReplyId
					FROM questionBoardReplies
					WHERE classId=? AND questionBoardId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}
	err = stmt.QueryRow(questionBoardReply.ClassId, questionBoardReply.QuestionBoardId, questionBoardReply.ClassId, questionBoardReply.QuestionBoardId).Scan(nextQuestionBoardReplyId)
	if err != nil {
		return
	}
	statement = `INSERT INTO questionBoardReplies
				(classId, year, questionBoardId, questionBoardReplyId, studentId, reply) VALUES
				(?, ?, ?, ?, ?, ?)`
	stmt, err = db.Prepare(statement)
	stmt.Close()
	if err != nil {
		return
	}
	_, err = stmt.Exec(questionBoardReply.ClassId, questionBoardReply.Year, questionBoardReply.QuestionBoardId, nextQuestionBoardReplyId, questionBoardReply.StudentId, questionBoardReply.Reply)
	if err != nil {
		return
	}
	questionBoardReply.QuestionBoardReplyId = nextQuestionBoardReplyId
	result = true
	return
}

func GetQuestionBoardReply(classId string, questionBoardId int) (questionBoardReplies []QuestionBoardReply, ok bool) {
	statement := `SELECT classId, questionBoardId, questionBoardReplyId, studentId, reply
					FROM questionBoardReplies
					WHERE claddId=? AND questionBoardId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}
	rows, _ := stmt.Query(classId, questionBoardId)
	for rows.Next() {
		questionBoardReply := new(QuestionBoardReply)
		err = rows.Scan(questionBoardReply.ClassId, questionBoardReply.QuestionBoardId, questionBoardReply.QuestionBoardReplyId, questionBoardReply.StudentId, questionBoardReply.Reply)
		if err != nil {
			return
		}
		questionBoardReplies = append(questionBoardReplies, *questionBoardReply)
	}
	ok = true
	return
}
