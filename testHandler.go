package main

import (
	"fmt"
	"net/http"
)

// 测试方法
func TestAloha(w http.ResponseWriter, r *http.Request) {
	respondByJSON(w, 200, struct {
		Info string
	}{
		Info: "aloha!",
	})
}

func TestErr(w http.ResponseWriter, r *http.Request) {
	respondByErr(w, 801, "Testing err")
}

func TestHealth(w http.ResponseWriter, r *http.Request) {
	respondByJSON(w, 200, "I AM ALIVE")
}

func PrintUser(w http.ResponseWriter, r *http.Request) {
	user, err := getJsonFromRequest(*r)
	if err != nil {
		fmt.Println("json convert failed")
		return
	}
	fmt.Println(user)
}

func TestSendAloha(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8086/hello/aloha")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Body)
}
