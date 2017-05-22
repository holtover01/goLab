// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

package main

import s "strings"
import "fmt"

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.
var p = fmt.Println

func main() {

    // Here's a sample of the functions available in
    // `strings`. Since these are functions from the
    // package, not methods on the string object itself,
    // we need pass the string in question as the first
    // argument to the function. You can find more
    // functions in the [`strings`](http://golang.org/pkg/strings/)
    // package docs.
    p("Contains:  ", s.Contains("test", "es"))//returns true
    p("Count:     ", s.Count("test", "t"))//returns two
    p("HasPrefix: ", s.HasPrefix("test", "te"))//returns true
    p("HasSuffix: ", s.HasSuffix("test", "st"))//returns true
    p("Index:     ", s.Index("test", "e"))//returns 1
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))//returns a-b
    p("Repeat:    ", s.Repeat("a", 5))//returns aaaaa
    p("Replace:   ", s.Replace("foo", "o", "0", -1))//returns f00
    p("Replace:   ", s.Replace("foo", "o", "0", 1))returns f0o
    p("Split:     ", s.Split("a-b-c-d-e", "-"))//returns an array of a b c d e
    p("ToLower:   ", s.ToLower("TEST"))//test
    p("ToUpper:   ", s.ToUpper("test"))//TEST
    p()

    // Not part of `strings`, but worth mentioning here, are
    // the mechanisms for getting the length of a string in
    // bytes and getting a byte by index.
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}

// Note that `len` and indexing above work at the byte level.
// Go uses UTF-8 encoded strings, so this is often useful
// as-is. If you're working with potentially multi-byte
// characters you'll want to use encoding-aware operations.
// See [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// for more information.
