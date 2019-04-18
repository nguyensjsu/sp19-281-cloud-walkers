var realProgrammersQ1     = ObjectId("5cb4048478163fa3c9726fd9")
var realProgrammersQ2     = ObjectId("5cb4048478163fa3c9726fda")
var zerosAndOnesQ1        = ObjectId("5cb4048478163fa3c9726fdb")
var zerosAndOnesQ2        = ObjectId("5cb4048478163fa3c9726fdc")
var whySoRelatableQ1      = ObjectId("5cb4048478163fa3c9726fdf")
var whySoRelatableQ2      = ObjectId("5cb4048478163fa3c9726fe0")
var travelTipsAndHacksQ1  = ObjectId("5cb4048478163fa3c9726fdd")
var travelTipsAndHacksQ2  = ObjectId("5cb4048478163fa3c9726fde")

var user1 = "1000000"
var user2 = "1231000"
var user3 = "4200304"
var user4 = "1466220"
var user5 = "9822142"

var realProgrammers     = ObjectId("5cb3c8ab78163fa3c9726fb2")
var zerosAndOnes        = ObjectId("5cb3c8ab78163fa3c9726fb3")
var whySoRelatable      = ObjectId("5cb3c8ab78163fa3c9726fb4")
var travelTipsAndHacks  = ObjectId("5cb3c8ab78163fa3c9726fb5")


db.questions.insert
(
    {
        "_id" : realProgrammersQ1,
        "spaceId" : realProgrammers,
        "body" : "What's the best way to find malloc leaks?",
        "createdOn" :  new Date(),
        "createdBy": user1
    }
)

db.questions.insert
(
    {
        "_id" : realProgrammersQ2,
        "spaceId" : realProgrammers,
        "body" : "Why would anyone want to allocate and free memory instead of using garbage collection?",
        "createdOn" :  new Date(),
        "createdBy": user2
    }
)


db.questions.insert
(
    {
        "_id" : zerosAndOnesQ1,
        "spaceId" : zerosAndOnes,
        "body" : "Were the compilers of the first programming languages written in machine code?",
        "createdOn" :  new Date(),
        "createdBy": user3
    }
)


db.questions.insert
(
    {
        "_id" : zerosAndOnesQ2,
        "spaceId" : zerosAndOnes,
        "body" : "Big Data is still evolving, After Hadoop industry is focusing on Spark, Now what next after Spark ?",
        "createdOn" :  new Date(),
        "createdBy": user4
    }
)


db.questions.insert
(
    {
        "_id" : travelTipsAndHacksQ1,
        "spaceId" : travelTipsAndHacks,
        "body" : "As someone who lives in the United Kingdom, what is the most ridiculous thing you have witnessed a tourist do?",
        "createdOn" :  new Date(),
        "createdBy": user5
    }
)

db.questions.insert
(
    {
        "_id" : travelTipsAndHacksQ2,
        "spaceId" : travelTipsAndHacks,
        "body" : "Is Iceland an underrated travel destination?",
        "createdOn" :  new Date(),
        "createdBy": user4
    }
)



db.questions.insert
(
    {
        "_id" : whySoRelatableQ1,
        "spaceId" : whySoRelatable,
        "body" : "What would your reaction be if Quora implemented a rule banning swear words?",
        "createdOn" :  new Date(),
        "createdBy": user1
    }
)

db.questions.insert
(
    {
        "_id" : whySoRelatableQ2,
        "spaceId" : whySoRelatable,
        "body" : "What is the single insight that most changed your life?",
        "createdOn" :  new Date(),
        "createdBy": user4
    }
)

