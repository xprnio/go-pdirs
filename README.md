# Project directory resolver

## Usage
```
go-pdirs [options] <project>

Options:
    --help                   View help
    --list                   List all projects
```

## Building
To build the application, make sure you have `go` and `make` installed.
```
make build
chmod +x bin/go-pdirs

# Optionally, move the binary to /usr/local/bin
sudo mv "$(pwd)/bin/go-pdirs" /usr/local/bin/go-pdirs

# Or create a symlink in case you want to make it easier to update
sudo ln -fs "$(pwd)/bin/go-pdirs" /usr/local/bin/go-pdirs
```
