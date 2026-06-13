package main

import (
	"fmt"
	"errors"

)

func divide(a int, b int) (int, error) {

	if b == 0 {
		return 0 , errors.New("0 division not allowed")
	}

	return a / b, nil
}

func defferfunct(){
	fmt.Println("Deffered")
	r := recover() // this should be used in deferred function help us catch the error with out and allow us handle the panic
	if r == nil {
		fmt.Println("No error")
	}else {
		fmt.Println(r)
	}
}

func main4() {
	/* An error that happens at run time is called a Panic
		Alway put a defer statement as the first statetment in out function we can have more than one
		Panic() throws and error and every things that comes after it will will not run.

	*/
	// defer defferfunct() //will run even if a panic happens usefull for clean up operations ad defer waites till the end of the code
	// panic("There was a panic error")
	// fmt.Println("ok panic")

	/*Alternatively */

	result, err := divide(3, 5)

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}