# judgebool

The boolcmp tool is a static analysis tool for Go that detects variables with bool values that use comparison operators, and reports them as potentially simplifiable. 

# install

```
$ go install github.com/kyosu-1/boolcmp
```

# Usage

```
$ go vet -vettool=$(which boolcmp) path/to/your/package
```

# Example

Consider the following Go code:

```go
package main

func main() {
    var a bool = true
    if a == true {
        // do something
    }
}
```

The boolcmp tool will detect the comparison `a == true` and report it as a warning. The output might look something like this: 

```
main.go:5:10: bool value used in comparison
```

The warning indicates that the comparison `a == true` can be simplified to just a. This can make the code more concise and easier to read, since the comparison with `true` is unnecessary.