package main

import (
	"log"
	"user-backend/pkg/crypto"
	"user-backend/pkg/mongo"
	"user-backend/pkg/server"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()

	h := crypto.Hash{}
	u := mongo.NewUserService(ms.Copy(), "user-backend", "user", &h)
	s := server.NewServer(u)

	s.Start()
}
