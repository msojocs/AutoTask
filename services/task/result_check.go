package task

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
)

func init() {
	checkMap = make(map[string]func(ret Result, exp Expected) error)
	checkMap["jsonBody"] = checkJsonBody
	checkMap["textBody"] = checkTextBody
	checkMap["status"] = checkStatus
}

var checkMap map[string]func(ret Result, exp Expected) error

func checkResponse(resp Result, exp Expected) error {

	// Body/head/Status...
	if checkMap[exp.Target] != nil {
		err := checkMap[exp.Target](resp, exp)
		if err != nil {
			return err
		}
	} else {
		return errors.New("unsupported check type")
	}
	return nil
}

func checkStatus(resp Result, exp Expected) error {
	v, err := strconv.Atoi(exp.Value)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to parse Status Value, Expected:%s", exp.Value))
	}
	switch exp.Exp {
	case ">":
		if resp.Status <= v {
			return errors.New("status not matched")
		}
		break
	case "=":
		if resp.Status != v {
			return errors.New(fmt.Sprintf("Status not matched, Expected:%d, actual:%d", v, resp.Status))
		}
		break
	}
	return nil
}

func checkJsonBody(result Result, exp Expected) error {
	ret := gjson.Get(result.Body, exp.Path)
	switch exp.Exp {
	// int类型相等
	case "integerEqual":
		v, err := strconv.ParseInt(exp.Value, 10, 64)
		if err != nil {
			return errors.New("failed to parse Value format")
		}
		if ret.Int() != v {
			return errors.New(fmt.Sprintf("check failed! Expected:%s actual:%d", exp.Value, ret.Int()))
		}
		//log.Panicln("check failed!", "Expected:", Value, " actual:", ret.Int())
		break

	//	字符串要相等
	case "stringEqual":
		if ret.String() != exp.Value {
			return errors.New(fmt.Sprintf("check failed! Expected:%s actual:%s", exp.Value, ret.String()))
		}
		//log.Panicln("check failed!", "Expected:", Value, " actual:", ret.String())
		break

	//	字符串非空
	case "stringNotEmpty":
		if len(ret.String()) == 0 {
			return errors.New("check failed! string is empty")
		}
		//log.Panicln("check failed!", "Expected:", Value, " actual:", ret.String())
		break

	//	数组长度为特定值
	case "arrayLength":
		if !ret.IsArray() {
			return errors.New(fmt.Sprintf("ret is not array"))
		}
		v, err := strconv.Atoi(exp.Value)
		if err != nil {
			return errors.New("failed to parse Value format")
		}
		if len(ret.Array()) != v {
			return errors.New(fmt.Sprintf("check failed! Expected:%d actual:%d", v, len(ret.Array())))
		}
		break

	case "regex":
		_, err := regexp.MatchString(exp.Value, ret.String())
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func checkTextBody(result Result, exp Expected) error {
	ret := result.Body
	switch exp.Exp {

	//	字符串要相等
	case "=":
		if ret != exp.Value {
			return errors.New(fmt.Sprintf("check failed! Expected:%s actual:%s", exp.Value, ret))
		}
		//log.Panicln("check failed!", "Expected:", Value, " actual:", ret.String())
		break

	//	字符串不相等
	case "!=":
		if ret == exp.Value {
			return errors.New("check failed! Content is equal")
		}
		//log.Panicln("check failed!", "Expected:", Value, " actual:", ret.String())
		break

	//	字符串非空
	case "!empty":
		if len(ret) == 0 {
			return errors.New("check failed! string is empty")
		}
		//log.Panicln("check failed!", "Expected:", Value, " actual:", ret.String())
		break

	//	字符串空
	case "empty":
		if len(ret) > 0 {
			return errors.New("check failed! string is not empty")
		}
		//log.Panicln("check failed!", "Expected:", Value, " actual:", ret.String())
		break

	case "regex":
		_, err := regexp.MatchString(exp.Value, ret)
		if err != nil {
			return err
		}
		break
	}

	return nil
}
