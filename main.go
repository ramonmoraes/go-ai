package main

import (
	"fmt"

	"github.com/ramonmoraes/go-ai/nlp"
)

func main() {
	tweets := nlp.LoadTweetList(10)
	fmt.Println(tweets)
}
