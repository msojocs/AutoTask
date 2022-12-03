package task

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

func init() {
	checkMap = make(map[string]func(ret Result, exp Expected) error)
	checkMap["body"] = checkBody
	checkMap["status"] = checkStatus
}

var checkMap map[string]func(ret Result, exp Expected) error

func checkResponse(resp Result, exp Expected) error {
	idx := strings.Index(exp.path, ".")
	path := exp.path
	if idx == -1 {
		idx = len(path)
	} else {
		exp.path = path[idx+1:]
	}
	part := path[0:idx]

	// body/head/status...
	if checkMap[part] != nil {
		err := checkMap[part](resp, exp)
		if err != nil {
			return err
		}
	} else {
		return errors.New("unsupported check type")
	}
	return nil
}

func checkStatus(resp Result, exp Expected) error {
	v, err := strconv.Atoi(exp.value)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to parse status value, expected:%s", exp.value))
	}
	if resp.status != v {
		return errors.New(fmt.Sprintf("status not matched, expected:%d, actual:%d", v, resp.status))
	}
	return nil
}

func checkBody(result Result, exp Expected) error {
	ret := gjson.Get(result.body, exp.path)
	switch exp.vType {
	// int类型相等
	case "integerEqual":
		v, err := strconv.ParseInt(exp.value, 10, 64)
		if err != nil {
			return errors.New("failed to parse value format")
		}
		if ret.Int() != v {
			return errors.New(fmt.Sprintf("check failed! expected:%s actual:%d", exp.value, ret.Int()))
		}
		//log.Panicln("check failed!", "expected:", value, " actual:", ret.Int())
		break

	//	字符串要相等
	case "stringEqual":
		if ret.String() != exp.value {
			return errors.New(fmt.Sprintf("check failed! expected:%s actual:%s", exp.value, ret.String()))
		}
		//log.Panicln("check failed!", "expected:", value, " actual:", ret.String())
		break

	//	字符串非空
	case "stringNotEmpty":
		if len(ret.String()) == 0 {
			return errors.New("check failed! string is empty")
		}
		//log.Panicln("check failed!", "expected:", value, " actual:", ret.String())
		break

	//	数组长度为特定值
	case "arrayLength":
		if !ret.IsArray() {
			return errors.New(fmt.Sprintf("ret is not array"))
		}
		v, err := strconv.Atoi(exp.value)
		if err != nil {
			return errors.New("failed to parse value format")
		}
		if len(ret.Array()) != v {
			return errors.New(fmt.Sprintf("check failed! expected:%d actual:%d", v, len(ret.Array())))
		}
		break

	}
	return nil
}
