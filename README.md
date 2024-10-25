
# ChittyChat - Distributed Chat System

## How to Run the Program

### Server

To start the server, open a terminal and run the following command:

```bash
go run assignment3/server/server.go
```

The server will start and listen for incoming client connections.

### Client

#### Start a Client

Open a new terminal and enter the following command, replacing `[username]` with your chosen username:

```bash
go run assignment3/client/client.go -username=[username]
```

For example, if you want to start a client with the username `client1`, use:

```bash
go run assignment3/client/client.go -username=client1
```

To demonstrate the system’s functionality with multiple clients, you can start additional clients with unique usernames, such as:

```bash
go run assignment3/client/client.go -username=client2
```

```bash
go run assignment3/client/client.go -username=client3
```

#### Joining chat

Once the client has joined the chat, you will see the following prompt in your server log:

```
2024/10/25 16:21:59 client1 enter the chat system
2024/10/25 16:21:59 The message has been broadcasted
```

Other clients, such as `client2`, will see:

```
2024/10/25 16:22:01 [3] server: Participant client2 joined Chitty-Chat at Lamport time 2
```

#### Sending the message

Once client1 send the message, you will see below in the server log:
```
2024/10/25 16:22:10 client1 sent the message
2024/10/25 16:22:10 The message has been broadcasted
```

Other client log:
```
2024/10/25 16:22:10 [7] client1: I am client1
```
### Leaving the Chat

To leave the chat, simply type `/quit`. The following message will appear:
Other clients will be notified:
```
2024/10/25 16:22:49 [14] System: Participant client1 left Chitty-Chat at Lamport time 13
```


And the server will log the departure:

```
2024/10/25 16:22:49 client1 leave the chat system
2024/10/25 16:22:49 The message has been broadcasted
```

Note that the lamport time in the message type is the server's lamport time, and the client will add 1 to the server's lamport time after receiving the message, so the number in [] is 1 more than the number in the message.

