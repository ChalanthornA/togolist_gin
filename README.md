# togolist_gin
Change togolist from fiber to gin framework.

# Run
```
$ go run main.go
```

API:
- POST "/auth/signup":
  - require "username", "name", "password"
- POST "/auth/signin":
  - require "username", "password"
- GET "/auth/info" get user data
  - require "accesstoken" 
- POST "/todo/create" create new todo
  - require "accesstoken", "todo", "description"
- GET "/todo/usertodo" get all user todo
  - require "accesstoken"
- DELETE "/todo/deletetodo/:id" delete todo by todo_id
  - require "accesstoken", "todo_id"
- PATCH "/todo/updatetodo" update todo or description by id
  - require "accesstoken", "todo_id", "todo", "description"
