// Marcelo Heredia e Pedro Castro
package main

import (
	"math/rand"
	"time"
)

var end chan bool
const Nc = 10 //numero de clientes

func cliente(id int, pedido chan bool){
	for{
		rand.Seed(time.Now().UnixNano())
		nmbr := rand.Intn(1)+1 //escolha aleatoria de pedido (horario atual ou valor aleatorio na tela)
		servidor(id,pedido,nmbr)
	}
	end <- true
}

func servidor(id_cli int, atendido chan bool, req int){
	
}

func dizerHorario(id_cli int, atendido chan bool){
	println("Cliente ",id_cli, " horario atual: ", time.Now().String())
}

func printAleatorio(id_cli int, atendido chan bool){
	println("Cliente ",id_cli," valor aleatorio: ", rand.Intn(999)+1)
}

func main(){
	var clientes[Nc] chan bool
	for i:=0; i<Nc; i++{
		clientes[i] = make (chan bool)
		clientes[i] <- false
		go cliente(i, clientes[i])
	}

	end = make (chan bool)
	<- end

}