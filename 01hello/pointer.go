package main

import "fmt"

func change(x *int) *int {
	*x += 100 // we need to be explicit with the dereferencing
	return x
}

type Book struct {
	title string
}

func (this *Book) setTitle(title string){
	this.title = title // here we do not need to be explicit about dereferencing go does that for us
}



func main6() {
	x := 3
	y := &x // this saves the memeory location of x as the value of y hence y points to the memory location of the the vale of x

	*y = 80 // *helps us go to the memeory adderess in order to change it (dereference)

	fmt.Println(x, *y)
	fmt.Printf("%T %T \n",x, y) // y is a *int type because is saves memory location

	/*passing points to the function*/
	w := change(&*y) // is now a pointer  to x directly
	x =9 
	fmt.Println(*w, x, *y)

	/* struct */
	newBook := Book{}

	newBook.setTitle("Things fall apart") // we dont need to do (&newbook) since it is of type *Book. again go handles this for us
	fmt.Println(newBook.title)
}