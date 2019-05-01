package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/allegro/bigcache"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var userActivityServerAddr = "http://34.222.10.95:3000"

var userFollowsCache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))


func flushCacheForUser(userId string){
	userFollowsCache.Delete(userId)
}

func getUserFollows(userId string, userToken string) ([]string, error){

	var ret []string

	if(userFollowsCache != nil){


		byteVal, err := userFollowsCache.Get(userId)

		if(err == nil){
			buffer := bytes.NewBuffer(byteVal)
			gob.NewDecoder(buffer).Decode(&ret)
		}
	}

	// this is where we will call out to server
	ret = []string{}

	if(userActivityServerAddr == ""){
		userActivityServerAddr = os.Getenv("USER_ACTIVITY_SERVER")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", userActivityServerAddr + "/userFollow", nil)
	if err != nil {
		return []string{}, err
	}

	req.Header.Add("Authorization", "JWT " + userToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if(err != nil){
		log.Println("Unexpected error from host: ", err)
		return []string{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return []string{}, err
	}

	var followedTopics FollowedTopics

	err = json.Unmarshal(body, &followedTopics)
	if err != nil {
		return []string{}, err
	}

	for _, topic := range followedTopics.FollowedTopics {
		ret = append(ret, topic.Label)
	}


	if(userFollowsCache != nil){
		buffer := &bytes.Buffer{}

		gob.NewEncoder(buffer).Encode(ret)

		userFollowsCache.Set(userId, buffer.Bytes())

	}


	return ret, nil
}