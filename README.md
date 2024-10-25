
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

#### Sending Messages

Once the client has joined the chat, you will see the following prompt in your server log:

```
2024/10/25 14:29:31 client enter the chat system
2024/10/25 14:29:31 The message has been broadcasted
```

Other clients, such as `client2`, will see:

```
Client joined with Lamport time: 2
2024/10/25 14:30:08 client enter the chat system
2024/10/25 14:34:52 client sent the message
```

### Leaving the Chat

To leave the chat, simply type `/quit`. The following message will appear:
Other clients will be notified:
```
2024/10/25 14:37:50 [6] System: Participant client left Chitty-Chat at Lamport time 6
```


And the server will log the departure:

```
2024/10/25 14:37:50 client leave the chat system
2024/10/25 14:37:50 The message has been broadcasted
```

### Additional Information

- **Lamport Timestamps**: The system uses Lamport timestamps to maintain message order. Each message will display the current Lamport time for consistent event sequencing across all clients.
- **Real-Time Updates**: Clients receive real-time notifications for joining, leaving, and messaging events, facilitating smooth communication in a distributed environment.

This completes the setup and usage instructions for the ChittyChat system. Happy chatting!
