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
	fmt.Printf("start tracking: %s: %v\n", taskname, t.Format(time.Stamp))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	elapsed := time.Since(t)

	fmt.Printf("finish tracking: %s: %v\n", taskname, elapsed.Truncate(time.Second))

	f, err := os.OpenFile(tthist, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	histEntry := fmt.Sprintf("%s\t%d\t%d\n", taskname, t.Unix(), int64(elapsed.Seconds()))
	_, err = f.WriteString(histEntry)
	if err != nil {
		panic(err)
	}
}
