# Chalk MVP

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

### Requirements 
- [ ] Students can search other students
- [x] Students can bookmark other students 
- [ ] Students can message each other
