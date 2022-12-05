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
	Status int
	Header map[string]string
	Body   string
}

// Expected 验证数据
type Expected struct {
	Path  string
	Value string
	Vtype string
}
type taskBody struct {
	// body类型 file/string/binary/json/form
	t    string
	data interface{}
}

func init() {
}

func (task *Task) exec() (Result, error) {
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
	if body == nil {
		return nil
	}
	if task.Header == nil {
		task.Header = make(map[string]string)
	}

	// form/string(json...)/file/binary
	switch body.t {
	case "string":
		task.Header["Content-Type"] = "text/plain"
		v, _ := body.data.(string)
		return strings.NewReader(v)

	case "json":
		task.Header["Content-Type"] = "application/json"
		v, _ := body.data.(string)
		return strings.NewReader(v)

	case "form":
		task.Header["Content-Type"] = "application/x-www-form-urlencoded"
		v, _ := body.data.(map[string]string)
		ret := ""
		if v != nil {
			for key := range v {
				ret += fmt.Sprintf("%s=%s&", key, v[key])
			}
			ret = strings.TrimRight(ret, "&")
		}
		return strings.NewReader(ret)

	case "binary":
		// 文件路径
		s, ok := body.data.(string)
		if !ok {
			log.Println("failed to convert Body data")
			return nil
		}
		storage := conf.Conf.Storage.Path + "/" + s
		d, err := os.ReadFile(storage)
		if err != nil {
			log.Println("failed to read file Data")
		}
		return bytes.NewReader(d)

	case "file":
		boundary := "--------------------------462569855119802584810426"
		task.Header["Content-Type"] = "multipart/form-data; boundary=" + boundary
		dataMap, ok := body.data.(map[string]string)
		if !ok {
			return nil
		}

		var fileData string
		if dataMap != nil {
			for name := range dataMap {
				file := conf.Conf.Storage.Path + "/" + dataMap[name]
				filename := path.Base(file)
				fileContent, err := os.ReadFile(file)
				if err != nil {
					continue
				}
				fileData = "--" + boundary + "\r\n"
				fileData = fileData + "Content-Disposition: form-data; name=\"" + name + "\"; filename=\"" + filename + "\"\r\n"
				fileData = fileData + "Content-Type: application/octet-stream\r\n\r\n"
				fileData = fileData + string(fileContent) + "\r\n"
			}
			fileData += "--" + boundary + "--\r\n"
		}
		return strings.NewReader(fileData)

	}
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

	log.Println("exec")
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
