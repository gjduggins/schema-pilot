package db

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"sigs.k8s.io/controller-runtime/pkg/log"
)

// GitHubAdapter provides methods to fetch files from GitHub by tag, with caching
type GitHubAdapter struct {
	Url        string
	Owner      string
	Repo       string
	Token      string // optional GitHub token for private repos
	httpClient *http.Client
}

// NewGitHubAdapter creates a new GitHubAdapter
func NewGitHubAdapter(url string, owner, repo, token string) *GitHubAdapter {

	return &GitHubAdapter{
		Url:        url,
		Owner:      owner,
		Repo:       repo,
		Token:      token,
		httpClient: &http.Client{},
	}
}

// FetchFile fetches a file at a specific tag, using cache if available
func (g *GitHubAdapter) FetchFile(filePath, tag string) ([]byte, error) {
	log.Log.Info("Ftech File")

	// Build GitHub API URL
	url := fmt.Sprintf(
		"%s/repos/%s/%s/contents/%s?ref=%s",
		g.Url, g.Owner, g.Repo, filePath, tag,
	)
	log.Log.Info("Fetching file from GitHub", "url", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if g.Token != "" {
		req.Header.Set("Authorization", "Bearer "+g.Token)
	}

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s\n%s", resp.Status, string(body))
	}

	var fileData struct {
		Content  string `json:"content"`
		Encoding string `json:"encoding"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&fileData); err != nil {
		return nil, err
	}

	if fileData.Encoding != "base64" {
		return nil, fmt.Errorf("unexpected encoding: %s", fileData.Encoding)
	}

	data, err := base64.StdEncoding.DecodeString(fileData.Content)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// saveToCache writes data to disk
func (g *GitHubAdapter) saveToCache(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
