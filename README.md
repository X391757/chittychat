
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
2024/10/25 15:04:37 client1 enter the chat system
2024/10/25 15:04:37 The message has been broadcasted
```

Other clients, such as `client2`, will see:

```
2024/10/25 15:04:37 [1] server: Participant client1 joined Chitty-Chat at Lamport time 1
```

#### Sending the message

Once client1 send the message, you will see below in the server log:
```
2024/10/25 15:05:39 client1 sent the message
2024/10/25 15:05:39 The message has been broadcasted
```

Other client log:
```
2024/10/25 15:05:39 [6] client1: hi
```
### Leaving the Chat

To leave the chat, simply type `/quit`. The following message will appear:
Other clients will be notified:
```
2024/10/25 15:13:09 [16] System: Participant client1 left Chitty-Chat at Lamport time 16
```


And the server will log the departure:

```
2024/10/25 15:13:03 client1 sent the message
2024/10/25 15:13:03 The message has been broadcasted
```

### Additional Information

- **Lamport Timestamps**: The system uses Lamport timestamps to maintain message order. Each message will display the current Lamport time for consistent event sequencing across all clients.
- **Real-Time Updates**: Clients receive real-time notifications for joining, leaving, and messaging events, facilitating smooth communication in a distributed environment.

This completes the setup and usage instructions for the ChittyChat system. Happy chatting!
