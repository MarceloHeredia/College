//Marcelo Heredia


package main

import (
	"fmt"
	"math/rand"
	"time"
	"../MCCSemaforo"
)

const N = 20 //tamanho do vetor
var vet []int



func main(){
	rand.Seed(time.Now().UnixNano())

	//cria um semaforo para aguardar o fim da execucao do programa
	var end = MCCSemaforo.NewSemaphore(0)


	vet = createVector(N) 

	fmt.Println("vetor inicial: ",vet)

	go quickSort(vet, 0, len(vet)-1,end)
	
	end.Wait()//aguarda o fim da execucao do quicksort

	fmt.Println("vetor final: ",vet)
}

//vetor utilizando slices
func createVector(size int) []int{
	var vet []int
	
	for i:=0; i<N; i++{
		vet = append(vet,rand.Intn(999))
	}
	return vet
}

//alem do normal do quicksort, recebe como parametro o semaforo da chamada anterior
func quickSort(a []int, lo, hi int, sf *MCCSemaforo.Semaphore){
	if lo < hi {
		//declara dois semaforos, um para cada chamada concorrente do quicksort
		var sem1 = MCCSemaforo.NewSemaphore(0)
		var sem2 = MCCSemaforo.NewSemaphore(0)
		pivot := partition(a, lo, hi)

		//executa chamada concorrente dos dois quicksorts 
		go quickSort(a, lo, pivot-1,sem1)
		go quickSort(a, pivot+1, hi,sem2)
	
		//aguarda as duas chamadas de quicksort retornarem
		sem1.Wait()
		sem2.Wait()
	}
	//por fim, da signal e retorna
	sf.Signal()
	return
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] <= pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}