package main

import (
	"fmt"
	"log"
	"user-backend/pkg"
	"user-backend/pkg/config"
	"user-backend/pkg/mongo"
	"user-backend/pkg/server"
)

type App struct {
	server  *server.Server
	session *mongo.Session
	config  *pkg.Config
}

func (a *App) Initialize() {
	a.config = config.GetConfig()
	var err error
	a.session, err = mongo.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	u := mongo.NewUserService(a.session.Copy(), a.config.Mongo)
	a.server = server.NewServer(u, a.config)
}

func (a *App) Run() {
	fmt.Println("Run")
	defer a.session.Close()
	a.server.Start()
}
