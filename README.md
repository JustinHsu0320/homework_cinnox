## DEV Setup
```
# Clone
$ git clone https://github.com/JustinHsu0320/homework_cinnox

cd homework_cinnox

# Install Packages
go mod tidy

# Use Makefile to set up DEV env by Docker
$ make network
$ make mongodb
$ make collection NAME=users
$ make collection NAME=messages

# Init APP
$ make server

# Expose APP for Linebot
$ ngrok http 8080
```

## APIs
```
# 1. Requested by Linebot
[POST] http://0.0.0.0:8080/webhook

# 2. Respond in Line Chatroom
[POST] http://0.0.0.0:8080/messages/send
{
    "message": "Good day, mate 👍"
}

# 3. Get individual users from messages
[GET]  http://0.0.0.0:8080/messages/users
```
