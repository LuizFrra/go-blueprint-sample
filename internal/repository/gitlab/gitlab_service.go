package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test-blueprint/internal/repository/model"
)

// GitLabService implements the RepositoryService interface for GitLab.
type GitLabService struct{}

// ListUserRepos fetches the repositories for a given GitLab username.
func (g *GitLabService) ListUserRepos(ctx context.Context, username string) ([]model.Repository, error) {
	url := fmt.Sprintf("https://gitlab.com/api/v4/users/%s/projects", username)
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error fetching GitLab repositories for user %s: %v", username, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch repositories for user %s: %s", username, resp.Status)
		return nil, fmt.Errorf("failed to fetch repositories: %s", resp.Status)
	}

	var repos []GitlabGetRepoDTO
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Printf("Error decoding response for user %s: %v", username, err)
		return nil, err
	}

	// Convert DTOs to Domain Models
	var result []model.Repository
	for _, repo := range repos {
		result = append(result, model.Repository{
			Name: repo.Name,
			URL:  repo.URL,
		})
	}

	return result, nil
}
