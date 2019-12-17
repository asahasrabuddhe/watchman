package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"

	"log"
)

var Version string

func main() {
	app := cli.NewApp()

	app.Name = "watchman"
	app.Usage = "A simple program to watch file(s) or folder(s) and execute commands when something changes"
	app.Version = Version

	app.Authors = []*cli.Author{
		{
			Name:  "Ajitem Sahasrabuddhe",
			Email: "ajitem.s@outlook.com",
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "Path to the `FILE` to monitor for changes",
		},
		&cli.StringSliceFlag{
			Name:    "directory",
			Aliases: []string{"d"},
			Usage:   "Path to the `DIRECTORY` to monitor for changes",
		},
		&cli.StringFlag{
			Name:    "exec",
			Aliases: []string{"e"},
			Usage:   "The `COMMAND` to execute when a change is detected",
		},
	}

	app.Action = func(c *cli.Context) error {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return err
		}
		defer watcher.Close()

		done := make(chan bool)

		go func() {
			<-c.Context.Done()
			done <- true
		}()

		go func() {
			for {
				select {
				case event := <-watcher.Events:
					log.Printf("EVENT %#v", event)

					op, err := exec.Command(c.String("exec")).CombinedOutput()
					if err != nil {
						log.Println(err)
					}

					fmt.Println(string(op))
				case err := <-watcher.Errors:
					fmt.Println("ERROR", err)
				}
			}
		}()

		files := c.StringSlice("f")
		for _, file := range files {
			if err := watcher.Add(file); err != nil {
				return err
			}
		}

		dirs := c.StringSlice("d")
		for _, dir := range dirs {
			if err := watcher.Add(dir); err != nil {
				return err
			}
		}

		<-done

		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
