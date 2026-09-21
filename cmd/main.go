package main

import (
	"fmt"
	"log"
	server "ypp"
	"ypp/config"
	"ypp/internal/repository"
	"ypp/internal/repository/postgres"
	"ypp/internal/service"
	handlers "ypp/internal/transport/http"
)

func main() {
	fmt.Println("Starts server running")
	
	//Init config
	conf, err := config.NewConfig()
	if err != nil{
		log.Fatal("failed to read config", err)
	}
	
	//Init handler, server, repository, service
	db, err := postgres.NewDB(postgres.Config(conf.DB))
	if err != nil{
		log.Fatal(err)
	}
	repo := repository.New(db)
	servic := service.NewService(repo)
	handler := handlers.Handler{Service: servic}
	serv := server.Server{HttpServer: handler}
	//Run server 
	log.Println("Serv runnig")
	if err := serv.Run(conf.Port); err != nil{
		log.Fatal("Cant run server", err)
	}
	db.Close()
}
