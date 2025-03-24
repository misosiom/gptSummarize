package main

import (
	"fmt"
	"log"
	"os"

	"summarize/internal/fetch"
	"summarize/internal/summarize"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("使い方: summarize <URL>")
	}

	url := os.Args[1]

	text, err := fetch.FetchTextFromURL(url)
	if err != nil {
		log.Fatalf("取得エラー: %v", err)
	}

	result, err := summarize.SummarizeText(text)
	if err != nil {
		log.Fatalf("要約エラー: %v", err)
	}

	fmt.Println("----- 要約と翻訳 -----")
	fmt.Println(result)
}
