package main

import (
	"task-handler/internal/config"
	"task-handler/internal/server"
)

func main() {
	cfg := config.LoadConfig("configs/config.yaml")
	s := server.NewServer(cfg)
	s.Start()
}
