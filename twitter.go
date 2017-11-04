package main

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

//Followers will store pnly id and screen_name of the followers
type Followers map[string]string

//WhoUnfollowedMe is the main function
func WhoUnfollowedMe(timer time.Duration, v url.Values) {
	//store new followerslist from twitter here
	newList := make(Followers)
	//store main list of followers here
	followers := make(Followers)

	for {
		//First get fresh list of your followers
		newList = scanFollowers(v)

		//if we start the program for the first time
		//it will copy them to the main list
		if len(followers) == 0 {
			followers = newList
			log.Println("Fresh followers list copied to main one")
		}

		log.Println("Comparing cached followers list with the fresh one")

		//get userId from followers and search it
		//in the newList. If isn't there, send DM
		for user := range followers {
			_, ok := newList[user]
			if !ok {
				log.Println(followers[user], "Unfollwed you")

				sendDM(user, followers[user], v)
			}
		}

		followers = newList
		log.Printf("Job is done. Time for sleep until next scan: %v hours", timer)
		time.Sleep(timer * time.Minute)
	}

}

func scanFollowers(v url.Values) (followers Followers) {
	log.Println("scanFollowers started")
	followers = Followers{}

	pages := TwitterAPI.GetFollowersListAll(v)
	for page := range pages {
		for _, user := range page.Followers {
			followers[user.IdStr] = user.ScreenName
		}
	}

	return
}

func sendDM(id string, user string, v url.Values) {
	msg := fmt.Sprintf("Unfollowed by User %s | handle @%s", id, user)
	TwitterAPI.PostDMToScreenName(msg, v.Get("screen_name"))
	log.Println(msg)
}
