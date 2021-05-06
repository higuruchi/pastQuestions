package pastQuestionsFront

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./pastQuestionsObj"

	// _ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"

	// "regexp"
	"../common"
	// "fmt"
	// "os"
	// "io"
	// "math/rand"
	// "time"
)

// ファイル名は乱数で決定
// apiフォーマット
// POST : /pastQuestion/<classId>/<year>/<semester>
// GET : /pastQuestion/<classId>/<fileId>/
// 該当するファイルが複数ある場合は、fileIdの最も大きいものを送る

func PastQuestions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		result := new(pastQuestionsObj.Result)
		pastQuestion := new(pastQuestionsObj.PastQuestion)
		parsedUri := strings.Split(r.RequestURI, "/")
		var flg bool
		err := r.ParseMultipartForm(32 << 20) // maxMemory
		if err != nil {
			return
		}

		pastQuestion.ClassId, flg = common.CheckInput(parsedUri[2], `[0-9]{0,7}`)
		year, _ := strconv.Atoi(parsedUri[3])
		pastQuestion.Year = year
		semester, _ := strconv.Atoi(parsedUri[4])
		pastQuestion.Semester = semester

		if flg {
			w.WriteHeader(400)
		} else {
			fileHeader := r.MultipartForm.File["pastQuestion"][0]
			file, err := fileHeader.Open()
			if err != nil {
				return
			}

			data, _ := ioutil.ReadAll(file)
			result.Result = pastQuestion.SavePastQuestion(data)
			result.Body = append(result.Body, *pastQuestion)
			json, _ := json.Marshal(result)
			w.Write(json)
		}

	case "GET":
		var flg bool
		var tmp string
		var classId string
		var fileId int
		// pastQuestion := new(pastQuestionsObj.PastQuestion)
		result := new(pastQuestionsObj.Result)
		parsedUri := strings.Split(r.RequestURI, "/")

		classId, flg = common.CheckInput(parsedUri[2], `[0-9]{0,7}`)
		tmp, flg = common.CheckInput(parsedUri[3], `\d+`)
		fileId, _ = strconv.Atoi(tmp)
		// tmp, flg = common.CheckInput(parsedUri[3], `\d{4}`)
		// pastQuestion.Year, _ = strconv.Atoi(tmp)
		// tmp, flg = common.CheckInput(parsedUri[4], `[1-4]`)
		// pastQuestion.Semester, _ = strconv.Atoi(tmp)

		if flg {
			w.WriteHeader(400)
		} else {
			pastQuestions, flg := pastQuestionsObj.GetPastQuestion(classId, fileId)
			fmt.Printf("%v\n", flg)

			if !flg {
				w.WriteHeader(404)
			} else {
				// fileName, _ := pastQuestion.GetPastQuestion()
				// data, _ := ioutil.ReadFile(fileName)
				// buf := make([]byte, 32 << 20)
				// data.Read(buf)
				// w.Header().Set("Content-Disposition", "attachment; filename=pastQuestion.pdf")
				// w.Header().Set("Content-Type", "application/pdf")
				// w.Header().Set("Content-Length", string(len(data)))
				// w.Write(data)
				result.Result = true
				result.Body = pastQuestions
				json, _ := json.Marshal(result)
				w.Write(json)
			}
		}

	}
}
