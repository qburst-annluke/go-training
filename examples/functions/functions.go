package main

import "fmt"

// function that takes two integers and returns
// their sum as an integer.
func plus(a int, b int) int {
    return a + b
}

// like-typed parameters
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
