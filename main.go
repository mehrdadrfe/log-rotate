package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "cli", "Name of the directory (for logging)")
	dir := flag.String("dir", "/var/log/nginx", "Directory to clean")
	days := flag.Int("days", 7, "Retention period in days")

	flag.Parse()

	fmt.Printf("Running log-rotate: name=%s, dir=%s, retention=%d days\n", *name, *dir, *days)

	config := Config{
		NameDir:       *name,
		LogDir:        *dir,
		RetentionDays: *days,
	}

	ConcurrentFilesToDelete([]Config{config})
}
