
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


function getNextId(idx){
    return answerIds[idx]
}

var user1 = "5cbe464972c9bfd18c02df81"
var user2 = "5cbeb9ae72c9bfdd4416f969"
var user3 = "5cbebb2572c9bfdd80680719"
var user4 = "5cbf42ef199080b5a1c486a2"
var user5 = "5cbf439e199080b5a1c486a3"

var realProgrammersQ1     = ObjectId("5cb4048478163fa3c9726fd9")
var realProgrammersQ2     = ObjectId("5cb4048478163fa3c9726fda")
var zerosAndOnesQ1        = ObjectId("5cb4048478163fa3c9726fdb")
var zerosAndOnesQ2        = ObjectId("5cb4048478163fa3c9726fdc")
var whySoRelatableQ1      = ObjectId("5cb4048478163fa3c9726fdf")
var whySoRelatableQ2      = ObjectId("5cb4048478163fa3c9726fe0")
var travelTipsAndHacksQ1  = ObjectId("5cb4048478163fa3c9726fdd")
var travelTipsAndHacksQ2  = ObjectId("5cb4048478163fa3c9726fde")


db.answers.insert
(
    {
        "_id" : getNextId(0),
        "questionId" : realProgrammersQ1,
        "answerText" : "The best option is to use a 3rd-party heap analysis tool such as SmartHeap?",
        "createdOn" :  new Date(),
        "createdBy": user5
    }
)

db.answers.insert
(
    {
        "_id" : getNextId(1),
        "questionId" : realProgrammersQ2,
        "answerText" : "Garbage collection comes at a cost of efficiency and performance.  For time-citical applications, " +
            "carefully planed memory management is needed. ",
        "createdOn" :  new Date(),
        "createdBy": user3
    }
)


db.answers.insert
(
    {
        "_id" : getNextId(2),
        "questionId" : zerosAndOnesQ1,
        "answerText" : "Well, no, actually! They were written in Assembly Language.\n" +
            "\n" +
            "Now, for the non-computer folks out there, the basic definition of an Assembly Language is that it " +
            "translates the instructions on a one to one basis into a machine code.\n" +
            "\n" +
            "The advantages of an assembly language include, but are not limited to allowing the Assembler Program to:\n" +
            "\n" +
            "    Use “tags” to represent an address\n" +
            "    Use combinations of instructions when a “pseudo-instruction” is referenced\n" +
            "    Use #2 above in a more complex form as a “macro”\n" +
            "\n" +
            "Oh! And, by the bye, the first Assembler programs? They were written in machine code. (You had to start " +
            "somewhere.)\n" +
            "\n" +
            "PS. A Compiler normally operates in a 2-step process:\n" +
            "\n" +
            "    The Compiler spits out Assembly Language\n" +
            "    Which is then assembled into machine code by an assembler program.",
        "createdOn" :  new Date(),
        "createdBy": user2
    }
)


db.answers.insert
(
    {
        "_id" : getNextId(3),
        "questionId" : zerosAndOnesQ2,
        "answerText" : "well, the next big thing after Apache Spark is Apache Flink. However, to prove this point we can see " +
            "the comparison between Spark and Flink. As there are limitations of Apache Spark, industries have started " +
            "shifting to Apache Flink– 4G of Big Data, because Flink generally overcomes the limitations of Apache Spark",
        "createdOn" :  new Date(),
        "createdBy": user1
    }
)


db.answers.insert
(
    {
        "_id" : getNextId(4),
        "questionId" : travelTipsAndHacksQ1,
        "answerText" : "When I was a student at Salford University I met a group of American exchange students from Detroit.\n" +
            "\n" +
            "They wanted to see as much of the U.K. as they could whilst over here and one of the trips they booked was a " +
            "coach tour around North Wales. Excitement started to grow when I told them Wales is another country, separate " +
            "from England.\n" +
            "\n" +
            "A couple of days before they were due to leave I asked if they’d managed to get their entry visas through in " +
            "time. They all started to get very worried as it hadn’t occurred to them they’d need a visa.\n" +
            "\n" +
            "\“Don’t worry,\” I said, \“they hardly ever check them anyway and the coach probably won’t even stop at the " +
            "border. Just wave your passport at the window as you drive past, they’ll see you’re American and everything " +
            "will be fine.\”\n" +
            "\n" +
            "When they got back a few days later I had a massive roasting. Apparently the bus was full of Americans and " +
            "the girls had asked the driver to let them know when they were approaching the border. As they drove through " +
            "an entire bus full of Americans all waved their passports at the sheep in the neighbouring fields. The " +
            "driver, so I’m told, didn’t stop laughing for the rest of their excursion.",
        "createdOn" :  new Date(),
        "createdBy": user2
    }
)

db.answers.insert
(
    {
        "_id" : getNextId(5),
        "questionId" : travelTipsAndHacksQ2,
        "answerText" : "I don’t think it’s underrated. I would say it is “rated” for all the wrong reasons. People come here " +
            "for Iceland’s nature under the claim that it is unique and unspoilt. The truth is that Iceland is a land " +
            "that was laid waste by man more than a 1000 years ago and has not recovered since. The lack of " +
            "distinguishable technological elements such as roads, wires and antennas give the impression that the land " +
            "is unspoiled, but where we now see rocky deserts there used to be a delicate sun-Arctic ecosystem of birch " +
            "forests. Nowhere else in Europe has man managed to strip the land bare to quite the same extent.",
        "createdOn" :  new Date(),
        "createdBy": user1
    }
)



db.answers.insert
(
    {
        "_id" : getNextId(6),
        "questionId" : whySoRelatableQ1,
        "answerText" : "Gosh darn diddly dang, are you serious? (Sorry to disappoint, but this will not be followed by a " +
            "present tense verb beginning with “F,” nor its gerund form, nor its four-letter noun taking wing, nor….)\n" +
            "\n" +
            "I am more of a private purveyor of cussing (alone and “closed company,”) and like a sailor, I can. But as " +
            "I tell my own kids, no one is offended when you don’t cuss. I substitute teach and should one of my " +
            "sweeties find my written word online, I hope I am sparing enough in the use of cursing to not be considered " +
            "a hypocrite for my “watch your language” admonitions in compliance with school policy.\n" +
            "\n" +
            "To be honest, as a writer, I have actually struggled with how far I want to “get real” in some of the " +
            "dialogue I might want to do or situations characters find themselves in. There is so much ugliness and " +
            "brutality in the world that I wouldn't mind writing something more on the beautiful side. When the time " +
            "comes, I guess I will know when to let my inner bawdy mouth loose in print.\n" +
            "\n" +
            "All that being freakin’ said, I think readers on Quora are quite astute in drawing their own conclusions " +
            "about whether a writer has used choice words well or dished them up to a fornicative level of excess and " +
            "meaninglessness.\n" +
            "\n" +
            "In short, censorship sucks.\n" +
            "\n" +
            "(I am old enough to remember when”sucks” used this way was considered cursing. Today, meh.)?",
        "createdOn" :  new Date(),
        "createdBy": user5
    }
)

db.answers.insert
(
    {
        "_id" : getNextId(7),
        "questionId" : whySoRelatableQ2,
        "answerText" : "In 7th grade (or was it 8th?) I first really recognized the beauty of the ordinary. It was the play " +
            "\“Our Town\” by Thornton Wilder; I had been assigned to read it for school. In that play, Emily dies, but i" +
            "s given the opportunity to return home, for one day. The other dead strongly advise her not to go. When she " +
            "insists, they tell her, they plead with her, to just pick a random, ordinary day. But she doesn't take their " +
            "advice. She picks a special day, one of the happiest days in her life, her 12th birthday.\n" +
            "\n" +
            "It was a mistake. She finds it was too much, too intense. Life was so beautiful, every moment! Why didn't " +
            "she savor it as it happened! She finds joy in the voice of her mother, of her friends, in the beauty of " +
            "her home. Filled with remorse at how little she appreciated the beauty of the ordinary, she returns to the dead.\n" +
            "\n" +
            "It was in reading that play, part of a school assignment, that I became aware of how blessed I was; how " +
            "wonderful life is, not in the great events or special moments, but every moment, whether it is a moment " +
            "of pleasure or pain.\n" +
            "\n" +
            "I went to a movie and was entranced by the beauty of the images and their color. But when I walked outside, " +
            "I realized that the real world had as much beauty and color, just in the middle of the Bronx, on an ordinary street.\n" +
            "\n" +
            "When my children were born, I created a little ceremony that we still celebrate. Every Good Friday (I’ll " +
            "explain why I chose that day in a moment) I notice little things, maybe a table top, maybe a bird outside " +
            "the window, maybe a pencil. I’ll look at it and enjoy its beauty. I brought my (then) young children into " +
            "this ceremony, on Good Friday pointing out how pretty the water was in their cups; how it shimmered and how " +
            "looking through it made funny images and how its surface fluttered. I told them that this was “Good Friday’s " +
            "Magic”. They too found the day special. Then a year later, when they were enjoying the Good Friday Magic, my " +
            "older daughter said to me, “Daddy … things are more beautiful on Good Friday … but isn’t that only because that’s " +
            "the day we notice them?” I told her she was growing wise.\n" +
            "\n" +
            "My younger daughter now lives away, but I know that on Good Friday I’ll receive a text message, something " +
            "about how pretty some ordinary object is on that day.\n" +
            "\n" +
            "Why Good Friday? In Wagner’s opera “Parsifal”, the protagonist Parsifal is returning home from the Holy " +
            "Land when he notices how beautiful the world is. He asks his squire why? Gurnemanz replies, “That is Good " +
            "Friday’s magic, my lord.” The music from this part of the opera is often played as a separate piece, called " +
            "\“The Good Friday Spell\” or “The Good Friday Magic.” I listen to the opera every Good Friday.\n" +
            "\n" +
            "When I pay attention to the ordinary, I feel a deep love of God. I feel blessed. The most wonderful part: " +
            "I can invoke this feeling at any time, any day, now, by simply noticing the beauty of the ordinary.",
        "createdOn" :  new Date(),
        "createdBy": user3
    }
)

