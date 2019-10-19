package main

import "fmt"

func main() {
	defer println("defer 1")
	defer func() {
		//if err := recover(); err != nil {
			//fmt.Println("go routine recover :", err)
		fmt.Println("defer2.1")
		println("defer2")
		fmt.Println("defer2.2")
		//}
	}()
	defer println("defer3")
	//panic("panic 1")
	//panic("panic2")

}

/*
defer2
defer 1
go routine recover : panic 1
 */