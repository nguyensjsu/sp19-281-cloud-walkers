package main

import (
	"gopkg.in/mgo.v2/bson"
)

const(
	UserActivityPostMessage = 1
	UserActivityPutMessage = 2
	UserActivityPostAnswer = 3
	UserActivityPutAnswer = 4
	UserActivityPostComment = 5
	UserActivityPostCommentReply = 6
	UserActivityPutComment = 7
)

type UserActivity struct {
	UserActivity int
	UserToken string
	ObjectId bson.ObjectId
}

var userActivityChan chan UserActivity

/**
thread to keep topics in sync with question topics (add, but never delete)
*/
func notifyUserActivity(topicChan <-chan Topic){

	for activity := range userActivityChan {
		switch activity.UserActivity {
		case UserActivityPostMessage:
		case UserActivityPutMessage:
		case UserActivityPostAnswer:
		case UserActivityPutAnswer:
		case UserActivityPostComment:
		case UserActivityPostCommentReply:
		case UserActivityPutComment:
		default:
			continue

		}

	}
}

