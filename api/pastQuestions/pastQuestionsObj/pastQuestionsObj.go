package pastQuestionsObj

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	// "io"
	"os"
	"math/rand"
	"time"
	// "fmt"
)

type PastQuestion struct {
	ClassId string `json:"classId"`
	Year int `json:"year"`
	Semester int `json:"semester"`
	FileId int `json:"fileId"`
	FileName string `json:"fileName"`
}
type Result struct {
	Result bool `json:"result"`
	Body []PastQuestion `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Fumiya_0324@/pastQuestions")
	if err != nil {
		panic(err)
	}		
}

// ファイル名は乱数で決定
// apiフォーマット
// POST : /pastQuestion/upload/<classId>/<year>/<semester>
// GET : /pastQuestion/download/<classId>/year/<semester>
// 該当するファイルが複数ある場合は、fileIdの最も大きいものを送る


func (pastQuestion *PastQuestion)SavePastQuestion(file []byte) (result bool) {
	flg := true
	var newFileId int
	var fileName string

	for flg {
		num := 0
		fileName = MakeRandomStr(30)
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
					WHEN COUNT(*)=0 THEN 0
					ELSE  (SELECT fileId
							FROM pastQuestions
							WHERE classId=? AND year=? AND semester=?
							ORDER BY fileId DESC LIMIT 1) END AS newFileId
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
	newFileId++
	statement = `INSERT INTO pastQuestions (classId, year, semester, fileId, fileName)
				VALUES (?, ?, ?, ?, ?)`
	
	stmt, err = db.Prepare(statement)
	if err != nil {
		return
	}
	_, err = stmt.Exec(pastQuestion.ClassId, pastQuestion.Year, pastQuestion.Semester, newFileId, fileName)
	if err != nil {
		fmt.Println("6\n\n")
		return
	}

	f, _ := os.Create("./pastQuestions/"+fileName+".pdf")
	defer f.Close()

	f.Write(file)
	pastQuestion.FileId = newFileId
	result = true
	return
}

func MakeRandomStr(digit uint32) (string) {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // 乱数を生成
    b := make([]byte, digit)
	rand.Seed(time.Now().UnixNano())
    if _, err := rand.Read(b); err != nil {
        return ""
    }

    // letters からランダムに取り出して文字列を生成
    var result string
    for _, v := range b {
        // index が letters の長さに収まるように調整
        result += string(letters[int(v)%len(letters)])
    }
    return result
}

