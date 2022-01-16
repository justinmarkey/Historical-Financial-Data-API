package main

import (
	"fmt"
	"log"
	//"threadhandler/getsys"
)
func Testfunc (name string) {
    fmt.Println ("Hello" , name) 
}
func checkErr(err error) {
    if err != nil {
        log.Fatal("Error tripped : ", err)
    }
}

func main() {
    
}
