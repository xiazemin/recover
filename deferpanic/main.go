package main

func main() {
	defer func() {
		panic("panic 0")
	}()
	defer func() {
		if err:=recover();err!=nil{
			panic(err)
		}
	}()
	defer func() {
	  if err:=recover();err!=nil{
		  panic(err)
	  }
	}()
	defer func() {
		if err:=recover();err!=nil{
			panic(err)
		}
	}()
	defer func() {
		panic("panic 1")
	}()

	defer func() {
		panic("panic2")
	}()
	panic("panic ")
}

/*
panic: panic
	panic: panic2
	panic: panic 1 [recovered]
	panic: panic 1 [recovered]
	panic: panic 1 [recovered]
	panic: panic 1
	panic: panic 0
*/
