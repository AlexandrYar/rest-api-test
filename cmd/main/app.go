package main

import (
	"Rest-Api-App/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Info("create router...")
	//Инициализация объекта router
	router := httprouter.New()
	logger.Fatal("create handler...")
	handler := user.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application...")
	listner, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server is listening 0.0.0.0:1234 ...")
	log.Fatal(server.Serve(listner))
}
