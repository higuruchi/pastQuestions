package common

import (
	"regexp"
	"math/rand"
	"time"
	"io"
	"fmt"
	"crypto/md5"
)

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

func CheckInput(val string, check string) (ret string, err bool) {
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

func EncriptPassword(password string) (string) {
	h := md5.New()
	io.WriteString(h, password)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	salt1 := "@#$%"
	salt2 := "^&*()"
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)
	password = fmt.Sprintf("%x", h.Sum(nil))
	return password
}
