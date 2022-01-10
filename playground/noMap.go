package main

import "fmt"

func main() {
	// map / 非推奨
	d := map[string]string{
		"name":  "Gopher",
		"title": "programmer",
	}

	// こう書く
	type data struct {
		name  string
		title string
	}

	d2 := data{
		name:  "Gopher",
		title: "programmer",
	}

	fmt.Println(d["name"], d["title"])
	fmt.Println(d2.name, d2.title)
}

// goではmapは非推奨
// 型があるので、型を指定して初期化する
