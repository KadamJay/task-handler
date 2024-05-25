package main

import (
	"fmt"
	"task-handler/internal/config"
	"task-handler/internal/server"
)

func main() {
	fmt.Println("Hello World!")
	cfg := config.LoadConfig("configs/config.yaml")
	s := server.NewServer(cfg)
	s.Start()
}
