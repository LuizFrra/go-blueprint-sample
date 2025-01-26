package github

type GithubGetRepoDTO struct {
	Name string `json:"name"`
	URL  string `json:"html_url"` // GitHub uses "html_url" for the repository URL
}
