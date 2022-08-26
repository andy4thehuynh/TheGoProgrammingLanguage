Package main defines a standalone executable program, not a librarys.

## Loops

// Basic for loop
for initialization; condition; post {
        zero or more statements
}

// traditional "while" Loops
for condition {
        // ...
}

// use "blank idenfiers" when index of a loop not needed. Compiler will yell otherwise
for <underscore>, args := range os.Args[1:] {
        // ....
}

## Variables (only use the following types according to Brian and Alan)
// short variable declaration - only used in functions (not package level variables)
s := ""

// Default initialization to zero value
var s string

## String concatenation
- using += operator on a string will allow old contents to be garbage collected when no longer in use

- more efficient to use Join function from Strings package

## Built in Function - Make

- `make` creates a new empty map. it has other uses too

## Maps (like hashes)
// trick to set a map key and increment
```
counts := make(map[string]int)
input := bufio.NewScanner(os.Stdin)
for input.Scan() {
        counts[input.Text()]++
}
```

- the order of a primitive `map` is assumed to be random

## Printf (verbs)
- fmt.Printf has conversions (like %d%s), which Go Programmers call verbs
- escape sequences are like \n or \t in string literals

## Functions
- functions and other package-level entities may be declared in any order
