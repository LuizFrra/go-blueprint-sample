package web

import (
	"net/http"
	"test-blueprint/internal/github"
	"test-blueprint/internal/gitlab"
	"test-blueprint/internal/repository/service"
	"test-blueprint/internal/web/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://*", "http://*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},

		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.GET("/health", s.healthHandler)

	githubService := github.NewGitHubService(&http.Client{})
	gitlabService := gitlab.NewGitLabService(&http.Client{})

	listUserReposByPlatformHandler := handlers.NewListUserReposByPlatformHandler(service.NewListUserReposByPlatformServiceBuilder().
		AddPlatform("github", githubService).
		AddPlatform("gitlab", gitlabService).
		Build())

	e.GET("/:service/:username/repos", listUserReposByPlatformHandler.ListUserReposHandler)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
