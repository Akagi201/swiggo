# swiggo

Call C/C++ lib with the help of swig

## Generate Go package file and C/C++ wrapper file
* `mv interface.swig/interface.swigxx interface.i`
* For C: `swig -go -cgo -intgosize 64 interface.i`
* For C++: `swig -go -cgo -c++ -intgosize 64 interface.i`

## Usage
* `go get github.com/Akagi201/swiggo/simple`

## TODO
- [ ] Auto gen Go package files for godoc and IDE to browser Go definitions.
