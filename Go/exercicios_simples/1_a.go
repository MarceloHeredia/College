/** enum:
 Desenvolva um algoritmo que faz o merge de
duas sequências de dados (inteiros).
Ele recebe em dois canais de entrada, faz merge e
escreve em um canal de saída. O merge é apenas a
junção, sem critério, transforma duas correntes de
itens em uma.
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

func main () {
		generator1 := make (chan int,5)
		generator2 := make (chan int,5)
		chanMerge := make (chan int)

		go createValues(generator1, leng1, 1)
		go createValues(generator2, leng2, 2)
		go merge(generator1, generator2, chanMerge)
		go consumer(chanMerge, leng1+leng2)

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
	}
}

//takes the data from the channels and merge into one channel 
func merge(gen1 chan int, gen2 chan int, mrg chan int){
	for {
		select{
		case data := <-gen1:
			println("Generator 1 - data: ", data)
			mrg <- data
		case data := <-gen2:
			println("Generator 2 - data: ", data)
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