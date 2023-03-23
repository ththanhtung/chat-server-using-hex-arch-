package main

import (
	"lab3/internal/adapter/framework/left/udp"
	"lab3/internal/application/api"
	"lab3/internal/application/core/chat"
)

func main(){
	core := chat.NewHub()
	application := api.NewApplicaion(core)
	server := udp.NewAdapter(application)
	server.Listen(3000)
}