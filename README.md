# spark
Golang client for Cisco Spark API

## Supported API calls

Method | Description | Example
--- | --- | --- 
ListRooms | Lists Rooms Available | 

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
  err := s.PostMessage(roomName, "Hello, world!")
  if err != nil {
    panic(err)
  }
}
```

## Inspiration

The Slack [bluele/slack](https://github.com/bleule/slack)
