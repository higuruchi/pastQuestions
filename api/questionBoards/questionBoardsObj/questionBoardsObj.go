package questionBoardsObj

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	// "fmt"
	"database/sql"
)

type Result struct {
	Result bool            `json:"result"`
	Body   []QuestionBoard `json:"body"`
}

type ResultRep struct {
	Result bool                 `json:"result"`
	Body   []QuestoinBoardReply `json:"body"`
}

type QuestionBoard struct {
	ClassId            string               `json:"classId"`
	ClassName          string               `json:"className"`
	Year               int                  `json:"year"`
	QuestionBoardId    int                  `json:"questionBoardId"`
	StudentId          string               `json:"studentId"`
	Question           string               `json:"question"`
	QuestoinBoardReply []QuestoinBoardReply `json:"questionBoardReply"`
}

type QuestoinBoardReply struct {
	ClassId              string `json:"classId"`
	ClassName            string `json:"className"`
	QuestionBoardId      int    `json:"questionBoardId"`
	QuestionBoardReplyId int    `json:"questionBoardReplyId"`
	StudentId            string `json:"studentId"`
	Reply                string `json:"reply"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:F_2324@a@tcp(172.25.0.2:3306)/pastQuestion")
	if err != nil {
		panic(err)
	}
}

func AddQuestionBoard(classId string, year int, studentId string, question string) (result bool, questionBoard []QuestionBoard) {
	var (
		num              int
		newQuestionBoard QuestionBoard
	)

	statement := `SELECT CASE
					WHEN COUNT(*)=0 THEN 1
					ELSE (
						SELECT questionBoardId+1
						FROM questionBoards
						WHERE classId=?
						ORDER BY questionBoardId DESC LIMIT 1
					) END
				FROM questionBoards
				WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}
	err = stmt.QueryRow(classId, classId).Scan(&num)
	if err != nil {
		return
	}

	statement = `INSERT INTO questionBoards (classId, questionBoardId, studentId, question, year)
				VALUES (?, ?, ?, ?, ?)`
	stmt, err = db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}
	_, err = stmt.Exec(classId, num, studentId, question, year)
	if err != nil {
		fmt.Printf("%v\n%v\n", err, classId)
		return
	}
	newQuestionBoard.ClassId = classId
	newQuestionBoard.Year = year
	newQuestionBoard.QuestionBoardId = num
	newQuestionBoard.StudentId = studentId
	newQuestionBoard.Question = question
	questionBoard = append(questionBoard, newQuestionBoard)
	result = true
	return
}

func GetQuestionBoard() (resultQuestionBoard []QuestionBoard, result bool) {
	statement := `SELECT questionBoards.classId, questionBoards.year, questionBoards.questionBoardId, questionBoards.studentId, questionBoards.question, classes.className
				FROM questionBoards INNER JOIN classes ON questionBoards.classId = classes.classId
				ORDER BY questionBoardId
				LIMIT 10`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	rows, errs := stmt.Query()
	if errs != nil {
		return
	}
	for rows.Next() {
		tmpQuestionBoard := new(QuestionBoard)
		err = rows.Scan(&tmpQuestionBoard.ClassId, &tmpQuestionBoard.Year, &tmpQuestionBoard.QuestionBoardId, &tmpQuestionBoard.StudentId, &tmpQuestionBoard.Question, &tmpQuestionBoard.ClassName)
		statement = `SELECT classId, questionBoardReplyId, studentId, reply
					FROM questionBoardReplies
					WHERE questionBoardId=? AND classId=?
					ORDER BY questionBoardReplyId`
		stmt, err := db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			return
		}
		replyRows, err := stmt.Query(tmpQuestionBoard.QuestionBoardId, tmpQuestionBoard.ClassId)
		for replyRows.Next() {
			tmpQuestionBoardReply := new(QuestoinBoardReply)
			err = replyRows.Scan(&tmpQuestionBoardReply.ClassId, &tmpQuestionBoardReply.QuestionBoardReplyId, &tmpQuestionBoardReply.StudentId, &tmpQuestionBoardReply.Reply)
			if err != nil {
				return
			}
			tmpQuestionBoard.QuestoinBoardReply = append(tmpQuestionBoard.QuestoinBoardReply, *tmpQuestionBoardReply)
		}
		resultQuestionBoard = append(resultQuestionBoard, *tmpQuestionBoard)
	}
	result = true
	return
}

func GetQuestionBoardSelectedByClassId(classId string) (resultQuestionBoard []QuestionBoard, result bool) {
	statement := `SELECT classId, year, questionBoardId, studentId, question
					FROM questionBoards
					WHERE classId=?
					ORDER BY questionBoardId`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	rows, errs := stmt.Query(classId)
	if errs != nil {
		return
	}
	for rows.Next() {
		tmpQuestionBoard := new(QuestionBoard)
		err = rows.Scan(&tmpQuestionBoard.ClassId, &tmpQuestionBoard.Year, &tmpQuestionBoard.QuestionBoardId, &tmpQuestionBoard.StudentId, &tmpQuestionBoard.Question)
		statement = `SELECT classId, questionBoardReplyId, studentId, reply
						FROM questionBoardReplies
						WHERE classId=? AND questionBoardId=?
						ORDER BY questionBoardReplyId`
		stmt, err := db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			return
		}
		replyRows, err := stmt.Query(tmpQuestionBoard.ClassId, tmpQuestionBoard.QuestionBoardId)
		for replyRows.Next() {
			tmpQuestionBoardReply := new(QuestoinBoardReply)
			err = replyRows.Scan(&tmpQuestionBoardReply.ClassId, &tmpQuestionBoardReply.QuestionBoardReplyId, &tmpQuestionBoardReply.StudentId, &tmpQuestionBoardReply.Reply)
			if err != nil {
				return
			}
			tmpQuestionBoard.QuestoinBoardReply = append(tmpQuestionBoard.QuestoinBoardReply, *tmpQuestionBoardReply)
		}
		resultQuestionBoard = append(resultQuestionBoard, *tmpQuestionBoard)
	}
	result = true
	return
}

func (questionBoardReply *QuestoinBoardReply) AddQuestionBoardReply() (result bool, questionBoardReplies []QuestoinBoardReply) {
	statement := `SELECT COUNT(*)
					FROM questionBoards
					WHERE classId=?`
	var num int

	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}
	err = stmt.QueryRow(questionBoardReply.ClassId).Scan(&num)
	if num != 0 {
		statement = `SELECT CASE
						WHEN COUNT(*)=0 THEN 1
						ELSE (
							SELECT questionBoardReplyId+1
							FROM questionBoardReplies
							WHERE classId=?
							ORDER BY questionBoardReplyId DESC LIMIT 1
						) END
					FROM questionBoardReplies
					WHERE classId=?`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			return
		}
		err = stmt.QueryRow(questionBoardReply.ClassId, questionBoardReply.ClassId).Scan(&questionBoardReply.QuestionBoardReplyId)
		if err != nil {
			return
		}
		statement = `INSERT INTO questionBoardReplies
					(classId, year, questionBoardId, questionBoardReplyId, studentId, reply)
					VALUES (?, 2020, ?, ?, ?, ?)`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			return
		}
		_, err = stmt.Exec(questionBoardReply.ClassId, questionBoardReply.QuestionBoardId, questionBoardReply.QuestionBoardReplyId, questionBoardReply.StudentId, questionBoardReply.Reply)
		if err != nil {
			return
		}
		questionBoardReplies = append(questionBoardReplies, *questionBoardReply)
		result = true
		return
	}
	return
}
