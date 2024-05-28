package main

import (
	"encoding/json"
	"net/http"
)

// 数据转json并响应
func respondByJSON(w http.ResponseWriter, code int, input any) error {
	//各类型数据转json
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	//写入Header中Content-Type信息
	w.Header().Add("Content-Type", "application/json")
	//响应Header
	w.WriteHeader(code)
	//响应body
	w.Write(data)
	return nil
}

// json格式报错
func respondByErr(w http.ResponseWriter, code int, errInfo string) {
	respondByJSON(w, code, struct {
		Err string
	}{
		Err: errInfo,
	})
}

// 从请求中获取json
func getJsonFromRequest(r http.Request) (interface{}, error) {
	var data interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 从响应中获取json
func getJsonFromResponse(r http.Response) (interface{}, error) {
	var data interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
