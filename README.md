# Go tutorial

## Getting started

``` bash
# Initialize go module
go mod init <module-path>
```

in `go`, all things must be inside a package. So, in the first line we must put `package <PACKAGE-NAME>`

Also, the `main` function is the entrypoint for a Go program. You can only have `one` _main_ function, because there can only exist `one` entrypoint.

> Importing a Package

To import a Package use `import <PACKAGE-NAME>`

``` go
/** Example */
import "fms"
```


``` go
package main

import 

func main() {
	print("hello world")
}
```

## Compiling

To compile and run a Go application you just need to type

``` bash
go run <FILE-NAME>
```
