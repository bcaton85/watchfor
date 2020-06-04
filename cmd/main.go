package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func main() {
	directory := os.Args[1]
	log.Printf("Listening on file path: \"%v\"", directory)
	if os.Args[2] != "--" {
		log.Fatal("Illegal format")
	}

	commandWithArgs := os.Args[3:]

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}

	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}

				log.Println("Change in directory, running command...")
				executeCommand(commandWithArgs)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println(err)
			}
		}
	}()

	err = watcher.Add(directory)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func executeCommand(commandWithArgs []string) {
	cmd := exec.Command(commandWithArgs[0], commandWithArgs[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(stdout))
}
