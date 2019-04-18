function getNextId(){
    return ObjectId()
}


userIds = [
    "1000000",
    "1231000",
    "4200304",
    "1466220",
    "9822142",
    "3112225",
    "6213498",
    "7855542"
]

answerIds = [
                ObjectId("5cb572fb78163fa3c9727019"),
                ObjectId("5cb572fb78163fa3c972701a"),
                ObjectId("5cb572fb78163fa3c972701b"),
                ObjectId("5cb572fb78163fa3c972701c"),
                ObjectId("5cb572fb78163fa3c972701d"),
                ObjectId("5cb572fb78163fa3c972701e"),
                ObjectId("5cb572fb78163fa3c972701f"),
                ObjectId("5cb572fb78163fa3c9727020")
]

comments = [
    "That's a very good answer",
    "Are you kidding?",
    "Dude, what are you smoking?",
    "Couldn't have said it better myself",
    "Jane, you ignorant ...",
    "One of the best answers I have ever read",
    "Preach it brother!",
    "Man, I gotta tell you, that's ignorant"
]

repliesToComments = [
    "I agree",
    "I disagree",
    "I neither agree nor disagree",
    "I concur",
    "I object",
    "What a maroon!",
    "Looooooooser",
    "Brilliant, absolutely brilliant!"

]


for(i = 0; i < 10; i++){
    curId = getNextId()
    db.comments.insert
    (
        {
            "_id" : curId,
            "answerId" : answerIds[i],
            "body" : comments[i],
            "createdOn" :  new Date(),
            "createdBy": userIds[i]
        }
    )

    db.comments.insert(
        {
            "_id" : getNextId(),
            "answerId" : answerIds[i],
            "parentCommentId" : curId,
            "body" : repliesToComments[i],
            "createdOn" :  new Date(),
            "createdBy": userIds[i]

        }
    )
}


