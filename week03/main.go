package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Week 03: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 03 課題")

	// 以下に実装してください
	http.HandleFunc("/uranai", uranai)

	http.ListenAndServe(":8080", nil)
}

func uranai(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	d := rand.New(rand.NewSource(seed))

	luck := d.Int31n(4)

	if luck == 0 {
		fmt.Fprintln(w, "今の運勢は大吉です。")
	} else if luck == 1 {
		fmt.Fprintln(w, "今の運勢は中吉です。")
	} else if luck == 2 {
		fmt.Fprintln(w, "今の運勢は吉です。")
	} else {
		fmt.Fprintln(w, "今の運勢は凶です。")
	}
}
