package main

import (
	"log"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

var (
	TwitterAPI               *anaconda.TwitterApi
	timer                    time.Duration
	twitterScreenName        string
	twitterConsumerKey       string
	twitterConsumerSecret    string
	twitterAccessToken       string
	twitterAccessTokenSecret string
)

func main() {
	//config
	timer = 888 //time in minutes
	twitterScreenName = "TWITTER_HANDLE_TO_WATCH"
	twitterConsumerKey = "TWITTER_CONSUMER_KEY"
	twitterConsumerSecret = "TWITTER_CONSUMER_SECRET"
	twitterAccessToken = "TWITTER_ACCESS_TOKEN"
	twitterAccessTokenSecret = "TWITTER_ACESS_TOKEN_SECRET"

	//api settings
	anaconda.SetConsumerKey(twitterConsumerKey)
	anaconda.SetConsumerSecret(twitterConsumerSecret)
	TwitterAPI = anaconda.NewTwitterApi(twitterAccessToken, twitterAccessTokenSecret)
	//TwitterAPI.EnableThrottling(10*time.Second, 5)

	v := url.Values{}
	v.Set("count", "200")                   //read followers on a batch of 200 people
	v.Set("screen_name", twitterScreenName) //screen name to watch

	log.Println("Who WhoUnfollowedMe started")
	log.Println("Timer set to: ", timer)

	//Main function
	WhoUnfollowedMe(timer, v)
}
