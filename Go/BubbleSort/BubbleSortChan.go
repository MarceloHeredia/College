//Marcelo Heredia

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 20    //tamanho da lista
const MAX = 1000 // valores de 0..MAX
var fin chan struct{}

func celulaBSort(id int, meuValor *int, entradaE chan int, saidaD chan int, entradaD chan int, saidaE chan int) {
	passaE := *meuValor
	passaD := *meuValor
	for i:= 0; i<N; i++{
		
		saidaD <- passaD
		saidaE <- passaE

		lft := <- entradaE
		rgt := <- entradaD

		l,m,h := ordena(lft,*meuValor,rgt)

		if id == 0{ //se estou na primeira posicao do vetor
			*meuValor = l //meu valor deve ser o menor
			passaD = m //o valor a minha direita deve ser o do meio
			passaE = h //o valor a minha esquerda deve ser o maior, afinal e a ultima posicao do vetor
		}else if id==N{ //se estou na ultima posicao do vetor
			*meuValor = h //meu valor deve ser o maior
			passaD = l //o valor a minha direita deve ser o menor, afinal e a primeira pos do vetor
			passaE = m // o valor a minha esquerda deve ser o do meio
		}else{// se estou em uma posicao qqr no meio
			passaE = l
			*meuValor = m
			passaD = h 
		}

	}
	fin <- struct{}{}
}

func ordena(a,b,c int)(l,m,h int){
	if a < b && a<c{
		l = a
		if b < c{
			m = b
			h = c
		}else{
			m = c
			h = b
		}
	}else if b < a && b < c{
		l = b
		if a < c{
			m = a
			h = c
		}else {
			m = c
			h = a
		}
	}else{
		l = c
		if a < b{
			m = a
			h = b
		}else{
			m = b
			h = a
		}
	}
	return l,m,h
}

func print(list [N]int) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fin = make(chan struct{})

	lista := [N]int{} // preenche lista
	for i := 0; i < N; i++ {
		lista[i] = rand.Intn(MAX)
	}
	print(lista)

	var paraDireita [N]chan int // declara e inicia canais
	var paraEsquerda [N]chan int
	for i := 0; i < N; i++ {
		paraDireita[i] = make(chan int, 1)
		paraEsquerda[i] = make(chan int, 1)
	}

	//posicao 0 conecta com a ultima posicao do vetor
	go celulaBSort(0 ,&(lista[0]), paraDireita[0], paraDireita[1], paraEsquerda[0], paraEsquerda[N-1])
	for i := 1; i < N-1; i++ { // monta processos
		go celulaBSort(i, &(lista[i]), paraDireita[i], paraDireita[i+1], paraEsquerda[i], paraEsquerda[i-1])
	}
	//ultima posicao do vetor conecta com a inicial
	go celulaBSort(N-1,&(lista[N-1]), paraDireita[N-1], paraDireita[0], paraEsquerda[N-1], paraEsquerda[N-2])
	
	for i := 0; i < N; i++ {
		<-fin
	}
	print(lista)
}
