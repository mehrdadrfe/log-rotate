package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	name := flag.String("name", "cli", "Label name for this run")
	dir := flag.String("dir", "", "Directory to scan (e.g., /var/log/nginx)")
	days := flag.Int("days", 7, "Retention period in days")
	dryRun := flag.Bool("dry-run", true, "Only print files to delete; set false to actually delete")
	logPath := flag.String("log", "log-rotate.log", "Path to log file")
	logFile, err := os.OpenFile("log-rotate.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	flag.Parse()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Running log-rotate: name=%s, dir=%s, retention=%d days, You can use dry-run to just prints what it would delete, dry-run=%v \n", *name, *dir, *days, *dryRun)

	config := Config{
		NameDir:       *name,
		LogDir:        *dir,
		RetentionDays: *days,
		DryRun:        *dryRun,
		logPath:       *logPath,
	}

	ConcurrentFilesToDelete([]Config{config})
}
