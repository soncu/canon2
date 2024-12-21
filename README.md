# canon2
Go - Simple shredding a file

This is an implementation of shred feature for given file. This is happened by writing random data to the file 3 times.
This file may contain any type of data. After overwriting the file is deleted.

There are also unit tests for checking the functionality. Tests are included for 
`Valid file`,
`Non existing file`,
`Empty file`,
`Read only file`,
`Large file`.


To run the application simple make sure your environment contains go compiler and runner (check `go version`).
Just do `go run shred.go` for a quick run.


To apply the tests run `go test -v` and see the results.
