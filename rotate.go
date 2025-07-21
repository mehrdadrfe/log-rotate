package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func ConcurrentFilesToDelete(configs []Config) {
	fileChan := make(chan string, 100)
	var wg sync.WaitGroup

	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for path := range fileChan {
				if configs[0].DryRun {
					log.Printf("[Worker %d] Would delete: %s\n", workerID, path)
				} else {
					err := os.Remove(path)
					if err != nil {
						log.Printf("[Worker %d] Failed to delete %s: %v\n", workerID, path, err)
					} else {
						log.Printf("[Worker %d] Deleted: %s\n", workerID, path)
					}
				}
			}
		}(i)
	}

	// Producer: walk configs, find old files, send to channel
	for _, config := range configs {
		log.Printf("Scanning directory: %s (retention: %d days)\n", config.LogDir, config.RetentionDays)
		entries, err := os.ReadDir(config.LogDir)
		if err != nil {
			log.Printf("Error reading dir %s: %v\n", config.LogDir, err)
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
				log.Printf("Error stat %s: %v\n", path, err)
				continue
			}
			if info.ModTime().Before(cutoff) {
				fileChan <- path
			}
		}
	}

	close(fileChan)
	wg.Wait()
}
