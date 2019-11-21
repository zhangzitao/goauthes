# goauthes

goauthes is a simple oauth2 server written in golang. Different from normal authorize server, goauthes is separated, so verify server and authorize server will be no longer just one server. There is remote options and local options you could choose. Configuring remote URL and other options in ```.env``` file. You can even set max amount of sessions one user could get at the same time.

Most parts of this project is implemented by interface, it's so convienent to work with different implementations. Thanks to interface, this project's struct is loose coupling. 

## implementations examples
- server
  - standard server
  - gin
  - ...
- authorize
  - password mode
  - authorization code mode
  - ...
- storage
  - memory
  - redis
  - ...

## How to run
1. clone ```examples``` folder
2. run ```go run remote_password_verify_server``` , this will open localhost:12321/user api. The api handle post to make password verify.
3. run ```go run authorize_server```, this will open localhost:9096/token api
4. make request to get new token
```
http://localhost:9096/token?grant_type=password&username=123&password=321&scope=read
```
  this will return new token: 
```json
{"AccessToken":"ce285e89-2f05-4aec-9874-e93c78dc85f9","TokenType":"bearer","ExpiresIn":3600,"RefreshToken":"56722f33-84f4-4fbe-a259-e50b7d9b29b3"}
```
5. make request to refresh the token
```
http://localhost:9096/token?grant_type=refresh_token&refresh_token=56722f33-84f4-4fbe-a259-e50b7d9b29b3
```

## Contribute
You are welcome to pullrequest or point out issues