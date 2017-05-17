package main

import "fmt"
import "math"
import "time"

const s string = "constant"

func main() {
    fmt.Println("hello world")
    fmt.Println("go" + "lang")
    //Integers and floats.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)
    //Booleans, with boolean operators as you’d expect.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

    //var declares 1 or more variables.
    var a string = "initial"
    fmt.Println(a)

    //You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)

    //Go will infer the type of initialized variables.
    var testTrue = true
    fmt.Println(testTrue)

    //Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
    var e int
    fmt.Println(e)

    //The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.
    f := "short"
    fmt.Println(f)


    const n = 500000000
    const d = 3e20 / n
    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")




    //There is no ternary if statement in Go
    // https://gobyexample.com/arrays

    //https://github.com/golang/example

    // https://github.com/golang/go/wiki/Projects
}
