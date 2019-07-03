package main

import "fmt"

func main() {

    // Declaring one varable
    var a = "initial"
    fmt.Println(a)

    // Declaring multiple variables
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Infering type of initialized variables
    var d = true
    fmt.Println(d)

    //  _zero-valued_ initialization
    var e int
    fmt.Println(e)

    // The `:=` syntax
    // shorthand for declaring and initializing a variable
    f := "apple"
    fmt.Println(f)
}
