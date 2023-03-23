
# Net Centric - Quiz 3
## Objective

 Get familiar with UDP Socket Programming using Golang

### Questions
To build a chat application using udp sockets, the application include one server and multiple clients.
Server: server receives various incoming clients’ requests. It stores the information of the clients i.e. IP address and 
name in a list. Whenever a client logs out, it deletes that particular client entry from its list and updates it 
accordingly. 
Client: It first registers itself by sending its username to the server. Client can send a private message to a particular 
user by using the command @<username> or send a message to all user using @all
## Note

In order to run this program your computer have to install Go Programming language

## Run Locally

Start the server first
```bash
  cd server
```
```bash
  go run cmd/main.go
```

Start the client
```bash
  cd client
```
```bash
  go run cmd/main.go
```
