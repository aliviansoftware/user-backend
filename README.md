# user-backend

| Endpoint              |          Params            |        Description           |
|-----------------------|:--------------------------:|:----------------------------:|
| PUT /user/            | username and password json | Creates new user             |
| GET /user/{username}  | none                       | Gets user by username        |
| POST /user/login      | username and password json | Provides password validation |