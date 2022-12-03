package task

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type Task struct {
	// POST/GET/PUT/DELETE
	method string
	// 路径
	url string
	// 请求头
	header map[string]string

	// 请求体
	body taskBody

	//	结果测试
	expected []Expected
}

type Result struct {
	status int
	header map[string]string
	body   string
}

type Expected struct {
	path  string
	value string
	vType string
}
type taskBody struct {
	// 类型
	t    string
	data interface{}
}

var requestMap map[string]func(task Task) (Result, error)

func init() {
	// 初始化
	requestMap = make(map[string]func(task Task) (Result, error))
	//requestMap["POST"] = POST
	requestMap["GET"] = GET
	//requestMap["DELETE"] = DELETE
	//requestMap["PUT"] = PUT
	//requestMap["PATCH"] = PATCH
}
func (task Task) exec() (Result, error) {
	log.Println("request start")

	var result Result
	// 转换大写
	funcName := strings.ToUpper(task.method)
	if requestMap[funcName] != nil {
		// 请求
		result, err := requestMap[funcName](task)
		if err != nil {
			//log.Panicln(err.Error())
			return result, err
		}

		log.Println("test")

		// 检测请求结果
		if task.expected != nil {
			for i := range task.expected {
				exp := task.expected[i]
				err = checkResponse(result, exp)
				if nil != err {
					return result, err
				}
			}
		}
	} else {
		log.Println("func not found!")
		return result, errors.New("method not found")
	}

	log.SetPrefix("")
	log.Println("request end")
	return result, nil
}

func GET(task Task) (Result, error) {
	log.SetPrefix("GET:")
	log.Println("GET func")

	result := Result{}

	req, err := http.NewRequest("GET", task.url, nil)
	if err != nil {
		return result, err
	}

	log.Println("handle header")
	// 请求头处理
	header := task.header
	if header != nil {
		//	添加请求头
		for h := range header {
			req.Header.Set(h, header[h])
		}
	}

	log.Println("exec")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return result, err
	}

	log.Println("read")
	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("err2", err.Error())
		return result, err
	}

	log.Println("to string")
	retStr := string(respByte)
	result.body = retStr
	log.Println("resp:", retStr)
	log.Println(io.EOF)
	result.status = resp.StatusCode
	err = resp.Body.Close()
	if err != nil {
		log.Println("err3", err.Error())
		return result, err
	}

	return result, nil
}

func POST(task Task) (Result, error) {
	log.Println("POST func")
	result := Result{}
	return result, nil
}

func PUT(task Task) (Result, error) {
	log.Println("PUT func")
	result := Result{}
	return result, nil
}

func DELETE(task Task) (Result, error) {
	log.Println("DELETE func")
	result := Result{}
	return result, nil
}

func PATCH(task Task) (Result, error) {
	log.Println("PATCH func")
	result := Result{}
	return result, nil
}
