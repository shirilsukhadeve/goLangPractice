package main

import (
        "fmt"
        "bufio"
        "os"
       )

func printinteger() {
    var i int = 4
    fmt.Printf("%v, %T \n" , i,i)

    j := 10
    fmt.Printf("%v, %T \n" , j,j)
}

func getinput() {
    var i int
    fmt.Scan(&i)
    fmt.Printf("%v, %T \n" , i,i)

    in := bufio.NewReader(os.Stdin)

    line, err:= in.ReadString('\n')
    fmt.Printf("%v, %T \n" , line,line)
    fmt.Printf("%v, %T \n" , err,err)

}

func main() {
    fmt.Println("Hello World")
    printinteger()
    getinput()
}
