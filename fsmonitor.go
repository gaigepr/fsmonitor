package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"code.google.com/p/go.exp/fsnotify"
	"github.com/hishboy/gocommons/lang"
)

func isMember(element string, array []string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return true
		}
		return false
	}
}

func collectPaths(path []string) []string {
	// paths to be returned
	collectedPaths := make([]string, 1, 1)

	//
	for _, thisPath := range paths {
		err := filepath.Walk(thisPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				collectedPaths = append(collectedPaths, path)
			}
			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	return collectedPaths
}

func SetupWatch(paths []string, excludes []string) (int, *fsnotify.Watcher) {
	var watchedCount int

	paths = collectPaths(paths)
	excludes = collectPaths(excludes)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error establishing watcher: ", err)
	}

	// establish watches
	for _, path := range paths {
		if !(isMember(path, excludes)) {
			err = watcher.Watch(path)
			if err != nil {
				fmt.Println("Error: ", err, "  establishing watch on: ", path)
			}
			watchedCount++
		}
	}
	return watcher, watchedCount
}
