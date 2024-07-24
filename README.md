# Chalk MVP

## Tech Stack
- [Golang](https://go.dev/), [Gin](https://gin-gonic.com/), [fx](https://uber-go.github.io/fx/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Docker](https://www.docker.com/)

## Set up and Running
- Copy the contents of [`.env.example`](./.env.example) into `.env` file and update the contents to suit your requirements.
- Next you can run the app using the [`docker-compose.yml`](./docker-compose.yml) file provided.
- If you want to run the app without docker, the entry file is [main.go](./cmd/app/main.go)

## API Endpoints

<details>
  <summary>
    <code>GET</code>
    <code>/healthcheck/healthcheck</code>
    <code>Checks if user is authenticated or not</code>
  </summary>

  ### Parameters
  >|name|type|data type|description|
  >|----|----|---------|-----------|
  >|Token|Header|string|Format: `Bearer <Token>` where `Token` is the auth token received when logged in|

  ### Responses
  >|http code|response|
  >|---------|--------|
  >|200|`{"message": "ok"}`|
</details>

<details>
  <summary>
    <code>POST</code>
    <code>/auth/register</code>
    <code>Register a new student</code>
  </summary>

  ### Parameters
  >|name|type|data type|description|
  >|----|----|---------|-----------|
  >|StudentData|Body|JSON|Registers a new student and returns the ID back|

  ### Responses
  >|http code|response|
  >|---------|--------|
  >|201|`{"id": id}`|
  >|422|{"error": "User with email {email} already exists."}|
  >|400|`No information`. Returned when server cannot process json|

  ### StudentData 
  ```json
  {
    "name": "string",               
    "email": "string",              
    "password": "string",           
    "description?": "string",        
    "university?": "string",         
    "degree?": "string",             
    "skills?": "string",             
    "year_of_graduation?": "string"
  }
```
</details>

<details>
  <summary>
    <code>POST</code>
    <code>/auth/signin</code>
    <code>Sign in with details</code>
  </summary>

  ### Parameters
  >|name|type|data type|description|
  >|----|----|---------|-----------|
  >|LoginData|Body|JSON|Sign into the server. Returns an authentication token to be used with other API endpoints|

  ### Responses
  >|http code|response|
  >|---------|--------|
  >|200|`{"token": token}`|
  >|401|{"error": "User with email {email} does not exist"} or {"error": "Invalid credentials"}|
  >|400|`No information`. Returned when server cannot process json|

  ### LoginData 
  ```json
  {
    "email": "string",              
    "password": "string"          
  }
```
</details>

<details>
  <summary>
    <code>POST</code>
    <code>/auth/signout</code>
    <code>Sign out</code>
  </summary>

  ### Parameters
  >|name|type|data type|description|
  >|----|----|---------|-----------|
  >|Token|Header|String|Format: `Bearer <Token>`|

  ### Responses
  >|http code|response|
  >|---------|--------|
  >|202|`{"message": ok}`|
</details>

<details>
  <summary>
    <code>POST</code>
    <code>/bookmark/add</code>
    <code>Bookmark another student</code>
  </summary>

  ### Parameters
  >|name|type|data type|description|
  >|----|----|---------|-----------|
  >|Token|Header|String|Format: `Bearer <Token>`|
  >|BookmarkData| Body|JSON|Student email you want to bookmark. This also creates a code for a room which can later be used to send messages in|


  ### Responses
  >|http code|response|
  >|---------|--------|
  >|202|`{"message": ok}`|
  >|400|`No information`. Returned when server cannot process json|

  ### BookmarkData 
  ```json 
  {
    "student_email": "string"
  }
```
</details>

<details>
  <summary>
    <code>POST</code>
    <code>/bookmark/view</code>
    <code>View another student's bookmarks</code>
  </summary>

  ### Parameters
  >|name|type|data type|description|
  >|----|----|---------|-----------|
  >|Token|Header|String|Format: `Bearer <Token>`|
  >|BookmarkData| Body|JSON|Student's bookmarks you want to see|


  ### Responses
  >|http code|response|
  >|---------|--------|
  >|202|`{"message": ok}`|
  >|400|`No information`. Returned when server cannot process json|

  ### BookmarkData 
  ```json 
  {
    "student_email": "string"
  }
```
</details>




### TODO 
- [ ] Students can search other students
- [x] Students can bookmark other students 
- [ ] Students can message each other
- [ ] Implement tests
