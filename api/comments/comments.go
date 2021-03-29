import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	// "fmt"
	// "strings"
)

struct type {
	classId string
	commentId int
	comment string
	studentId string
	goodFlg bool = false
	badFlg bool = false
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
		case "GET":
		case "PUT":
		case "DELETE":
	}
}

func (comment *Comment)AddComment()(flg bool) {
	statement := `SELECT CASE WHEN commentId IS NULL THEN 0 ELSE commentId END commentId FROM comments WHERE classId="5005070" ORDER BY commentId DESC LIMIT 1`
}