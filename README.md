# go-morsecode
`morse.go` is command line utility built with go which encodes and decodes morsecode.

## Installation
`morse.go` uses goolang and codegangsta/cli for generating help text and parsing commands.
golang installation instructions can be found here: 
https://golang.org/doc/install
After installing golang, go-morsecode can be installed by running the following commands
`git clone https://github.com/aryasaatvik/go-morsecode.git` or by downloading it as zip archive.
change directory
`cd go-morsecode`
go install
`go install morse.go`

##Usage
`morse help` shows the help text
`morse encode` encodes text to morsecode
`morse decode` decodes morsecode to text
