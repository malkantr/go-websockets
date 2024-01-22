# Introduction - Mastering WebSockets With Go - An in-depth tutorial by ProgrammingPercy

[youtube link](https://youtu.be/pKpKv9MKN-E?si=4Bq5UBJbu7LzOYz7)
[github link](https://github.com/percybolmer/websocketsgo)

WebSockets is one of the oldest ways to communicate in a bi-directional way and is widely used today. It is supported by most browsers and is relatively easy to use. WebSockets are really useful when building real time apis.

In this tutorial, we will cover what WebSockets are and how they work, how to use them in Go to communicate between servers and clients. We will also explore some regular pitfalls that I've seen in WebSocket APIs, and how to solve them.

During the tutorial, we will be building a chat application where you can enter different chat rooms. The WebSocket server will be built using Go, and the client connecting in vanilla JavaScript. The patterns we learn and apply could easily be adapted when connecting using a Websocket Client written in Go, Java, React, or any other language.

We will learn about Authentication, Heartbeats using Ping & Pong, Cross Origin and a number of other useful WebSocket related things.

If you prefer a written format of the video instead you can visit my blog
https://programmingpercy.tech/blog/mastering-websockets-with-go/

- 00:00 Introduction
- 03:04 What are WebSockets
- 05:46 Project setup
- 18:46 Connecting Websocket
- 28:14 Adding Clients
- 33:46 Reading & Writing Messages
- 53:00 Events
- 01:12:24 HeartBeating
- 01:22:27 Jumbo Frames
- 01:24:36 Cross Origin
- 01:27:37 Authentication
- 01:55:10 TLS
- 02:02:18 Finalizing


## Tools

```
-- install openssl

cd tools/
bash certgen.bash
```

## Notes
