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

> Typing

Note that you can specify the type of a variable using the type, or the 'syntactic sugar' `:=`, as following

``` go
var untyped = "i'm untyped"

// default way
var typed int = 35

// You can let Go to put the type un autommatic manner using :=
// Note that automatic types doesn't work with constants
autoTyped := 56
```

## Using arrays

``` go
const ARRAY_SIZE = 50
var arrayExample [ARRAY_SIZE]string 
```

## Compiling

To compile and run a Go application you just need to type

``` bash
go run <FILE-NAME>
```

## Packages

To create and manage packages inside our package, we'll need the `go.mod` file created with the `go mod init <MODULE-NAME>` command.

In that file, we have defined the `MODULE-NAME`.

If we want to create a new package, we can create a new folder, and put inside the line

``` go
// package <NEW-PACKAGE-NAME>
package random-package-name
```

And then, for using it inside other modules, import it as follows

``` go
// import "<MODULE-NAME>/<NEW-PACKAGE-NAME>"
import "some-package/random-package-name"
```

### Exporting

To export something from a package, al we need to do is `make Capital` the first letter of the desired export (constant, variable, function).

Example

``` go
package random-package-name

// Unexported function
func greet() {
	print("holis")
}

// Exported function. Can be used inside other packages
func ExportMe() bool {
	amIExported := true

	return amIExported
}
```

### Usage

To use a package element, just import the package and use the function with the syntax `packageName.ElementName`, as follows

``` go
import "some-package/random-package-name"

random-package-name.ExportMe()
```

