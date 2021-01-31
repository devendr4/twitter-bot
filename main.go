package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-co-op/gocron"
	"os"
	"time"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error) {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	fmt.Println("ACCOUNT: \n", user)
	return client, nil
}

func postTweet(client *twitter.Client) {

	currentTime := time.Now()
	tweet, _, _ := client.Statuses.Update(".@XenoAFRO "+currentTime.Format("01-02-2006 03:04:05 pm"), nil)

	fmt.Println(tweet)
}

func main() {
	creds := Credentials{
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN_DEVENDR4"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET_DEVENDR4"),
		ConsumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
	}
	client, err := getClient(&creds)
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(10).Minutes().Do(postTweet, client)
	s1.StartBlocking()

	if err != nil {
		fmt.Println(err)
	}
}
