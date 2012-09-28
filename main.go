package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	exit := make(chan int, 0)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
		exit <- 1
	}()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Watch(pwd)
	if err != nil {
		log.Fatal(err)
	}

	<-exit

	/* ... do stuff ... */
	watcher.Close()
}
