package pastQuestionsObj

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"strings"
	"strconv"
	"./commentMains"
	"./commentReplies"
	"regexp"
	"io"
	// "fmt"
)

// // ファイル名は乱数で決定
// // apiフォーマット
// // POST : /pastQuestion/upload/<classId>/<year>/<semester>
// // GET : /pastQuestion/download/<classId>/year/<semester>
// // 該当するファイルが複数ある場合は、fileIdの最も大きいものを送る


func SavePastQuestion(file []byte) (result bool) {
	fileName := MakeRandomStr(30)
	f, _ := os.Create("./pastQuestions/"+fileName+".pdf")
	defer f.Close()

	f.Write(file)
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

