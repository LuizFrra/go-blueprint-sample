package gitlab

type GitlabGetRepoDTO struct {
	Name string `json:"name"`
	URL  string `json:"web_url"` // GitLab uses "web_url" for the repository URL
}
