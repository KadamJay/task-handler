package server

import (
	"context"
	"log"
	"net/http"
	v1 "task-handler/internal/api/v1"
	"task-handler/internal/config"
	repository "task-handler/internal/repo"
	"task-handler/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	config     *config.Config
	HTTPServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

func (s *Server) InitDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(s.config.MongoDB.URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	return client.Database("taskdb")
}

func (s *Server) Start() {
	db := s.InitDB()
	taskRepo := repository.NewTaskRepository(db.Collection("tasks"))
	taskService := service.NewTaskService(taskRepo)
	taskHandler := v1.NewTaskHandler(taskService)

	router := http.NewServeMux()
	router.HandleFunc("/tasks", taskHandler.CreateTask)
	router.HandleFunc("/task", taskHandler.GetTasks)

	s.HTTPServer = &http.Server{
		Addr:    s.config.Server.Port,
		Handler: router,
	}

	log.Printf("Starting server on %s\n", s.config.Server.Port)
	if err := s.HTTPServer.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
