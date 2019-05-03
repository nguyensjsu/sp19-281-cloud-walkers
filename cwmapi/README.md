#CWMAPI
##Overview
CWMapi is the CWoura Message service.  It provides scalable and stateless access to the CWoura message data.  The current implentation is using MongoDB, but the service and API has been modeled such that it would be easy to move to a different NOSQL store such as Cassandra or Riak.

## Data Model
The CWuora app is built around topics, questions, answers, and comments.  Topics are tags appolied to questions.  They are free-form, and can be created and applied without any restrictions.  If a new question has a topic that is no in the topic table, it will be added.  A question has text, user id, zero or more answers, and zero or more topic tags.  Answers have text, user id, and zero or more comments.  Comments have text, user id, and zero or more nested comments (replies).  This is shown in the figure below.

todo: figure

##cwmapi API
The cwmapi API was built using OpenIO 3.0, which provides for both always current and upto date [online documentation](https://app.swaggerhub.com/apis-docs/jonathannah/cwmapi/1), and code stubs for both client and server.




