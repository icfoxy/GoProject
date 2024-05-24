package main

import "net/http"

//两个测试方法
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
