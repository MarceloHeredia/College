// Marcelo Heredia e Pedro Castro
package main

import(
    "fmt"
    "math/rand"
)
const N int = 10

func servidor(id int, req chan chan int){
   for{
       fmt.Println("Cliente: ", id, " Fez uma requisição")
       go trabalhador(<-req)
       fmt.Println(<-req)
   }
}

func trabalhador(req chan int){
    var numero int
    for {
        numero = rand.Intn(999)
        fmt.Println(numero)
        req <- numero
    }
}

func main(){
   var cli[N] chan chan int
    for i:=0; i < N; i++{
        cli[i] = make(chan chan int)
    }

    for i:= 0; ; i++{
        go servidor(i % N, cli[i % N])
    }
}