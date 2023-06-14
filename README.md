# assignment_2023

![Tests](https://github.com/vrsb99/assignment_2023/actions/workflows/test.yml/badge.svg)

This is the backend assignment of 2023 TikTok Tech Immersion.

## Usage

1. To launch the services: 
```
docker-compose up -d --build
```
2. To send a POST request message
```
curl -X POST -H "Content-Type: application/json" -d '
    {
    "chat":"senderName:receiverName",
    "text":"hello",
    "sender":"senderName"
    }' 
    http://localhost:8080/api/send
```
3. To receive a GET request message
```
curl -X GET -H "Content-Type: application/json" -d '
    {
    "chat":"senderName:receiverName",
    "cursor":0,
    "limit":10,
    "reverse":true
    }' 
    http://localhost:8080/api/pull
```

## Tests

While the Github Actions workflow pass the http-server and docker-compose services, rpc-server will fail. It is suspected that this is due to the non-json api test requests.

## Sending Requests with `api/send`
General format should include a ':' between senderName and receiverName
The following is an example:
```
{
"chat":"senderName:receiverName",
"text":"hello",
"sender":"senderName"
}
```

## Pulling Requests with `api/pull`
The following is an example:

1. Chat specifies the senderName and Receivername with a ':'
2. Cursor specifies the starting point of the messages
3. Limit specifies the number of messages to be pulled
4. Reverse specifies the order of the messages
```
{
"chat":"senderName:receiverName",
"cursor":0,
"limit":10,
"reverse":true
}
```
Expected output
```
{
    "messages": [
        {
            "chat":"senderName:receiverName",
            "text":"Most recent message since reverse is true",
            "sender":"senderName"
        },
        {
            "chat":"senderName:receiverName",
            "text":"Second most recent message",
            "sender":"senderName"
        }
    ]
}
```
