package main

import "net/http"

//两个测试方法
func TestAloha(w http.ResponseWriter, r *http.Request) {
	respondByJSON(w, 202, "Aloha!")
}

func TestHealth(w http.ResponseWriter, r *http.Request) {
	respondByJSON(w, 201, "I AM ALIVE")
}
