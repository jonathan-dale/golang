package main

import "fmt"


// Go requires explicit return types
func plus(a int, b int) int {

    return a + b
}


// you can omit the type name for like param's
// until the last one like this
func plusplus(a, b, c int) int {

    return a + b + c
}

func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res1 := plusplus(1, 2, 3)
    fmt.Println("1+2+3 = ", res1)

}
