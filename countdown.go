package main

import (
        "fmt"
        "time"
        "os"
)


func main() {
        timer := 60
        nano_count := time.Nanoseconds()
        for {
        
          fmt.Printf("%v: Timer: %v \n",timer,time.Nanoseconds() )
          timer--
          time.Sleep(1000000000 - (time.Nanoseconds() - nano_count))
          if timer < 1 {
                os.Exit(0)
          }
          nano_count = time.Nanoseconds()
        }
}

