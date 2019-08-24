package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

type tweet struct {
	sentiment int64 // 0 = negative, 2 = neutral, 4 = positve
	text      string
}

func main() {
	tweets := loadTweetList(10)
	fmt.Println(tweets)
}

// maxTweetsCount should be -1 if want to lose entire file
func loadTweetList(maxTweetsCount int) []tweet {
	input, err := ioutil.ReadFile("dataset/training.1600000.processed.noemoticon.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bytes.NewReader(input))
	tweets := []tweet{}
	count := 0
	for {
		record, err := r.Read()

		if maxTweetsCount != -1 && count >= maxTweetsCount {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		sentimentValue, err := strconv.ParseInt(record[0], 10, 8)
		if err != nil {
			log.Fatal(err)
		}

		tweets = append(tweets, tweet{
			sentiment: sentimentValue,
			text:      record[5],
		})

		count++
	}
	return tweets
}
