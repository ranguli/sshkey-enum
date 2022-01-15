package main

import (
         "flag"
         "fmt"
         "os"
)

func usage() {
     fmt.Printf("usage : %s -inputValue=123\n", os.Args[0])
     flag.PrintDefaults()
     os.Exit(0)
}

var input = flag.String("inputValue", "", "String value to display")

func main() {
     flag.Usage = usage
     flag.Parse()

     if flag.NFlag() == 0 {
         usage()
         os.Exit(-1)
     }

    fmt.Println("String value to display is : ", *input)
}
