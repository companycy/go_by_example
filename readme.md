Why are map operations not defined to be atomic?

Why is len a function and not a method?

Why don't maps allow slices as keys?

What's the difference between new and make?

In short: new allocates memory, make initializes the slice, map, and channel types.

The method set of any other named type T consists of all methods with receiver type T.
The method set of the corresponding pointer type *T is the set of all methods with receiver *T or T (that is, it also contains the method set of T).


return with defer

取决于该变量是不是事先已经存在.
如果不是事先存在, 则会受影响


It's easy to address the situation, though. Use the blank identifier to let unused things persist while you're developing.
use goimports -w

go get -v golang.org/x/tools/cmd/goimports
goimports -w nil_ex.go 


import "unused"

// This declaration marks the import as used by referencing an
// item from the package.
var _ = unused.Item  // TODO: Delete before committing!

func main() {
    debugData := debug.Profile()
    _ = debugData // Used only during debugging.
    ....
}


slice vs array
An array type definition specifies a length and an element type
An array's size is fixed; its length is part of its type 

A slice literal is declared just like an array literal, except you leave out the element count:



