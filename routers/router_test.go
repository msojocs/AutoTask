package router

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/msojocs/AutoTask/v1/bootstrap"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
	"github.com/stretchr/testify/assert"
)

func init() {
	bootstrap.Init()
}

func TestRouter(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin get method", w.Body.String())
}

func TestIndexPostRouter(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin post method", w.Body.String())
}

func TestIndexPatchRouter(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin patch method", w.Body.String())
}

func TestUserSave(t *testing.T) {
	username := "lisi"
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已保存", w.Body.String())
}

func TestUserSaveByQuery(t *testing.T) {
	username := "lisi"
	age := 18
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户："+username+",年龄："+strconv.Itoa(age)+"已保存", w.Body.String())
}

func TestUserSaveByQuery2(t *testing.T) {
	username := "lisi"
	age := 20
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户："+username+",年龄："+strconv.Itoa(age)+"已保存", w.Body.String())
}

func TestUserRegister(t *testing.T) {
	value := map[string]string{
		"username": "123@gmail.com",
		"nick":     "testuser",
		"password": "123456",
	}
	data, err := json.Marshal(value)

	if err != nil {
		log.Fatalln("请求 JSON 转换失败 ", err.Error())
	}

	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	var resp serializer.Response
	err2 := json.Unmarshal(w.Body.Bytes(), &resp)

	if err2 != nil {
		log.Println("响应 JSON 转换失败 ", err2.Error())
	}

	assert.Equal(t, 0, resp.Code)
}

func TestUserLogin(t *testing.T) {

	login := map[string]string{
		"userName": "test@gmail.com",
		"Password": "123456",
	}

	data, err := json.Marshal(login)

	if err != nil {
		log.Fatalln("请求 JSON 转换失败 ", err.Error())
	}

	log.Println("JSON ", string(data))

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var resp serializer.Response
	err2 := json.Unmarshal(w.Body.Bytes(), &resp)

	if err2 != nil {
		log.Println("响应 JSON 转换失败 ", err2.Error())
	}

	assert.Equal(t, 0, resp.Code)

}
