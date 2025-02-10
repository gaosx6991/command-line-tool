# Go Command Line Tools

A collection of common Unix/Linux command line tools implemented in Go, including `ls`, `cat`, `cp`, and more.

## Available Commands

### ls
Lists directory contents

```go
ls # List contents of current directory
ls /path # List contents of specified directory
```

### cat
Concatenate and print files

```go
cat filename   # Display contents of a file
```

### cp
Copy files and directories
```go
cp [OPTION]... SOURCE DEST
```
Supports various options including:
- `-h, --help`: Display help information
- More features coming soon...

### cd (Deprecated)
Change directory command implementation (Note: Due to Go's process isolation, this command cannot actually change the parent shell's working directory)

## Requirements

- Go 1.19 or higher

## Installation

1. Clone the repository
```bash
git clone [repository-url]
```

2. Build the commands
```bash
go build ./ls
go build ./cat
go build ./cp
```

## Development

This project is under active development. Feel free to contribute by:
- Adding new commands
- Improving existing implementations
- Adding more command options
- Writing tests
- Improving documentation
