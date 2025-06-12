# gmfi
Get file info fast and simple
 
<img src="photo.png" width="500px">

#### Feel free to contribute! 
 
## Install

#### Fastest way 
Run this command in terminal and it will install everything itself
```sh
curl https://raw.githubusercontent.com/jvqtil/gmfi/refs/heads/main/install.sh | sh
```
or if you prefer GoLang package manager use
```sh
go install github.com/jvqtil/gmfi@latest
```
#### Manual way
Go to [releases](https://github.com/jvqtil/gmfi/releases/) and download latest binary for your OS, then move it to `/usr/local/bin/` and enjoy with simple `gmfi` in terminal!

## Building
- Install [Go](https://go.dev/) and make sure it's working with `go version`
- Clone repo
- Run `go build` in repo directory, then move it to `/usr/local/bin/`

## Usage
`gmfi [filename]` to see file / dir info

## Options: 
	-h, --help       Show the help 
	-v, --version    Show version information
