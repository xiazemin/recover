package main

import "fmt"

func main() {
	defer func() { // defer1
		err:=recover()
		fmt.Println(err)
	}()
	defer func() { // defer1
		err:=recover()
		fmt.Println(err)
	}()
	panic1()
	defer panic("error2") // 根本执行不到

}

func panic1() {
	defer func() { // defer2
		panic("error1") // panic2
	}()
	panic("error") // panic1
}
