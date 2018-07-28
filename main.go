package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"time"
)

func homeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}

func main() {
	t := time.Now()

	tthist, ok := os.LookupEnv("TT_HIST_FILE")
	if !ok {
		tthist = path.Join(homeDir(), ".tthist")
	}

	if len(os.Args) != 2 {
		fmt.Print("usage: tt taskname\n\n")
		os.Exit(1)
	}

	taskname := os.Args[1]
	fmt.Printf("start tracking: %s\n", taskname)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	elapsed := time.Since(t).Round(time.Second)

	fmt.Printf("tracking finished: %s: %v\n", taskname, elapsed)

	f, err := os.OpenFile(tthist, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(fmt.Sprintf("\n%s\t%d", taskname, elapsed/time.Second))
	if err != nil {
		panic(err)
	}
}
