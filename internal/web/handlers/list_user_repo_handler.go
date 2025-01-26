package handlers

import (
	"net/http"
	"test-blueprint/internal/repository/service"

	"github.com/labstack/echo/v4"
)

type ListUserReposByPlatformHandler struct {
	listUserReposByPlatformService service.ListUserReposByPlatformService
}

func NewListUserReposByPlatformHandler(listUserReposByPlatformService service.ListUserReposByPlatformService) *ListUserReposByPlatformHandler {
	return &ListUserReposByPlatformHandler{listUserReposByPlatformService: listUserReposByPlatformService}
}

func (s *ListUserReposByPlatformHandler) ListUserReposHandler(c echo.Context) error {
	username := c.Param("username")
	service := c.Param("service")

	if username == "" || service == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "username and service are required"})
	}

	repos, err := s.listUserReposByPlatformService.ListUserRepos(c.Request().Context(), username, service)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, repos)
}
