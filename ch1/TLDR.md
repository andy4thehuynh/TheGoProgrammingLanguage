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
// using += operator on a string will allow old contents to be garbage collected when no longer in use

// more efficient to use Join function from Strings package
