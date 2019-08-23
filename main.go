package main

import (
    "net/http"
    "log"
    handlers "./handlers"
)

func main() {
  router := handlers.Router()

  log.Print("The service is ready to listen and serve.")
  log.Fatal(http.ListenAndServe(":8000", router))
}
