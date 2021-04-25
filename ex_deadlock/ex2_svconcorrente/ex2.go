// Marcelo Heredia e Pedro Castro
package main

import(
    "fmt"
    "time"
)
const N int = 10

func main(){
   var cli[N] chan chan string


   for i:=0; i<N; i++{
       cli[i] = make (chan chan string)
       go client(i, cli[i])
   }

   go server(cli)
   time.Sleep(time.Second*10)

}


func client(id int, request chan chan string){
    for{
        fmt.Println("Client: ",id," made a request to the server")
        responseChan := make (chan string)
        request <- responseChan

        response := <- responseChan

        fmt.Println("Response for client ",id," : ", response)
    }
}

func server(clients [10]chan chan string){
    for{
        for i:=0; i<len(clients);i++{
            select{
            case hasRequest := <- clients[i]:
                clients[i] <- hasRequest
                go worker(clients[i])
            }
        }
    }
}

func worker(request chan chan string){
    responseChan := <- request
    responseChan <- time.Now().String()
}