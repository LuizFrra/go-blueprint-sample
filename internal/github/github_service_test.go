package github_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"test-blueprint/internal/github"
)

// MockHTTPClient is a mock implementation of the HTTPClient interface.
type MockHTTPClient struct {
	Response *http.Response
	Error    error
}

// Implement the Get method
func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.Response, m.Error
}

// TestListUserRepos tests the ListUserRepos method of GitHubService.
func TestListUserRepos(t *testing.T) {
	// Define the mock response
	repos := []github.GithubGetRepoDTO{
		{Name: "repo1", URL: "http://github.com/user/repo1"},
		{Name: "repo2", URL: "http://github.com/user/repo2"},
	}
	repoJSON, _ := json.Marshal(repos)

	// Create a mock HTTP client that returns the mock response
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(repoJSON)),
		},
		Error: nil,
	}

	// Create the GitHubService with the mock client
	service := github.NewGitHubService(mockClient)

	// Call the method under test
	result, err := service.ListUserRepos(context.Background(), "user")

	// Assertions
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("expected 2 repositories, got %d", len(result))
	}

	if result[0].Name != "repo1" || result[0].URL != "http://github.com/user/repo1" {
		t.Errorf("expected first repo to be repo1, got %+v", result[0])
	}

	if result[1].Name != "repo2" || result[1].URL != "http://github.com/user/repo2" {
		t.Errorf("expected second repo to be repo2, got %+v", result[1])
	}
}

// TestListUserRepos_Error tests the ListUserRepos method when an error occurs.
func TestListUserRepos_Error(t *testing.T) {
	// Create a mock HTTP client that returns an error
	mockClient := &MockHTTPClient{
		Response: nil,
		Error:    fmt.Errorf("network error"),
	}

	// Create the GitHubService with the mock client
	service := github.NewGitHubService(mockClient)

	// Call the method under test
	result, err := service.ListUserRepos(context.Background(), "user")

	// Assertions
	if err == nil {
		t.Fatalf("expected an error, got none")
	}

	if result != nil {
		t.Fatalf("expected no repositories, got %+v", result)
	}
}
