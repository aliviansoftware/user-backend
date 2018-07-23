# User Backend

Provides user management for the application

| Endpoint              |          Params            |        Description           |
|-----------------------|:--------------------------:|:----------------------------:|
| PUT /user/            | username and password json | Creates new user             |
| POST /user/login      | username and password json | Provides password validation |

## Example

### Create a user
```javascript
PUT http://127.0.0.1:1377/user/ 
// with the payload
{
	"username": "test",
	"password": "mypass"
}
```
A 200 status will be returned with the user being created

### Log a user in
```javascript
POST http://127.0.0.1:1377/user/ 
// with the payload
{
	"username": "test",
	"password": "mypass"
}
```
A 200 status will be if the credentials are correct

A 404 will be returned if the username doesnt exist

A 500 will be returned withthe response body of ```"error": "Incorrect password"``` when the password doesnt match

