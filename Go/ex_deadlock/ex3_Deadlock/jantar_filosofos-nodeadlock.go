// Marcelo Heredia e Pedro Castro
package main

import (
	"time"
)

const PLAYERS = 5
const HUNGER = 3 //numero de vezes que cada um se alimenta
const EATnTHINK = time.Second/2 //tempo que cada filosofo gasta se alimentando

var end chan bool
var filosofos = []string{"Jorge", "Joao", "Michel", "Edgar", "Ernesto"} //nome dos 5 filosofos
var hungryLevels = []int {3,3,3,3,3} //nivel de fome para cada filosofo nomeado acima


func Diner(id int, left chan bool, right chan bool){
	for {
		//tempo pensando
		if hungryLevels[id] > 0 {
			
			println("Filosofo: ",filosofos[id], " esta pensando...")
			time.Sleep(EATnTHINK)
			println("Filosofo: ",filosofos[id], " esta faminto...")

			lft,rgt := false,false
			select {
			case lft = <- left: //tenta pegar o hashi da esquerda
				break
			default: 
			 lft = false
			 break
			}

			if !lft { //se nao conseguiu pegar o hashi a esquerda
				println("Filosofo: ",filosofos[id], "nao conseguiu pegar o hashi da esquerda")
				continue
			}

			select {
			case rgt = <- right: //tenta pegar o hashi da direita
				break
			default:
				rgt = false
				break
			}

			if !rgt{//se nao conseguir o hashi da direita (deve soltar o da esquerda tambem)
				left <- true
				println("Filosofo: ",filosofos[id], "nao conseguiu pegar o hashi da direita")
				continue
			}

			if lft && rgt {//se pegou os dois
				println("Filosofo: ",filosofos[id], "pegou ambos hashis e esta comendo")
				time.Sleep(EATnTHINK)
				//aguarda terminar de comer
				println("Filosofo: ",filosofos[id], "terminou de comer e largou os garfos")
				hungryLevels[id]--
				left <- true
				right <- true
			}
		}else if hungryLevels[id]==0{
			println("Filosofo: ",filosofos[id], "esta completamente alimentado e saiu da mesa")
			end <- true //avisa o main que este filosofo esta alimentado
			break
		}
	}
}


func main(){
	var hashi[PLAYERS] chan bool
	for i:=0; i<PLAYERS; i++{
		hashi[i] = make(chan bool, 1)
		hashi[i] <- true //true significa que o hashi esta livre
	}

	go Diner(0, hashi[0], hashi[PLAYERS-1])
	for i:=1; i<PLAYERS; i++{
		go Diner(i, hashi[i], hashi[i-1])
	}

	end = make(chan bool, 5)
	for i:=0; i<5; i++{
		<- end
	}
	
	println("Todos filosofos estão alimentados")
}

