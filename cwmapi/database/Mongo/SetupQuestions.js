var realProgrammersQ1     = ObjectId("5cb4048478163fa3c9726fd9")
var realProgrammersQ2     = ObjectId("5cb4048478163fa3c9726fda")
var zerosAndOnesQ1        = ObjectId("5cb4048478163fa3c9726fdb")
var zerosAndOnesQ2        = ObjectId("5cb4048478163fa3c9726fdc")
var whySoRelatableQ1      = ObjectId("5cb4048478163fa3c9726fdf")
var whySoRelatableQ2      = ObjectId("5cb4048478163fa3c9726fe0")
var travelTipsAndHacksQ1  = ObjectId("5cb4048478163fa3c9726fdd")
var travelTipsAndHacksQ2  = ObjectId("5cb4048478163fa3c9726fde")

var user1 = "5cbe464972c9bfd18c02df81"
var user2 = "5cbeb9ae72c9bfdd4416f969"
var user3 = "5cbebb2572c9bfdd80680719"
var user4 = "5cbf42ef199080b5a1c486a2"
var user5 = "5cbf439e199080b5a1c486a3"


var realProgrammers     = ObjectId("5cb3c8ab78163fa3c9726fb2")
var zerosAndOnes        = ObjectId("5cb3c8ab78163fa3c9726fb3")
var whySoRelatable      = ObjectId("5cb3c8ab78163fa3c9726fb4")
var travelTipsAndHacks  = ObjectId("5cb3c8ab78163fa3c9726fb5")


db.questions.insert
(
    {
        "_id" : realProgrammersQ1,
        "questionText" : "What's the best way to find malloc leaks?",
        "createdOn" :  new Date(),
        "createdBy": user1,
        "topics" : [
            {
                "label": "Computer Programming",
            },
            {
                "label": "Computer Science",
            },
            {
                "label": "Software Engineering"
            }
        ]
    }
)

db.questions.insert
(
    {
        "_id" : realProgrammersQ2,
        "questionText" : "Why would anyone want to allocate and free memory instead of using garbage collection?",
        "createdOn" :  new Date(),
        "createdBy": user2,
        "topics" : [
            {
                "label": "Computer Programming",
            },
            {
                "label": "Computer Science",
            },
            {
                "label": "Software Engineering"
            }
        ]
    }
)


db.questions.insert
(
    {
        "_id" : zerosAndOnesQ1,
        "questionText" : "Were the compilers of the first programming languages written in machine code?",
        "createdOn" :  new Date(),
        "createdBy": user3,
        "topics" : [
            {
                "label": "Computer Programming",
            },
            {
                "label": "Computer Science",
            },
            {
                "label": "Software Engineering"
            }
        ]
    }
)


db.questions.insert
(
    {
        "_id" : zerosAndOnesQ2,
        "questionText" : "Big Data is still evolving, After Hadoop industry is focusing on Spark, Now what next after Spark ?",
        "createdOn" :  new Date(),
        "createdBy": user4,
        "topics" : [
            {
                "label": "Computer Programming",
            },
            {
                "label": "Computer Science",
            },
            {
                "label": "Software Engineering"
            }
        ]
    }
)


db.questions.insert
(
    {
        "_id" : travelTipsAndHacksQ1,
        "questionText" : "As someone who lives in the United Kingdom, what is the most ridiculous thing you have witnessed a tourist do?",
        "createdOn" :  new Date(),
        "createdBy": user5,
        "topics" : [
            {
                "label": "International Travel",
            },
            {
                "label": "Tourism",
            },
            {
                "label": "Visiting and Travel"
            }
        ]
    }
)

db.questions.insert
(
    {
        "_id" : travelTipsAndHacksQ2,
        "questionText" : "Is Iceland an underrated travel destination?",
        "createdOn" :  new Date(),
        "createdBy": user4,
        "topics" : [
            {
                "label": "International Travel",
            },
            {
                "label": "Tourism",
            },
            {
                "label": "Visiting and Travel"
            }
        ]
    }
)



db.questions.insert
(
    {
        "_id" : whySoRelatableQ1,
        "questionText" : "What would your reaction be if Quora implemented a rule banning swear words?",
        "createdOn" :  new Date(),
        "createdBy": user1,
        "topics" : [
            {
                "label": "Education",
            },
            {
                "label": "Life and Living",
            },
            {
                "label": "Students"
            }
        ]
    }
)

db.questions.insert
(
    {
        "_id" : whySoRelatableQ2,
        "questionText" : "What is the single insight that most changed your life?",
        "createdOn" :  new Date(),
        "createdBy": user4,
        "topics" : [
            {
                "label": "Education",
            },
            {
                "label": "Life and Living",
            },
            {
                "label": "Students"
            }
        ]
    }
)

