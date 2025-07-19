package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "cli", "Label name for this run")
	dir := flag.String("dir", "", "Directory to scan (e.g., /var/log/nginx)")
	days := flag.Int("days", 7, "Retention period in days")
	dryRun := flag.Bool("dry-run", true, "Only print files to delete; set false to actually delete")

	flag.Parse()

	fmt.Printf("Running log-rotate: name=%s, dir=%s, retention=%d days, You can use dry-run to just prints what it would delete, dry-run=%v \n", *name, *dir, *days, *dryRun)

	config := Config{
		NameDir:       *name,
		LogDir:        *dir,
		RetentionDays: *days,
		DryRun:        *dryRun,
	}

	ConcurrentFilesToDelete([]Config{config})
}
