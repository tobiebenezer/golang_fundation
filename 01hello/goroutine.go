package main

import (
	"fmt"
	"time"
	"sync"
)

func addv (x int, y int, ch chan<- int, delay int) { //channel enable us to pass data between thread from our go thread to the main thread etc chan<- mean send only channel
	time.Sleep(time.Duration(delay) * time.Second)
	ch <- x + y
}

type Counter struct{
	count int
	lock sync.Mutex
}

// func (this *Counter) increaseCount(ch chan<- bool){
func (this *Counter) increaseCount(wg *sync.WaitGroup){
	this.lock.Lock() // lock the value
	defer this.lock.Unlock() // defer to the end of the function to unlock the thread resource
	this.count += 1
	fmt.Println(this.count)
	// ch <- true
	wg.Done()
}

func main(){
	// ch := make(chan int)

	// go addv(2300, 7888, ch)
	// x := <-ch // getting the value from the channel we use channel because we are unable to return a value from the function this is as a blocking operation
	// // deadlock is when two or more thread are waiting for a common resource
	// fmt.Println(x)

	// ch2 := make(chan int)
	// ch1 := make(chan int)

	// go addv(2300, 7888, ch1, 2)
	// go addv(200, 9, ch2, 2)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 		case x := <-ch1: // <-chan mean recieve only
	// 			fmt.Println(x)
	// 		case y := <-ch2:
	// 			fmt.Println(y)

	// 	}
	// }
	
	/* Buffer  Channel
		whenever a value is sent to a channel it wait for it to reciever to empty the sent
		the function of a buffer channel is to create space for receive and adding multiple send
	*/

	// ch := make(chan bool, 2) //buffer channel of two
	// ch <- true //send the first value
	// ch <- true // sending the secont value. a third send will lead to a deadlock channel keep waiting for reciever so a resever 
	// //must come next either a go route the receives it or a reciever
	// <-ch // another send can do work


	// fmt.Println("Done")

	/* Mutex thread lock 
		In a situation where threads access memory and update it  because the order of access is ot deterministic 
		we ge weird update behavior so a mutex thread lock help us fix this.
		Mutex block until a thread release it
	*/

	newCounter := Counter{}
	// ch := make(chan bool)

	//createing waithGroup
	//Instead of channels we could use waitGroups when we are not returning item from the channels
	wg := sync.WaitGroup{}
	wg.Add(100) // adding the number of theard or goroutine to be create so at the end of each we will signal done to decrease the wait number

	for i := 0; i < 100 ; i ++ {
		// go newCounter.increaseCount(ch)
		go newCounter.increaseCount(&wg)
		
	}

	//recieve the ch for each go routine dispatched
	// for i := 0; i < 100 ; i ++ {
	// 	<-ch	
	// }

	wg.Wait()
	
	
}