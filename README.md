# Project directory resolver
A program for resolving projects to directories.

## Usage
```
go-pdirs [options] <project>

Options:
    --help                   View help
    --scripts                Print bash scripts
    --list                   List all projects
```

### Scripts
The application comes with a bash function for navigating to project directories.
In order to call that function, you will need to include it in your `.bashrc` (or equivalent):
```bash
eval "$(go-pdirs --scripts)"
```

## Configuring directories
The application requires you to have a configuration file in `$HOME/.pdirs`.
The configuration file is currently just a newline-separated file of key-value pairs:
```
project_1=/home/user/path/to/project_1
project_2=/home/user/path/to/project_2
```

The configuration file also supports environment variables as such:
```
project_1=$HOME/path/to/project_1
project_2=$HOME/path/to/project_2
```

## Installing
Currently only AUR packages exist:
- [go-pdirs@0.1.2-1](https://aur.archlinux.org/packages/go-pdirs)
- [go-pdirs-git@0.1.2-1](https://aur.archlinux.org/packages/go-pdirs-git)

## Building
To build the application, make sure you have `go` and `make` installed.
```bash
make build
chmod +x bin/go-pdirs

# Optionally, move the binary to /usr/local/bin
sudo mv "$(pwd)/bin/go-pdirs" /usr/local/bin/go-pdirs
sudo mv "$(pwd)/scripts/pcd" /usr/local/bin/pcd

# Or create a symlink in case you want to make it easier to update
sudo ln -fs "$(pwd)/bin/go-pdirs" /usr/local/bin/go-pdirs
sudo ln -fs "$(pwd)/scripts/pcd" /usr/local/bin/pcd
```
