## 1. Creating a executable from a go file.

We can use the go compiler to compile a executable from a go file and run it like any other binary.
We can compile a executable with go compiler like below:

```console
go build hello.go
```

## 2. Running go files without creating an executable.

If we want to keep making changes to the file, we would also have to compile it everytime into an executable.
That's not really convient, that's why there is the `go run` command which will compile and run our go file
at the same time without creating an executable.

```console
go run hello.go
```

## 3. Go packages.

A go program can contain many go files, organized into packages. Packages are like directories containing
a set of go files dedicated to do a specific task/work on your go program. If we were to create a calculator
, we could have a package called `calc` to contain all the logical files and a `io` package to hold all the
files that handle the program's input and output process.

Below is how a common go package file looks like:

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

Breakdown:

```go
package main
```

This line is called package declaration. It tells the Go compiler which package this file belongs to.
If the package declaration ins `package main`, then this program will be compiled into a executable.

```go
import "fmt"
```

With the import statement we can import and use code from other packages inside our program.

```go
func main () {
    fmt.Println("hello world")
}
```

1. We use the `func` keyword to declare a Go function followed by the name of the function, in this case it is `main`
2. After the name there follows a pair of parenthesis () and a set of curly braces {}
3. The function code is written inside the set of curly braces.

> This is a pretty standard way to declare functions in many high level languages except the `func` keyword.

## Function invocation

Generally to invoke a function, we call the declared function with a set of parenthesis () like this:

```go
Hello()
```

That is also same for Go also except for `main` functions. The `main` function is different and
reserved if it resides inside the `main` package. When we have a `main` function inside `main`
package, this is special in Go. When compiled an executable is produced and when run the executable uses
the `main` function as the starting point.

## Importing multiple packages

To import multiple packages, we can add multiple statements:

```go
import "math"
import "html"
```

Or use a single `import` statement with parenthesis with all the packages inside:

> :bulb: We don't need to use comma `,` after newlines

```go
import (
    "math"
    "html"
)
```

## Package alias

We can also add aliases to the imported packages by assining a name before the original package name:

```go
import (
    m "math"
    h "html"
)
```

Now we can just use the alias instead of the original package name to call any function inside it lie below. Once we assign an alias
to a package we have to use the alias only and not the original package name.

```go
import (
    f "fmt"
    m "math"
)

func main(){
    f.Println(m.Sqrt(100))
}
```

## Comments

Like most programming languages Go supports two types of comments.

1. Line comments
2. Block comments

Line comments starts with `//` and rest of the line is ignored by the compiler.

```go
// This entire line and the one below is ignored by the compiler
// fmt.Println("hello world")
fmt.Println("This gets printed")
```

Block comments starts with `/*` and ends with `*/`

```go
This line is ignored
This line is ignored too
fmt.Println("This won't print")
```

## Go Resources

Go has a built-in cli program called `go doc` to extract and view documentation from go files.
For provided information about any Package, use the command `go doc` followed by the package name like below:

```go
go doc fmt
```

If you want to know about a specific exported function from a package you can use get it like below:

```go
go doc fmt.Println
```

> By default it will only show the exported function documentation, exported functions are those that are declared with an Uppercase letter at the start of the function name. If you want to see the unexported functions too, you need to use the flag `-u` flag

```go
go doc -u
```

# Review 1

This is the high level review of the things learnt so far:

- The Go compiler.
- Go packages, including the `main` for executables.
- Go functions, including the `main` function.
- Go comment types. `Line comments` and `Block comments`.
- Resource extraction in CLI with `go doc`.
