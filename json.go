package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 数据转json并响应
func respondByJSON(w http.ResponseWriter, code int, input any) {
	//各类型数据转json
	data, err := json.Marshal(input)
	if err != nil {
		log.Println("json marshal faild:", input)
		return
	}
	//写入Header中Content-Type信息
	w.Header().Add("Content-Type", "application/json")
	//响应Header
	w.WriteHeader(code)
	//响应body
	w.Write(data)
}

func respondByErr(w http.ResponseWriter, code int, errInfo string) {
	respondByJSON(w, code, struct {
		Err string
	}{
		Err: errInfo,
	})
}
