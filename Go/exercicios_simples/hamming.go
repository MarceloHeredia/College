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
var leng = 100

func main () {
		generator1 := make (chan int,2)
		generator2 := make (chan int,2)
		generator3 := make (chan int,2)
		chanMerge := make (chan int)

		go createValues(generator1, 1, 2)
		go createValues(generator2, 2, 3)
		go createValues(generator3, 3, 5)
		go merge(generator1, generator2, chanMerge)
		go consumer(chanMerge2, leng1+leng2)

		end = make(chan bool)
		<- end

}

//create random values and put on the parametrized channel
func createValues(gen chan int, idGenerator int, mult int){
	for i:=0; i<leng; i++{
		
		println("Generator: ",idGenerator, "\n Number generated: ", i*mult)

		gen <- i*mult

		time.Sleep(time.Duration(i*mult)*time.Millisecond)
	}
}

//takes the data from the channels and merge into one channel 
func merge(gen1, gen2, gen3, mrg  chan int){
		v1 = <- gen1
		v2 = <- gen2
		v3 = <- gen3
		for {
			min := v1
			if v2 < min{
				min = v2
			}
			if v3 < min{
				min = v3
			}
			mrg <- min

			if min == v1{v1 <- gen1}
			if min == v2{v2 <- gen2}
			if min == v3{v3 <- gen3}
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