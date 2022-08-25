package main

import ( "fmt" )


func main() {

    // implicit datatype conversion is not done by go

    // all var's is not initialized are initilized by 0 value.
    //boolean
    //var b bool default value = false
    var b bool = true
    fmt.Printf("%v, %T", b, b)

    /* int
    int8, int16, int32, int64
    more than that big package math lib
    */

    /* uint
    uint8, uint16, uint32. byte==uint8
    */

    // supported operations .. bit shifting using >> or << , AND(&), OR(|) , XOR(^), NOR(&^), +  - , / , %

    // float32 && float64 to store decimal time number,  suppoted op: +  - , / , %

    // complex64 and complex128 -> you can store complex numbers such as 1 + 2i, supported op : +  - , / , %
    // real(n) imag(n) are to functions to get real and imaginary part of the complex numbers.
    // complex(x,y) to make where x is the real part and y is imaginary part

    // text types:
    // string: declared in ""
    // string: utf-8 chars -> can be treated as array of char same as c
    // string in go are immutable
    // you can concatinate strings

    // rune -> utf-32 char -> declared in ''


}
