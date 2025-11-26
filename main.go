package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	// publicフォルダの中身を配信する設定
	http.Handle("/", http.FileServer(http.Dir("public/")))

	fmt.Println("Launch server...")
	// 8080番ポートでサーバーを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}
