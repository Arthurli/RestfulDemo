# RestfulDemo
这是一个使用 Go 语言编写的 Restful Demo, 使用 PostgreSQL 数据库

##实现 API

###Users:
####GET /users
List all users
####Example:

~~~
$curl -XGET "http://localhost:80/users"
[
  {
    "id": "21341231231",
    "name": "Bob",
    "type": "user"
}, {
    "id": "31231242322",
    "name": "Samantha",
    "type": "user"
} ]
~~~

####POST /users
Create a user
allowed fields:
  name = string
####Example:

~~~
$curl -XPOST -d '{"name":"Alice"}' "http://localhost:80/users"
{
  "id": "11231244213",
  "name": "Alice",
  "type": "user"
}
~~~

###Relationships:
####GET /users/:user_id/relationships
List a users all relationships
####Example:

~~~
$curl -XGET "http://localhost:80/users/11231244213/relationships"
[
  {
    "id": "222333444",
    "state": "liked",
    "type": "relationship"
}, {
    "id": "333222444",
    "state": "matched",
    "type": "relationship"
}, {
    "id": "444333222",
    "state": "disliked",
    "type": "relationship"
} ]
~~~

####PUT /users/:user_id/relationships/:other_user_id
Create/update relationship state to another user.
allowed fields:
   state = "liked"|"disliked"
If two users have "liked" each other, then the state of the relationship is "matched"
####Example:

~~~
$curl -XPUT -d '{"state":"liked"}'
"http://localhost:80/users/11231244213/relationships/21341231231"
{
  "id": "21341231231",
  "state": "liked",
  "type": "relationship"
}
$curl -XPUT -d '{"state":"liked"}'
"http://localhost:80/users/21341231231/relationships/11231244213"
{
  "id": "11231244213",
  "state": "matched",
  "type": "relationship"
}
$curl -XPUT -d '{"state":"disliked"}'
"http://localhost:80/users/21341231231/relationships/11231244213"
{
  "id": "11231244213",
  "state": "disliked",
  "type": "relationship"
}
~~~
