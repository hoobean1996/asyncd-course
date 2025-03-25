# Asyncd
- Ent
  - EntTask Schema
- GraphQL
  go get entgo.io/contrib/entgql@master

Ent 实际上是一个ORM

## 目标 设计一个异步任务的平台
Task:
- id
- handler
- parameter
- priority

asyncd
\_ ent
    \_ schema
        \_ enttask.go


# RESTful
/api/v1/foo?query_strings..

{

}

HTTP Verb


# GraphQL

{
  tasks {
    handler
  }
}

{
  "tasks": [
    {
      "handler": "echo1"
    },
    {
      "handler": "echo2"
    }
  ]
}

{
  tasks {
    handler
    priority

    owner {
      name
      id
    }
  }
}

{
  "tasks": [
    {
      "handler": "echo1",
      "priority": "0",
      "owner: {
        "name": "xx",
        "id": "1234"
      }
    }
  ]
}

GraphQL 底层还是HTTP协议

GraphiQL 用来debug的一个交互式的页面 (postman)