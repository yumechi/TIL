package main

import "fmt"

const USER_NAME = "yumechi"
const REPO_NAME = "til"

func get_base_url() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", USER_NAME, REPO_NAME)
}
