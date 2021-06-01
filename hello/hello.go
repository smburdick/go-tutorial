package main

import (
  "fmt"
  "log"
  "time"
  "example.com/greetings"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)
  // Get a greeting message and print it.
  message, err := greetings.Hello("")
  if err != nil {
    log.Fatal("[" + time.Now().Format("2006-01-02 15:04:05") + "] ", err)
  }
  fmt.Println(message)
}
