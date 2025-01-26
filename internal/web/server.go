package web

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"test-blueprint/internal/repository/github"
	"test-blueprint/internal/repository/gitlab"
	"test-blueprint/internal/repository/model"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type RepositoryService interface {
	ListUserRepos(username string) ([]model.Repository, error)
}

type Server struct {
	port          int
	githubService RepositoryService
	gitlabService RepositoryService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:          port,
		githubService: &github.GitHubService{},
		gitlabService: &gitlab.GitLabService{},
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
