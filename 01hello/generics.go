package main

import "fmt"

type Number interface {
	int | float64 | uint
}

func add[T Number](x T, y T) T {
	return x + y
}

func getValues [K comparable, V any] (mp map[K]V) []V {
	values := make([]V,1)

	for _, value := range mp {
		values = append(values, value)
	}

	return  values
}

/*Generic types with method*/
type GenericSlice[T any] []T

/* method to use with the created type */
func (self GenericSlice[T]) print() {
	for _,value := range self {
		fmt.Printf("%v ", value)
	}
}

/*Generic Struc */
type GenericStruc[T any] struct {
	name T
}

func main5() {
	// value := add(3, 4)
	// fmt.Println(value)

	// fmt.Println((getValues(map[string]int{"one":1 , "three":3})))

	//defining a variable with a custome generic
	v := GenericSlice[int]{1, 9, 9, 0}
	v.print()

	//initilising generic strucs
	me := GenericStruc[string]{"Tobi"}

	fmt.Println(me.name)
}