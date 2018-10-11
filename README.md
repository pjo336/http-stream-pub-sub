# http-stream-pub-sub

Random prototype to demonstrate publish/subscribe over http streaming.

#### Running:
1. Run the server
    - `go build server.go` (skip if using premade binary)
    - `./server`
    
2. Run the client (You can run multiple clients)
    - `node client.js`
    
3. Use NetCat or curl to tell the server an event occurred, and subscribers should be notified.
    - `curl localhost:8000/message`
    
    
 ##### TODO
 Currently, step 3 always sends a message to every registered client.
 
 In the future the server would find a way to distinguish who the message is meant for and only publish to that subscriber
 