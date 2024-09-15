package main

import (
  "log"

  "download-data-script/fetcher"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  fetcher.SaveRelevantInformation()
}

