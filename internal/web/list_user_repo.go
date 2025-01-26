package web

import (
	"net/http"
	"test-blueprint/internal/repository/model"

	"github.com/labstack/echo/v4"
)

func (s *Server) listUserReposHandler(c echo.Context) error {
	username := c.Param("username")
	service := c.Param("service")

	if username == "" || service == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "username and service are required"})
	}

	var repos []model.Repository
	var err error

	switch service {
	case "github":
		repos, err = s.githubService.ListUserRepos(username)
	case "gitlab":
		repos, err = s.gitlabService.ListUserRepos(username)
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "unsupported service"})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, repos)
}
