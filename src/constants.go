package main

import (
	"fmt"
	"os"
)

const USER_NAME = "yumechi"
const REPO_NAME = "til"

func get_base_url() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", USER_NAME, REPO_NAME)
}

func getGitHubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}
