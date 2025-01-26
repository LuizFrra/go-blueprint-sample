package github_test

import (
	"bytes"
	"context"
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

// TestListUserRepos tests the ListUserRepos method of GitHubService with various scenarios.
func TestListUserRepos(t *testing.T) {
	// Define test cases
	tests := []struct {
		name          string
		repos         []github.GithubGetRepoDTO
		expected      int
		mockResp      *http.Response
		expectError   bool
		expectedError string // New field for expected error message
	}{
		{
			name: "valid response with two repos",
			repos: []github.GithubGetRepoDTO{
				{Name: "repo1", URL: "http://github.com/user/repo1"},
				{Name: "repo2", URL: "http://github.com/user/repo2"},
			},
			expected:      2,
			mockResp:      &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(bytes.NewReader([]byte(`[{"name": "repo1", "url": "http://github.com/user/repo1"},{"name": "repo2", "url": "http://github.com/user/repo2"}]`)))},
			expectError:   false,
			expectedError: "",
		},
		{
			name:          "valid response with no repos",
			repos:         []github.GithubGetRepoDTO{},
			expected:      0,
			mockResp:      &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(bytes.NewReader([]byte(`[]`)))},
			expectError:   false,
			expectedError: "",
		},
		{
			name:          "internal server error",
			expected:      0,
			mockResp:      &http.Response{StatusCode: http.StatusInternalServerError, Status: "500 Internal Server Error", Body: ioutil.NopCloser(bytes.NewReader([]byte("")))},
			expectError:   true,
			expectedError: "failed to fetch repositories: 500 Internal Server Error", // Expected error message
		},
		{
			name:     "decode error",
			expected: 0,
			mockResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("invalid json"))),
			},
			expectError:   true,
			expectedError: "invalid character 'i' looking for beginning of value", // Expected error message
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client that returns the mock response or error
			mockClient := &MockHTTPClient{
				Response: tt.mockResp,
				Error:    nil, // No need to set Error here since we're using mockResp
			}

			// Create the GitHubService with the mock client
			service := github.NewGitHubService(mockClient)

			// Call the method under test
			result, err := service.ListUserRepos(context.Background(), "user")

			// Assertions
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error, got none")
				}
				if err.Error() != tt.expectedError {
					t.Fatalf("expected error message '%s', got '%v'", tt.expectedError, err)
				}
				if result != nil {
					t.Fatalf("expected no repositories, got %+v", result)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				if len(result) != tt.expected {
					t.Fatalf("expected %d repositories, got %d", tt.expected, len(result))
				}
			}
		})
	}
}
