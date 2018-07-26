package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	t := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("usage: tt taskname")
		os.Exit(1)
	}

	taskname := os.Args[1]
	fmt.Printf("start tracking: %s\n", taskname)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Printf("tracking finished: %s: %v\n", taskname, time.Since(t).Round(time.Second))
}
