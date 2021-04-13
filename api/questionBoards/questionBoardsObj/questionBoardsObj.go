package questionBoardsObj

import (
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
	"database/sql"
)

type Result struct {
	Result bool `json:"result"`
	Body []QuestionBoard `json:"body"`
}

type QuestionBoard struct {
	ClassId string `json:"classId"`
	Year int `json:"year"`
	QuestionBoardId int `json: "questionBoardId"`
	StudentId string `json:"studentId"`
	Question string `json:"question"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Fumiya_0324@/pastQuestions")
	if err != nil {
		panic(err)
	}		
}


func (questionBoard *QuestionBoard)AddQuestionBoard() (result bool) {
	var num int

	statement := `SELECT CASE
					WHEN COUNT(*)=0 THEN 0
					ELSE 1 END
					FROM classes
					WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	err = stmt.QueryRow(questionBoard.ClassId).Scan(&num)
	if err!= nil {
		return
	}

	if num == 1 {
		
		statement = `SELECT CASE
						WHEN COUNT(*)=0 THEN 1
						ELSE (
							SELECT questionBoardId+1
							FROM questionBoards
							WHERE classId=?
							ORDER BY questionBoardId DESC LIMIT 1
						) END
					FROM questionBoards
					WHERE classId=?`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			return
		}
		err = stmt.QueryRow(questionBoard.ClassId, questionBoard.ClassId).Scan(&num)
		if err != nil {
			return
		}


		statement = `INSERT INTO questionBoards (classId, questionBoardId, studentId, question, year)
					VALUES (?, ?, ?, ?, 2020)`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			return
		}
		_, err = stmt.Exec(questionBoard.ClassId, num, questionBoard.StudentId, questionBoard.Question)
		if err != nil {
			return
		}
		result = true
		return
	}
	return
}

func (questionBoard *QuestionBoard)GetQuestionBoard() (resultQuestionBoard []QuestionBoard, result bool) {
	statement := `SELECT classId, year, questionBoardId, studentId, question
					FROM questionBoards
					WHERE classId=?
					ORDER BY questionBoardId`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	rows, errs := stmt.Query(questionBoard.ClassId)
	if errs!= nil {
		return
	}
	for rows.Next() {
		tmpQuestionBoard := new(QuestionBoard)
		err = rows.Scan(&tmpQuestionBoard.ClassId, &tmpQuestionBoard.Year, &tmpQuestionBoard.QuestionBoardId, &tmpQuestionBoard.StudentId, &tmpQuestionBoard.Question)
		resultQuestionBoard = append(resultQuestionBoard, *tmpQuestionBoard)
	}
	result = true
	return
}