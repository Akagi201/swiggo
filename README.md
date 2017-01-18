# swiggo

[![Build Status](https://travis-ci.org/Akagi201/swiggo.svg)](https://travis-ci.org/Akagi201/swiggo)
[![Go Report Card](https://goreportcard.com/badge/github.com/Akagi201/swiggo)](https://goreportcard.com/report/github.com/Akagi201/swiggo)
[![GoDoc](https://godoc.org/github.com/Akagi201/swiggo?status.svg)](https://godoc.org/github.com/Akagi201/swiggo)

Call C/C++ lib with the help of swig

## Generate Go package file and C/C++ wrapper file
* `mv interface.swig/interface.swigxx interface.i`
* For C: `swig -go -cgo -intgosize 64 interface.i`
* For C++: `swig -go -cgo -c++ -intgosize 64 interface.i`

## Usage
* `go get github.com/Akagi201/swiggo`

## TODO
- [ ] Auto gen Go package files for godoc and IDE to browser Go definitions.
