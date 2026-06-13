package main

// import (
// 	"fmt"
// 	// "math" //math
// 	// "strconv" //string convations and other operations
// )

// type Class struct {
// 	name string
// 	population int
// 	subject string
// 	arm string
// }

// type Person struct {
// 	name string
// 	age int
// 	class Class
// }

// /*adding methods to struct*/
// func (this Person) getname() (name string){
// 	name = this.name
// 	return
// }

// func Main() {

	// fmt.Println("Hello World!")
	// uint - unsigned integer 2 ^n - 1
	// int - signed integer 2 ^(n-1) - 1
	// uint8, int8, uint16, int16, uint32, int32, uint64, int64
	// float - float for decimal numbers. default is float64
	// byte - 8-bit for any ASCII CODE OR CHARACTER
	// rune - int32 for any unicode character
	// bool - boolean values
	// string - sequence of characters in double quotes only
	// nil - same as null

	// const x uint8 = 255
	// y := 3000
	// type cast unit32(0)

	// fmt.Println(x)
	// fmt.Println("y: ", y)

	//fmt - format
	// fmt.Println - print line
	// fmt.Printf - print formated %T is type %v - value useful for string formating
	//%b - binary representation
	// %e - scientific notation
	// %f - float
	// %s - string
	// %.2f - 2dp
	// %10.2f - format to 2dp but add 10 space in front
	// %% - give persent
	// \" - give quote
	// fmt.Sprintf() - return the string fromating no printing involved

	// fmt.Printf("The value of x is: %v\n", x)

	// fmt.Println(math.Pow(4,2))

	//casting from string to int

	// x := "100000"

	// y, err := strconv.Atoi(x) //convert from string to int

	// price := "1234"

	// dbprice, _ := strconv.Atoi(price)

	// fmt.Printf("db price is: %v %v \n",dbprice, price)

	// we also have strconv.ParseInt, ParseFloat, ParseBool
	// y, err := strconv.ParseInt(x, 2, 0)

	// fmt.Println(y,err)

	/*
		comparisom operators
		<
		>
		<=
		>=
		==
		!=

		Logical operators
		|| - Or
		&& - and
		! - not
	*/

	// x := uint(8)
	// y := 10
	// z := x < uint(y) // if i am comparing two variable I need to have the right type for both but if I have a x < 9, 9 is converted automatically to the right typ fro the comparison to work

	// fmt.Println(z)

	/* If statement*/

	// if (int(x) % 5) == 0 &&  (int(x) % 3) == 0 {
	// 	fmt.Println("FizzBuzz")
	// } else if int(x) % 3 == 0 {
	// 	fmt.Println("Buzz")
	// } else {
	// 	fmt.Println("Fizz")
	// }

	/* Switch Statement Conditions and Conditionals */
	// a := 1

	// switch a {
	// case 1:
	// 	fmt.Println("One")
	// 	fallthrough // allow to check multiple cases
	// case 2 , 4 , 8: // checking multiple contitions
	// 	fmt.Println("Two")

	// default:
	// 	fmt.Println("Done")

	// }

	/* and way */

	// switch {

	// case a < 0:
	// 	fmt.Println("Negative numbers")

	// case a > 0, a > 5 , a < 90:
	// 	fmt.Println("Positive")

	// default:

	// 	fmt.Println("Zero")

	// }

	/* Forloop */

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	/* While like loop */

	// a := 40
	// for a != 0 {
	// 	fmt.Println(a)
	// 	a--
	// }

	/* Looping strings */

	// str := " Hello world!"

	// for _ , char := range str { // range is use the get the range for a given chracter since strings are represented in UTF-8 format which change take more that 1 byte to represent
	// 	fmt.Printf("%c ", char)
	// }

	/* Array
	The are fixed size data structure that stores object of the same type
	the size can not increase or reduce notation: var arr [size]type
	*/

	// arr := [2]int{2, 2}                    //array with two element
	// nest_arry := [2][2]int{{1, 2}, {2, 0}} //nested array

	// fmt.Println(nest_arry)
	// fmt.Println(arr)

	// arr2 := [...][2]int{{12, 32}, {8, 3}}

	// fmt.Println((arr2))

	// arr[0] = 9
	// arr2[0][0] = 80
	// nest_arry[1] = [2]int{59, 28}

	// fmt.Println(arr)
	// fmt.Println(arr2)
	// fmt.Println(nest_arry)

	// for _, nested := range nest_arry {
	// 	for _, x := range nested {
	// 		fmt.Println(x)
	// 	}
	// }

	/* Slices 
	pointer -> points the first element of slice to the the index on arr from which the slice was taken
	length -> len(slice)
	capacity -> capcity of slice is len(arr) from the index of slice[0]
	the start pointer does not change and we are allow to add more element from that point upto it capcity
	*/


	// arr := [7]int{40, 40, 0, 0, 23,2,4}
	// slice := arr[1:3]

	// fmt.Println(slice, len(slice), cap(slice))

	// sl := []string{"Hello", "World"} // a slice non array : and array will be created from this and become the refrence for the slice pointer, len, and capcity
	// fmt.Println(sl, len(sl), cap(sl))

	// //Appending to slice
	// for i := 0; i < 10; i++ {
	// 	sl = append(sl, fmt.Sprintf("%v", i))
	// 	fmt.Println(sl, len(sl), cap(sl))
	// }

	//Slice with Make
	// with make(arraytype, len, capcity)

	// sl := make([]int, 10) // for dynamaicaly reation and emty slice

	// sl := []string{"hello", "World", "Welcome", "To", "Go"}
	// test(sl)
	// for i, value := range sl {

	// 	ecape := "\n"
	// 	if i < len(sl) - 1{
	// 		ecape = " "
	// 	}

	// 	fmt.Printf("%v%v",value,ecape)
	// }

	/* Maps 
	similar to dictionary in python
	*/

	// mp := map[string]int {
	// 	"hello" : 1,
	// 	"world" : 300,
	// }

	// mp["god"] = 90000
	// delete(mp, "hello")

	// value, ok := mp["world"]
	// if !ok {
	// 	fmt.Println("Not Found!")
	// }else{
	// 	fmt.Println("value ",value)
	// }

	// fmt.Println(mp)

	// mp := make(map[uint]uint)
	// n := uint(100)

	// for number := uint(1); number <= n; number++ {
	// 	for i := uint(1) ; i <= 6; i++ {
	// 		if number % i == 0 {
	// 			mp[i]++
	// 		}
	// 	} 
	// }

	// fmt.Println(mp)

	/*
	Function
	// */

	// increament := func(arg int) int {
	// 		return arg+1
	// 	}

	// qu := callfunc( increament, 30)

	// fmt.Println(qu)

	// // addition := sum(1,34, 34, 90)
	// addition := sum([]int{1, 34, 90}...)

	// fmt.Println(addition)

	/* 
	Struct
	any item in a package file that should be accessed publicly my start with capital letter in the name other wise it is private
	*/

// 	p1 := Person{name:"Awodumila", age:34, class: Class{name: "Jss3", population: 30, subject: "Mathematics"}}
// 	fmt.Println(p1.class.name)
	
// }

//function callable
// func callfunc(callable func(int) int,arg int) int{
// 	return callable(arg)
// }

// func square(num int) int {
// 	return num * num
// }

//Variadic function

// func sum(num ...int) int{
//     fmt.Printf("%T  \n", num)

// 	count := 0
// 	for _, x := range num {
// 		count += x
// 	}

// 	return count
// }

// // name return value we can have multipule named return values
// func sum1(num ...int) (count int){
//     fmt.Printf("%T  \n", num)

// 	for _, x := range num {
// 		count += x
// 	}

// 	return 
// }

// func test(arr []string) {
// 	arr[0] = "Change this"
// }