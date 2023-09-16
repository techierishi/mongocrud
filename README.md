This is CRUD operations on MongoDB written in Golang. This is created to test the k8s operator present at (https://github.com/techierishi/k8soperator)

## Docker steps 

```bash
# Create image
docker build -t ghcr.io/techierishi/mongocrud:latest .


# Push the image to Docker Hub
docker login
docker push ghcr.io/techierishi/mongocrud:latest

```
## How to run?
Clone the repository:
```sh
git clone git@github.com:techierishi/mongocrud.git
```
Next, change the current directory to the repository:
```sh
cd mognocrud
```

```sh
docker run -p 7070:7070 --name mongocrud ghcr.io/techierishi/mongocrud:latest
```

## Endpoints:
```sh
GET    /users/:email
POST   /users
PUT    /users/:email
DELETE /users/:email
```

### Get User
Send a `GET` request to `/users/:email`:
```sh
curl -X GET 'http://127.0.0.1:7070/users/ric@gmail.com'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Ric",
    "email": "ric@gmail.com",
    "password": "oldpass"
  }
}
```
### Create User
Send a `POST` request to `/users`:
```sh
curl -X POST 'http://127.0.0.1:7070/users' -H "Content-Type: application/json" -d '{"name": "Ric", "email": "ric@gmail.com", "password": "oldpass"}'
```
Response:  
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Ric",
    "email": "ric@gmail.com",
    "password": "oldpass"
  }
}
```
### Update User
Send a `PUT` request to `/users/:email`:
```sh
curl -X PUT 'http://127.0.0.1:7070/users/ric@gmail.com' -H "Content-Type: application/json" -d '{"password": "newpass"}'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Ric",
    "email": "ric@gmail.com",
    "password": "newpass"
  }
}
```

### Delete User
Send a `DELETE` request to `/users/:email`:
```sh
curl -X DELETE 'http://127.0.0.1:7070/users/ric@gmail.com'
```
Response:
```sh
{}
```

### Errors
```sh
{
  "error": "user not found"
}
```
