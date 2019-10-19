package main

import "fmt"

func main() {
	defer println("defer 1")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("go routine recover :", err)
		}
	}()
	defer println("defer2")
	panic("panic 1")
	panic("panic2")

}

/*
defer2
defer 1
go routine recover : panic 1
 */