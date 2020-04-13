/** enum:
 versao do 1_a com um canal a mais e um merge a mais
**/

//@Author: Marcelo Heredia
package main

import (
	"math/rand"
	"time"
)


var end chan bool
var leng1 = 15
var leng2 = 15
var leng3 = 15

func main () {
		generator1 := make (chan int)
		generator2 := make (chan int)
		generator3 := make (chan int)
		chanMerge := make (chan int)
		chanMerge2 := make (chan int)

		go createValues(generator1, leng1, 1)
		go createValues(generator2, leng2, 2)
		go createValues(generator3, leng3, 3)
		go merge(generator1, generator2, chanMerge)
		go merge(chanMerge, generator3, chanMerge2)
		go consumer(chanMerge2, leng1+leng2)

		end = make(chan bool)
		<- end

}

//create random values and put on the parametrized channel
func createValues(gen chan int, len int, idGenerator int){
	for i:=0; i<len; i++{
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(999)+1
		
		println("Generator: ",idGenerator, "\n Number generated: ", number)

		gen <- number

		time.Sleep(time.Duration(number)*time.Millisecond)
	}
}

//takes the data from the channels and merge into one channel 
func merge(gen1 chan int, gen2 chan int, mrg chan int){
	for {
		select{
		case data := <-gen1:
			mrg <- data
		case data := <-gen2:
			mrg <- data
		}
	}
}

//consumption of the data on the channel
func consumer(mrg chan int, len int){
	for i:=0; i<len; i++{
		data := <-mrg
		println("Received data: ", data)
	}
	end <- true
}