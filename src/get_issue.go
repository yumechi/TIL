package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func getIssue() {
	req, _ := http.NewRequest( "GET", get_base_url(), nil )

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
	getIssue()
}
