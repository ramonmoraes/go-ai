package nlp

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

type Tweet struct {
	Sentiment int64 // 0 = negative, 2 = neutral, 4 = positve
	Text      string
}

// LoadTweetList should return the loaded datase
// tmaxTweetsCount should be -1 if want to lose entire file
func LoadTweetList(maxTweetsCount int) []Tweet {
	input, err := ioutil.ReadFile("dataset/training.1600000.processed.noemoticon.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bytes.NewReader(input))
	tweets := []Tweet{}
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

		tweets = append(tweets, Tweet{
			Sentiment: sentimentValue,
			Text:      record[5],
		})

		count++
	}

	return tweets
}
