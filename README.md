## DEV Setup
```
# Clone
$ git clone https://github.com/JustinHsu0320/homework_cinnox
$ cd homework_cinnox

# Install Packages
$ go mod tidy

# Use Makefile to set up DEV env by Docker
$ make network
$ make mongodb
$ make collection NAME=messages

# Init APP
$ make server

# Expose APP for Linebot
$ ngrok http 8080
```

## APIs
```
# 1. Requested by Linebot. Redirect by Ngrok
[POST] https://2b08-2001-b400-e251-e3ae-4c28-59f2-e4f0-30cb.ngrok-free.app/webhook

# 2. Respond in Line Chatroom
[POST] http://0.0.0.0:8080/messages/send
{
    "message": "Good day, mate 👍"
}

# 3. Get individual users from messages
[GET]  http://0.0.0.0:8080/messages/users
```
