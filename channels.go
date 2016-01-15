package main

import "fmt"

func emit(c *chan, pokes ...string) {
  fmt.Println(pokes);
  for _, poke := range pokes {
    messages <- poke
  }
}

func main() {
  messages := make(chan string)

  go func() { messages <- "ping" }()

  msg := <-messages
  fmt.Println(msg)
}
