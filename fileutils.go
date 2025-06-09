package main

import (
	"fmt"
	"os"
	"time"
)

func FilesToDelete(configs []Config) {
	for _, config := range configs {
		fmt.Printf("Checking directory: %s (retention: %d days)\n", config.LogDir, config.RetentionDays)

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
			path := config.LogDir + "/" + entry.Name()
			info, err := entry.Info()
			if err != nil {
				fmt.Printf("Error stat %s: %v\n", path, err)
				continue
			}
			if info.ModTime().Before(cutoff) {
				err := os.Remove(path)
				if err != nil {
					fmt.Printf("Failed to delete %s: %v\n", path, err)
				} else {
					fmt.Printf("Deleted: %s\n", path)
				}
			}
		}
	}
}
