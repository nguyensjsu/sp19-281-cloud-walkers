package main

import (
	"bytes"
	"encoding/gob"
	"github.com/allegro/bigcache"
	"time"
)

var userFollowsCache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))


func getUserFollows(userId string) ([]string){

	var ret []string

	if(userFollowsCache != nil){


		byteVal, err := userFollowsCache.Get(userId)

		if(err == nil){
			buffer := bytes.NewBuffer(byteVal)
			gob.NewDecoder(buffer).Decode(&ret)
		}
	}

	// this is where we will call out to server
	ret = []string{"Education", "Tourism"}

	if(userFollowsCache != nil){
		buffer := &bytes.Buffer{}

		gob.NewEncoder(buffer).Encode(ret)

		userFollowsCache.Set(userId, buffer.Bytes())

	}

	return ret
}