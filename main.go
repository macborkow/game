package main

import (
  "net/http"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "context"
  "log"
  "fmt"
  "time"
)

func main() {
  serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
  clientOptions := options.Client().
      ApplyURI("mongodb+srv://user:password1234@go.6wehgo4.mongodb.net/?retryWrites=true&w=majority").
      SetServerAPIOptions(serverAPIOptions)
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  _, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Success!")
  fs := http.FileServer(http.Dir("static/"))
  http.Handle("/", fs)

  http.ListenAndServe(":8080", nil)
}
