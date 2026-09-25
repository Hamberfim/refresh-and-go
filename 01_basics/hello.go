package main

import (
	"fmt"

	"rsc.io/quote"

	"time"
)

func main() {
	fmt.Println("Hello, It's Go Time!")
	fmt.Println(quote.Go())

	fmt.Println("Welcome, the time is", time.Now().Format(time.UnixDate))
}
