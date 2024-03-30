
go build creates an executable file for later use in Go lang

go run compiled the go file and run it to machine

Package main is special.  It defines a standalone executable
program, not a library.  Within package main the
function main is

also special—it’s where execution of the program begins

“The import declarations must follow the package declaration.”

## Variable declaration

This version of program uses a short variable declaration to declare and initialize `s` and `sep`, but we could equally well have declared the variables separately. There are several ways to declare a string variable; these are all equivalent.

```go
s := ""
var s string
var s = ""
var s string = ""
```

- The first form a short variable declaration, is the most compact, but it may be used only within a function, not for package-level variables.
- The second variable relies on default initialization to zero value for string, which is `""`.

The built-in function `make` creates a new empty `map`

Printf has over a dozen such conversions, which Go programmers call verbs. This table '

%d decimal integer
%x, %o, %b integer in hexadecimal, octal, binary


%f, %g, %e floating-point number: 3.141593 3.141592653589793 3.141593e+00


%t boolean: true or false


%c
rune (Unicode code point)


%s
string


%q
quoted string "abc" or rune 'c'


%v
any value in a natural format


%T
type of any value


%%
literal percent sign (no operand)”

Excerpt From
The Go Programming Language (Addison-Wesley Professional Computing Series)
Brian W. Kernighan
This material may be protected by copyright.