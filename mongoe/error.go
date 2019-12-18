package mongoe

import (
	"fmt"
	"regexp"
)

var ErrorCodeMap map[string]string

func init() {
	ErrorCodeMap = make(map[string]string)
	ErrorCodeMap["E11000"] = "%s 数据重复，请修改后提交"
	ErrorCodeMap["E00404"] = "无相关记录"
}

func Error(err error) error {
	return MongoeError{Err: err}
}

type MongoeError struct {
	Err error
}

// error 类型需要实现 Error() 方法
func (e MongoeError) Error() string {
	errStr := e.Err.Error()
	var body []interface{}

	code := getErrorCode(errStr)
	switch code {
	case "E11000":
		body = []interface{}{getStructBody(errStr)}
	}

	if ErrorCodeMap[code] == "" {
		body = []interface{}{errStr}
	}

	return fmt.Sprintf(ErrorCodeMap[code], body...)
}

// 在错误主体中匹配错误码
func getErrorCode(str string) string {
	codeRE := regexp.MustCompile("\\w?\\d{5}")
	code := codeRE.FindString(str)

	if code == "" {
		switch str {
		case "mongo: no documents in result":
			code = "E00404"
		}
	}

	return code
}

// 在错误主体中匹配结构主体
func getStructBody(str string) string {
	bodyRE := regexp.MustCompile("{ \\w+.* }")
	body := bodyRE.FindString(str)
	return body
}
