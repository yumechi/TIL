package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
)


func testData() string {
	return `{
	"title": "Goから来ました",
	"body": "Goから来ました。\nテストです。"
}`
}

func postIssue() {
	req, _ := http.NewRequest(
		"POST",
		"https://api.github.com/repos/yumechi/til/issues",
		bytes.NewBuffer([]byte(testData())),
	)

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", getGitHubToken())
	req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")

	client := &http.Client{}
	resp, _ := client.Do(req)

	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(dumpResp))
	// byteArray, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(byteArray))
}

func main() {
	postIssue()
}
