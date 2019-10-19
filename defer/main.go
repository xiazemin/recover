package main

import "fmt"

func main() {
	defer println("defer in main")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover :", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover1 :", err)
		}
	}()
	defer println("defer in main2")
	go func() {
		/*defer func() {
			if err := recover(); err != nil {
				fmt.Println("go routine recover :", err)
			}
		}()
		*/
		defer println("defer in goroutine")
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("go routine recover :", err)
			}
		}()
		defer println("defer in goroutine1")
		//func() {
		panic("panic in goroutine")
		panic("panic in goroutine1")
		//}()
	}()
	panic("panic in main")
	panic("panic in main 1")

}
/*


defer in main2
defer in goroutine1
defer in main
defer in goroutine
recover1 : panic in main
go routine recover : panic in goroutine
 */