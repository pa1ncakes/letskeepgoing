package main

import  (
  "fmt"
  "example.com/greetings"
)

func main() {
  // Gets a greeting message and print it.
  message := greetings.Hello("tarun")
  fmt.Println(message)
}


