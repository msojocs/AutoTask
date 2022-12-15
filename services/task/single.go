package task

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/msojocs/AutoTask/v1/pkg/conf"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

type Task struct {
	// POST/GET/PUT/DELETE
	Method string
	// 路径
	Url string
	// 请求头
	Header map[string]string
	// 代理
	Proxy string

	// 请求体
	Body taskBody

	// 结果测试
	Expected []Expected
}

// Result HTTP响应信息
type Result struct {
	Status int               `json:"status"`
	Header map[string]string `json:"header"`
	Body   string            `json:"body"`
}

// Expected 验证数据
type Expected struct {
	Path  string
	Value string
	Vtype string
}
type taskBody struct {
	// body类型 file/string/binary/json/form
	Type string   `json:"type"`
	Data BodyData `json:"data"`
}
type BodyData struct {
	FormData []FormData `json:"formData"`
	Form     []Form     `json:"form"`
	Raw      string     `json:"raw"`
	Binary   string     `json:"binary"`
}
type FormData struct {
	Enable bool
	Type   string
	Key    string
	Value  string
}
type Form struct {
	Enable bool
	Type   string
	Key    string
	Value  string
	Desc   string
}

func init() {
}

func (task *Task) Exec() (Result, error) {
	log.Println("request start")

	var result Result
	// 请求
	result, err := request(task)
	if err != nil {
		//log.Panicln(err.Error())
		return result, err
	}

	log.Println("test")

	// 检测请求结果
	if task.Expected != nil {
		for i := range task.Expected {
			exp := task.Expected[i]
			err = checkResponse(result, exp)
			if nil != err {
				return result, err
			}
		}
	}

	log.Println("request end")
	return result, nil
}

func (task *Task) genClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 5, //超时时间
	}
	if task.Proxy != "" {
		proxy, _ := url.Parse(task.Proxy)
		tr := &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client.Transport = tr
	}

	return client
}

func (task *Task) genBody(body *taskBody) io.Reader {
	log.Println("genBody", body.Type)
	if body == nil {
		return nil
	}
	if task.Header == nil {
		task.Header = make(map[string]string)
	}

	log.Println("switch start")
	// form/string(json...)/file/binary
	switch body.Type {
	case "formData":
		log.Println("formData")
		boundary := "--------------------------462569855119802584810426"
		task.Header["Content-Type"] = "multipart/form-data; boundary=" + boundary
		dataMap := body.Data.FormData

		var fileData string
		if dataMap != nil {
			for _, formData := range dataMap {
				if len(formData.Key+formData.Value) == 0 || !formData.Enable {
					continue
				}

				file := conf.Conf.Storage.Path + "/" + formData.Value
				filename := path.Base(file)
				fileData = "--" + boundary + "\r\n"
				fileData += "Content-Disposition: form-data; name=\"" + formData.Key + "\""

				if formData.Type == "file" {
					fileData += "; filename=\"" + filename + "\"\r\n"
					fileData = fileData + "Content-Type: application/octet-stream\r\n\r\n"
					fileContent, err := os.ReadFile(file)
					if err != nil {
						continue
					}
					fileData = fileData + string(fileContent) + "\r\n"
				} else {
					//	text
					fileData += "\r\n\r\n"
					fileData += formData.Value + "\r\n"
				}
			}
			fileData += "--" + boundary + "--\r\n"
		}
		return strings.NewReader(fileData)

	case "form":
		log.Println("form")
		task.Header["Content-Type"] = "application/x-www-form-urlencoded"
		v := body.Data.Form
		ret := ""
		if v != nil {
			for _, value := range v {
				if len(value.Key+value.Value) == 0 || !value.Enable {
					continue
				}

				ret += fmt.Sprintf("%s=%s&", value.Key, value.Value)
			}
			ret = strings.TrimRight(ret, "&")
		}
		return strings.NewReader(ret)

	case "json", "text", "javascript", "html", "xml":
		log.Println("raw")
		type2header := make(map[string]string)
		type2header["text"] = "text/plain"
		type2header["javascript"] = "application/javascript"
		type2header["json"] = "application/json"
		type2header["html"] = "text/html"
		type2header["xml"] = "application/xml"
		task.Header["Content-Type"] = type2header[body.Type]
		v := body.Data.Raw
		return strings.NewReader(v)

	case "binary":
		log.Println("binary")
		// 文件路径
		s := body.Data.Binary
		storage := conf.Conf.Storage.Path + "/" + s
		d, err := os.ReadFile(storage)
		if err != nil {
			log.Println("failed to read file Data")
		}
		return bytes.NewReader(d)
	default:
		log.Println("未知类型：", body.Type)
		break
	}
	log.Println("switch end")
	return nil
}

func request(task *Task) (Result, error) {
	log.Println("request func")

	result := Result{}

	var body io.Reader
	body = task.genBody(&task.Body)

	req, err := http.NewRequest(strings.ToUpper(task.Method), task.Url, body)
	if err != nil {
		return result, err
	}

	log.Println("handle Header")
	// 请求头处理
	header := task.Header
	if header != nil {
		//	添加请求头
		for h := range header {
			req.Header.Set(h, header[h])
		}
	}

	log.Println("Exec")
	client := task.genClient()
	resp, err := client.Do(req)
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
	result.Body = retStr
	log.Println("resp:", retStr)
	log.Println(io.EOF)
	result.Status = resp.StatusCode
	err = resp.Body.Close()
	if err != nil {
		log.Println("err3", err.Error())
		return result, err
	}

	return result, nil
}
