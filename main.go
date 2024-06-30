package main

import (
  "fmt"
  "github.com/joho/godotenv"
)

func main() {
  url := "http://api.weatherapi.com/v1"
  fmt.Println("holi")
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Unable to load .env file. Check if you have a .env")
  }
  api_key := os.Getenv("API_KEY")

}

