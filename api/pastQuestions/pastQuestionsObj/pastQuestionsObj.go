package pastQuestionsObj

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	// "io"
	"os"
	// "math/rand"
	// "time"

	// "fmt"
	"../../common"
)

type PastQuestion struct {
	ClassId  string `json:"classId"`
	Year     int    `json:"year"`
	Semester int    `json:"semester"`
	FileId   int    `json:"fileId"`
	FileName string `json:"fileName"`
}
type Result struct {
	Result bool           `json:"result"`
	Body   []PastQuestion `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:F_2324@a@tcp(172.28.0.3:3306)/pastQuestion")
	if err != nil {
		panic(err)
	}
}

// ファイル名は乱数で決定
// apiフォーマット
// POST : /pastQuestion/upload/<classId>/<year>/<semester>
// GET : /pastQuestion/download/<classId>/year/<semester>
// 該当するファイルが複数ある場合は、fileIdの最も大きいものを送る

func (pastQuestion *PastQuestion) SavePastQuestion(file []byte) (result bool) {
	flg := true
	var newFileId int
	var fileName string

	for flg {
		num := 0
		fileName = common.MakeRandomStr(30)
		statement := `SELECT
						COUNT(*)
						FROM pastQuestions
						WHERE fileName=?`
		stmt, err := db.Prepare(statement)
		if err != nil {
			return
		}
		err = stmt.QueryRow(fileName).Scan(&num)
		if num == 0 {
			flg = false
		}
		if err != nil {
			return
		}
	}

	statement := `SELECT CASE
					WHEN COUNT(*)=0 THEN 1
					ELSE (
					SELECT fileId+1
					FROM pastQuestions
					WHERE classId=? AND year=? AND semester=?
					ORDER BY fileId DESC LIMIT 1) END AS fileName
				FROM pastQuestions
				WHERE classId=? AND year=? AND semester=?`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(pastQuestion.ClassId, pastQuestion.Year, pastQuestion.Semester, pastQuestion.ClassId, pastQuestion.Year, pastQuestion.Semester).Scan(&newFileId)
	if err != nil {
		return
	}
	statement = `INSERT INTO pastQuestions (classId, year, semester, fileId, fileName)
				VALUES (?, ?, ?, ?, ?)`

	stmt, err = db.Prepare(statement)
	if err != nil {
		return
	}
	_, err = stmt.Exec(pastQuestion.ClassId, pastQuestion.Year, pastQuestion.Semester, newFileId, fileName)
	if err != nil {
		return
	}

	f, _ := os.Create("./pastQuestions/" + fileName + ".pdf")
	defer f.Close()

	f.Write(file)
	pastQuestion.FileId = newFileId
	result = true
	return
}

func (pastQuestion *PastQuestion) GetPastQuestion() (fileName string, result bool) {

	statement := `SELECT CASE
					WHEN COUNT(*)=0 THEN "nothing"
					ELSE (SELECT fileName
						FROM pastQuestions
						WHERE classId=? AND year=? AND semester=?
						ORDER BY fileId DESC LIMIT 1) END AS fileName
					FROM pastQuestions
				WHERE classId=? AND year=? AND semester=?`

	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(pastQuestion.ClassId, pastQuestion.Year, pastQuestion.Semester, pastQuestion.ClassId, pastQuestion.Year, pastQuestion.Semester).Scan(&fileName)
	if fileName == "nothing" {
		return
	}
	fileName = "./pastQuestions/" + fileName + ".pdf"
	// file, _ = os.Open("./pastQuestions/"+fileName+".pdf")
	// data, _ = ioutil.ReadAll(file)
	if err != nil {
		return
	}
	result = true
	return
}
