// Marcelo Heredia e Pedro Castro
package main

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

type Node struct {
	id int
	msg Message
	pos GridPosition
	channel chan Message
}

type Message struct {
	id int
	origin int
	text string
	steps int
}

type GridPosition struct{
	row,col int
}

type Report struct {
	node int
	steps int
}


const S = 10 //rede = SxS (S² >= N)
const N = 9 //numero de nodos
const RANGE = 1 //alcance de visão em todas direcoes do nodo

var end chan bool	

var nodes [N]Node
func main(){
	rand.Seed(time.Now().UnixNano())
	rt := make (chan Report)

	//start nodes
	for i:=0; i<N; i++{
		nodes[i].id = i;
		nodes[i].msg = Message{-1,-1,"",-1}
		nodes[i].channel = make(chan Message, 5)
		nodes[i].pos = GridPosition{rand.Intn(S),rand.Intn(S)} //put the node in a random position of the grid

		go processNodes(i, rt)
	}

	fmt.Println("processos inicializados")

	nodes[0].channel <- Message{0,0,"msg",0}

	//wait for reports
	var reports [N]Report
	
	for i:=0; i<N; i++{
		ret := <- rt
		fmt.Println("Passos dados: ",ret.steps)
		reports[ret.node] = ret
	}
	
	fmt.Println("\n\n\nRetornos: ")
	for i:=0; i<len(reports); i++{
		fmt.Println("nodo: ",reports[i].node,
			    "\npassos: ",reports[i].steps)
	}
}

func randomizeMovement(current GridPosition){
	signalY := 0 // even
	if time.Now().UnixNano()%2 != 0{
		signalY = 1 // odd
	}
	rand.Seed(time.Now().UnixNano())
	incrementY := rand.Intn(RANGE);

	signalX := 0
	if time.Now().UnixNano()%2 != 0{
		signalX = 1
	}
	rand.Seed(time.Now().UnixNano())
	incrementX := rand.Intn(RANGE);

	if incrementY == 0 &&
		incrementX == 0 {
			incrementX++
			incrementY++ //avoinding them to simply stand by
	}
	
	if signalY == 1{//they can go backwards
		incrementY = -incrementY
	}
	if signalX == 1{
		incrementX = -incrementX
	}

	//if passing the limits of the grid -> goes backwards
	futureX := current.col + incrementX
	if futureX < 0 || futureX > S{
		futureX = current.col - incrementX
	}
	
	futureY := current.row + incrementY
	if futureY < 0 || futureY > S{
		futureX = current.row - incrementY
	}

	current.col = futureX
	current.row = futureY
	
}

func findNodes(node_id int)[]int{
	this_row := nodes[node_id].pos.row
	this_col := nodes[node_id].pos.col

	var ngb []int

	for i:=0; i<N; i++{
		row := nodes[i].pos.row
		col := nodes[i].pos.col
		if Abs(this_row - row)<=RANGE &&
		   Abs(this_col - col)<=RANGE{
			ngb = append(ngb, i)
		}
	}

	return ngb
}

func processNodes (id int, ret chan Report){
    go func(){
        for {
            time.Sleep(time.Millisecond*time.Duration((id*10)+100))
            
            ngb := findNodes(id)
            fmt.Println("vizinhos: ",ngb)//vizinhos

            if nodes[id].msg.origin > -1{
                //tem msg valida
                for j:=0; j<len(ngb); j++{
                    nodes[ngb[j]].channel <- Message{nodes[id].msg.id, nodes[id].msg.origin, nodes[id].msg.text, nodes[id].msg.steps}
                }
            }
        }
    }()
    go func(){
        rcv := <-nodes[id].channel
        rcv.steps++
        fmt.Println("Mensagem recebida: ",printMsg(rcv))

        if nodes[id].msg.origin != rcv.origin{//nova msg
            nodes[id].msg = Message{rcv.id, rcv.origin, rcv.text, rcv.steps}
            ret <- Report{id, nodes[id].msg.steps}
        }
    }()
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printMsg(m Message) string{
	return "id: "+strconv.Itoa(m.id)+"\norigem: "+strconv.Itoa(m.origin)+
			"\ntexto: "+m.text+"\npassos: "+strconv.Itoa(m.steps)
}