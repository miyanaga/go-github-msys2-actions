package main

import (
	"fmt"
	"log"
	"os"

	cp "github.com/otiai10/copy"
)

func List() {
	// 現在のディレクトリリストを表示
	list, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range list {
		fmt.Println(v.Name())
	}
}

func CopyFile() {
	cp.Copy("./日本語を含む.txt", "./ascii.txt")
}

func Read() {
	// ファイルの読み込み

	file, err := os.ReadFile("./日本語を含む.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(file))
}
