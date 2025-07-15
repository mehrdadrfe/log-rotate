package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "cli", "Name of the directory (for logging)")
	dir := flag.String("dir", "/var/log/nginx", "Directory to clean")
	days := flag.Int("days", 7, "Retention period in days")
	dryRun := flag.Bool("dry-run", true, "If true, only print files; don’t delete")

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
