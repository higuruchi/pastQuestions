package questionBoardsObj

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
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
			fmt.Printf("%v\n", err)
			return
		}
		result = true
		return
	}
	return
}