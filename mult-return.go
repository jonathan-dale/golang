package main

import "fmt"


// the (int, int) shows that this 
// func  returns 2 ints
func vals() (int, int) {

    return 3, 5
}

func main() {

    // here we use 2 different return values to 
    // make mult assignments of vars a and b
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // to only use a subset of the returned values
    // use a blank identifier _.
    _, c := vals()
    fmt.Println(c)

}
