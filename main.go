package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gen2brain/beeep"
)

const (
	FolderListen = "C:\\SPI_LOG"
	AlertTitle   = "SPI LOG"
	expr         = "\\w{2}\\d{11}\\s+\\d{4}\\s+\\d{4}-\\d{2}-\\d{2}\\s+\\d{2}:\\d{2}:\\d{2}\\s+\\d{2}:\\d{2}:\\d{2}\\s+\\d\\s+\\w\\d{2}\\w\\d\\w{2}\\d\\w{2}\\d{2}\\s+\\w+\\s+\\w+\\s+\\w+\\s+\\w"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		stop("ERROR: Create watcher", err)
	}
	defer watcher.Close()
	if err := watcher.Add(FolderListen); err != nil {
		message := fmt.Sprintf("ERROR: Listen folder %s", FolderListen)
		stop(message, err)
	}
	beeep.Notify(AlertTitle, fmt.Sprintf("New SPI log system listen to %s folder", FolderListen), "")
	done := make(chan bool)
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				if event.Op == 0x1 {
					filePath := event.Name
					fileInfo, err := os.Stat(filePath)
					if err != nil {
						beeep.Notify(AlertTitle, fmt.Sprintf("ERROR: Get file into: %s", err), "")
						continue
					}
					if fileInfo.IsDir() {
						beeep.Notify(AlertTitle, fmt.Sprintf("Was created a subfolder: %s", filePath), "")
						continue
					}
					go func() {
						// handle file
						beeep.Notify(AlertTitle, fmt.Sprintf("HANDLE: %s", filePath), "")
					}()
				}
			case err := <-watcher.Errors:
				beeep.Notify(AlertTitle, fmt.Sprintf("ERROR: Watch error: %s", err), "")
			}
		}
	}()
	<-done
}

func stop(message string, err error) {
	beeep.Notify(AlertTitle, message, "")
	panic(err)
}
