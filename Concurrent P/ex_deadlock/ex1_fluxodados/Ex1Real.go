// Pedro Castro e Marcelo Heredia
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var fin chan bool

//para estes vetores, a posicao 0 refere-se ao valor gerado
// a posicao 1 refere-se ao ultimo valor gerado e exibido!
var gerados = []int{0,0}
var negativos = []int{0,0}
var positivos = []int{0,0}
var pares = []int{0,0}
var impares = []int{0,0}
var primo = []int{0,0}

func geradorValores(send chan int) {
	rand.Seed(time.Now().UnixNano())
	var numero int
	for {
		numero = rand.Intn(999) - rand.Intn(999)
		gerados[0]++
		send <- numero
	}
	fin <- true // termina o processo
}

func contaNegativos(rec chan int) {
	for {
		dado := <-rec
		if dado > 0 {
			positivos[0]++
		} else {
			negativos[0]++
		}
	}
}

func contaParesImpares(rec chan int) {
	for {
		dado := <-rec
		if dado%2 == 0 {
			pares[0]++
		} else {
			impares[0]++
			if dado > 0 {
				go ehPrimo(rec)
			}
		}
	}
}

func ehPrimo(rec chan int) {
	for {
		dado := <-rec
		if dado%2 == 0 {
			return
		}

		for i := 3; i*i <= dado; i += 2 {
			if dado%i == 0 {
				return
			}
		}
		primo[0]++
	}
}

func validaPrint(){
	for {
		if gerados[0]> gerados[1]+10 ||
		   negativos[0] > negativos[1]+10 ||
		   positivos[0] > positivos[1]+10 ||
		   pares[0] > pares[1]+10 ||
		   impares[0] > impares[1]+10 ||
		   primo[0] > primo[1]+10 {
			
			gerados[1] = gerados[0]
			negativos[1] = negativos[0]
			positivos[1] = positivos[0]
			pares[1] = pares[0]
			impares[1] = impares[0]
			primo[1] = primo[0]

			fmt.Println("<",gerados[1],",",positivos[1],",",negativos[1],",",pares[1],",",impares[1],",",primo[1],">")
		   }
		time.Sleep(time.Duration(5) * time.Millisecond)
	}
}

func main() {
	gerador1 := make(chan int)

	go geradorValores(gerador1)
	go contaNegativos(gerador1)
	go contaParesImpares(gerador1)
	go validaPrint()

	fin = make(chan bool)
	<- fin
}
