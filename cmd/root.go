package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "watchfor",
	Short: "Watchfor runs given commands in response to filesystem changes",
	Long:  `Watchfor will run a given command when notified of a read/write/create of a file in the specified path`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println(arg)
		}
		runWatcher(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runWatcher(args []string) {
	directory := args[0]
	log.Printf("Listening on file path: \"%v\"", directory)

	commandWithArgs := args[1:]

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
				runGivenCommand(commandWithArgs)

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

func runGivenCommand(commandWithArgs []string) {
	cmd := exec.Command(commandWithArgs[0], commandWithArgs[1:]...)
	stdout, _ := cmd.Output()
	log.Println(string(stdout))
}
