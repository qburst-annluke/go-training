package main

import "fmt"

func main() {

    // with a single condition
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // with initialization, condition and increment
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // without condition
    // using `break` to exit the loop
    for {
        fmt.Println("loop")
        break
    }

    // using `continue` to the proceed
    // to the next iteration of the loop.
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
