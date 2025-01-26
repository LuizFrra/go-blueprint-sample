package gitlab_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"test-blueprint/internal/gitlab"
	"testing"
)

// MockClient is a mock implementation of the internalHTTP.Client interface.
type MockClient struct {
	Response *http.Response
	Err      error
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	return m.Response, m.Err
}

func TestListUserRepos(t *testing.T) {
	// Create a mock server to simulate GitLab API response
	repos := []gitlab.GitlabGetRepoDTO{
		{Name: "Repo1", URL: "http://example.com/repo1"},
		{Name: "Repo2", URL: "http://example.com/repo2"},
	}
	repoJSON, _ := json.Marshal(repos)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(repoJSON)
	}))
	defer ts.Close()

	// Create a mock client that uses the test server
	mockClient := &MockClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(string(repoJSON))),
		},
		Err: nil,
	}

	service := gitlab.NewGitLabService(mockClient)

	// Call the method under test
	result, err := service.ListUserRepos(context.Background(), "testuser")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validate the result
	if len(result) != 2 {
		t.Fatalf("expected 2 repositories, got %d", len(result))
	}
	if result[0].Name != "Repo1" {
		t.Fatalf("expected Repo1, got %s", result[0].Name)
	}
	if result[1].Name != "Repo2" {
		t.Fatalf("expected Repo2, got %s", result[1].Name)
	}
}
