function getNextId(){
    return ObjectId()
}


userIds = [
    "5cbe464972c9bfd18c02df81",
    "5cbeb9ae72c9bfdd4416f969",
    "5cbebb2572c9bfdd80680719",
    "5cbf42ef199080b5a1c486a2",
    "5cbf439e199080b5a1c486a3",
    "5cbf439e199080b5a1c486a4",
    "5cbf44da199080b5a1c486a5",
    "5cbf44da199080b5a1c486a6"
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


for(i = 0; i < 8; i++){
    curId = getNextId()
    db.comments.insert
    (
        {
            "_id" : curId,
            "answerId" : answerIds[i],
            "commentText" : comments[i],
            "createdOn" :  new ISODate(),
            "createdBy": userIds[i]
        }
    )

    db.comments.insert(
        {
            "_id" : getNextId(),
            "answerId" : answerIds[i],
            "parentCommentId" : curId,
            "commentText" : repliesToComments[i],
            "createdOn" :  new ISODate(),
            "createdBy": userIds[i]

        }
    )
}


