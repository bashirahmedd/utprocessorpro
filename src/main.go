package main

/* import "fmt"

func main() {
    fmt.Println("Download using Video ID: ")
    i := 101
    fmt.Println(i)
} */

import(
    "os" 
    "fmt"
    "log"
  )
  
  func main() {
    fmt.Println("Launched the application...")
    dir, err := os.Getwd()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(dir)
  }