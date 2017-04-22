# spark
Golang client for Cisco Spark API

## Supported API calls

### Rooms
Method | Description | Example
--- | --- | --- 
ListRooms | Lists Rooms Available | 
GetRoom | Gets a Room with a Room ID |
GetRoomWithName | Gets the first Room where Name matches|

### Messages
Method | Description | Example
--- | --- | --- 
ListMessages | Lists all messages in a Room |
CreateMessage | Creates a message to a Room |
GetMessage | Gets a message by Id |

### Webhooks
Method | Description | Example
--- | --- | --- 
CreateWebhook | Creates a new webhook |
ListWebhooks | Lists existing webhooks |
DeleteWebhook | Deletes a webhook given a webhook ID | 
## Example

```go
package main

import (
  "github.com/vallard/spark"
)

const (
  token       = "your-spark-access-id"
  roomName = "your-spark-room-name"
)

func main() {
  s := spark.New(token)
  
  // Get the room ID of the room name
  room, err := s.GetRoomWithName(roomName)
  
  if err != nil {
  	panic(err)
  }
  
  // Create the message we want to send
  m := spark.Message{
  	RoomId: room.Id
  	Markdown: "# Big Message right here!"
  }
  
  // Post the message to the room
  _, err := s.CreateMessage(m)
  
  if err != nil {
    panic(err)
  }
}
```

Read the test cases for more ways to use the library.

## Inspiration

[bluele/slack](https://github.com/bleule/slack) 
