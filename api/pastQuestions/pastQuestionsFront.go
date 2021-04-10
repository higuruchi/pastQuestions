package pastQuestionsFront

import (
	"net/http"
	"encoding/json"
	"./pastQuestionsObj"
	"io/ioutil"
	// _ "github.com/go-sql-driver/mysql"
	"strings"
	"strconv"
	"regexp"
	// "fmt"
	// "os"
	// "io"
	// "math/rand"
	// "time"
)

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


// ファイル名は乱数で決定
// apiフォーマット
// POST : /pastQuestion/<classId>/<year>/<semester>
// GET : /pastQuestion/<classId>/year/<semester>
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

			pastQuestion.ClassId, flg = checkInput(parsedUri[2], `[0-9]{0,7}`)
			year, _ := strconv.Atoi(parsedUri[3])
			pastQuestion.Year = year
			semester, _ := strconv.Atoi(parsedUri[4])
			pastQuestion.Semester = semester

			if flg {
				w.WriteHeader(400)
			}

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
			
			

		
		case "GET":
			var flg bool
			var tmp string
			pastQuestion := new(pastQuestionsObj.PastQuestion)
			parsedUri := strings.Split(r.RequestURI, "/")

			pastQuestion.ClassId, flg = checkInput(parsedUri[2], `[0-9]{0,7}`)
			tmp, flg = checkInput(parsedUri[3], `\d{4}`)
			pastQuestion.Year, _ = strconv.Atoi(tmp)
			tmp, flg = checkInput(parsedUri[4], `[1-4]`)
			pastQuestion.Semester, _ = strconv.Atoi(tmp)

			if flg {
				w.WriteHeader(400)
			} else {
				data, flg := pastQuestion.GetPastQuestion()

				if !flg {
					w.WriteHeader(404)
				} else {
					// fileName, _ := pastQuestion.GetPastQuestion()
					// data, _ := ioutil.ReadFile(fileName)
					// buf := make([]byte, 32 << 20)
					// data.Read(buf)
					w.Header().Set("Content-Disposition", "attachment; filename=pastQuestion.pdf")
					w.Header().Set("Content-Type", "application/pdf")
					w.Header().Set("Content-Length", string(len(data)))
					w.Write(data)

				}
			}


	}
}