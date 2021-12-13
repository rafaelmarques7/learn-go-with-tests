# Learn Go with tests

This repo includes the source code that was created while following the [learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests/) tutorial.

Notes: 
* The code is self documenting. 
  * When something is not obvious, the code includes comments that are intended to provide clarification. I believe that this will help to better understand the code when re-reading it.
  * This may mean that code will be more verbose and may have some unecessary comments - or so it could seem - but these comments still add some value, at least for someone learning Go! 
* Lessons are organized into folders, because each folder can only have one package.
* In Go, when you call a function or a method the arguments are copied.


Useful go scripts:
* `go test`
* `go test -v` - verbose mode
* `go test -bench=.` - runs benchmark tests
* `go test -cover` - shows code coverage
* `go test -race` - checks to see if there are any race conditions
