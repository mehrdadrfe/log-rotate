package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func ConcurrentFilesToDelete(configs []Config) {
	fileChan := make(chan string, 100) // buffered channel for file paths
	var wg sync.WaitGroup

	// Start 4 worker goroutines
	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for path := range fileChan {
				info, err := os.Stat(path)
				if err != nil {
					fmt.Printf("[Worker %d] Error stat %s: %v\n", workerID, path, err)
					continue
				}
				// Example: get retention days from parent config — here we use 7 days for demo
				cutoff := time.Now().AddDate(0, 0, -7)
				if info.ModTime().Before(cutoff) {
					err := os.Remove(path)
					if err != nil {
						fmt.Printf("[Worker %d] Failed to delete %s: %v\n", workerID, path, err)
					} else {
						fmt.Printf("[Worker %d] Deleted: %s\n", workerID, path)
					}
				}
			}
		}(i)
	}

	// Walk through all configs and send files to channel
	for _, config := range configs {
		fmt.Printf("Scanning directory: %s (retention: %d days)\n", config.LogDir, config.RetentionDays)

		entries, err := os.ReadDir(config.LogDir)
		if err != nil {
			fmt.Printf("Error reading dir %s: %v\n", config.LogDir, err)
			continue
		}

		cutoff := time.Now().AddDate(0, 0, -config.RetentionDays)

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			path := filepath.Join(config.LogDir, entry.Name())
			info, err := entry.Info()
			if err != nil {
				fmt.Printf("Error stat %s: %v\n", path, err)
				continue
			}
			if info.ModTime().Before(cutoff) {
				// send to workers
				fileChan <- path
			}
		}
	}

	close(fileChan) // no more files coming
	wg.Wait()       // wait for all workers to finish
}
