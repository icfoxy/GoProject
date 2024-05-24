package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")     //载入配置文件
	port := os.Getenv("PORT") //从配置文件中获取参数
	router := chi.NewRouter() //建立并初始化路由
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//往路由中加入两个路径以及对应处理方法
	router.HandleFunc("/health", TestHealth)
	router.HandleFunc("/hello/aloha", TestAloha)

	//创建服务器
	serv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Println("server starts at:", port)
	//开启服务器
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal("something went wrong with server")
	}
}
