package github

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	internalHTTP "test-blueprint/internal/http"
	"test-blueprint/internal/repository/model"
)

// GitHubService implements the RepositoryService interface for GitHub.
type GitHubService struct {
	client internalHTTP.Client
}

func NewGitHubService(client internalHTTP.Client) *GitHubService {
	return &GitHubService{client: client}
}

// ListUserRepos fetches the repositories for a given GitHub username.
func (g *GitHubService) ListUserRepos(ctx context.Context, username string) ([]model.Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := g.client.Get(url)

	if err != nil {
		log.Printf("Error fetching GitHub repositories for user %s: %v", username, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch repositories for user %s: %s", username, resp.Status)
		return nil, fmt.Errorf("failed to fetch repositories: %s", resp.Status)
	}

	var repos []GithubGetRepoDTO
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
