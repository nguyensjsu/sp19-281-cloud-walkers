package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/allegro/bigcache"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Topic struct {

	Label string `json:"label" bson:"label"`
}

type FollowedTopics struct{
	FollowedTopics []Topic `json:"followed_topics" bson:"followed_topics"`
}

var userActivityServerAddr = "http://34.222.10.95:3000"

var useCache = false

func initProxy(){
	useCache = os.Getenv("USE_CACHE") == "true"
	if(useCache){
		log.Println("Cache Enabled")
	}
}


// OK, this is obviously a no-no for real code.  Never put password in code
const jwtToken = "secret"

func getUserToken(w http.ResponseWriter, r *http.Request)(string, bool){
	tknStr := r.Header.Get("Authorization")

	if(len(tknStr) == 0){
		http.Error(w, "JWT User token required", http.StatusUnauthorized)
		return "", false

	}

	tokens := strings.Split(tknStr, " ")

	if(len(tokens) != 2){
		http.Error(w, "Expected type JWT not found", http.StatusUnauthorized)
		return "", false

	}

	return tokens[1], true
}

func getUserTokenFromRequest(w http.ResponseWriter, r *http.Request)(string, bool){
	// validate JWT Token

	tokenStr, ok := getUserToken(w, r)

	if(!ok){
		return "", false
	}

	tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtToken), nil
	})

	if tkn == nil || !tkn.Valid {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return "", false
	}

	claims, valid := tkn.Claims.(jwt.MapClaims)

	if err != nil || !valid{
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return "", false
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", false
	}


	//fmt.Println("claims", claims)

	return claims["id"].(string), true

}
func Ping(w http.ResponseWriter, r *http.Request) {
	/*_, ok := getUserTokenFromRequest(w, r)

	if(!ok){
		return
	}*/

	//fmt.Fprintf(w, "ping!")
	jsonVal, _ := json.Marshal("pong");
	w.Write(jsonVal)
	w.WriteHeader(200)
	//fmt.Fprintf(w, "pong!\n")
}

// FlushCache - clear any caching for specific user.  Sent anytime a user makes a change, such as following/unfollowing a topic
func FlushCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// is authorized user?
	userId, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}

	flushCacheForUser(userId)

	w.WriteHeader(http.StatusNoContent)
}

var userFollowsCache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))


func flushCacheForUser(userId string){
	userFollowsCache.Delete(userId)
}

func UserFollows(w http.ResponseWriter, r *http.Request) {
	userId, ok := getUserTokenFromRequest(w, r)

	if(!ok){
		return
	}

	userToken, ok := getUserToken(w, r)
	if(!ok){
		return
	}

	followedTopics, err := getUserFollows(userId, userToken)

	if(err != nil){
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonVal, err := json.MarshalIndent(followedTopics, "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);

}


func getUserFollows(userId string, userToken string) (FollowedTopics, error){

	var ret FollowedTopics

	if(useCache && userFollowsCache != nil){


		byteVal, err := userFollowsCache.Get(userId)

		if(err == nil){
			buffer := bytes.NewBuffer(byteVal)
			gob.NewDecoder(buffer).Decode(&ret)
			return ret, nil
		}
	}

	// this is where we will call out to server
	if(userActivityServerAddr == ""){
		userActivityServerAddr = os.Getenv("USER_ACTIVITY_SERVER")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", userActivityServerAddr + "/userFollow", nil)
	if err != nil {
		return FollowedTopics{}, err
	}

	req.Header.Add("Authorization", "JWT " + userToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if(err != nil){
		log.Println("Unexpected error from host: ", err)
		return FollowedTopics{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return FollowedTopics{}, err
	}


	err = json.Unmarshal(body, &ret)
	if err != nil {
		return FollowedTopics{}, err
	}


	if(useCache && userFollowsCache != nil){
		buffer := &bytes.Buffer{}

		gob.NewEncoder(buffer).Encode(ret)

		userFollowsCache.Set(userId, buffer.Bytes())

	}


	return ret, nil
}