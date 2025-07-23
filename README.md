# log-rotate

A simple CLI tool written in Go to clean up old log files based on a retention policy.

## 🔧 Features

- Deletes log files older than a specified number of days
- Accepts multiple target directories
- Prints deleted files
- Easy to customize and extend

## 📦 Requirements

- Go 1.18 or higher

## 🚀 Getting Started

### Clone the repo:
```bash
git clone https://github.com/mehrdadrfe/log-rotate.git
cd log-rotate
go run main.go --name=nginx --dir=/var/log/nginx --days=7 --dry-run=false


## 🪵 Logging

By default, logs print to the console.

To log to a file (recommended for cron):

```bash
go run main.go --dir=/var/log/nginx --days=7 --dry-run=false > run.log 2>&1