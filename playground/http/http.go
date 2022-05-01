package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)



func M() {
	resp, _ := http.Get("http://www.example.com")
	if resp != nil {
    defer resp.Body.Close()    // ← ここで nil じゃないときは閉じるようにしておく
	}

	body, _ := ioutil.ReadAll(resp.Body)
	
	fmt.Println(string(body))
}
