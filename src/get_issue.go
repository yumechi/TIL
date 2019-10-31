package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func application() {
	resp, _ := http.Get(get_base_url())
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}

func main() {
	application()
}
