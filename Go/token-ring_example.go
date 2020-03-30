//example from the material of a graduation discipline 

package main

const N = 4

type Packet struct {
	sender int
	receiver int
	message string
}

type Token struct {
	empty bool
	packet Packet
}

func sender(id int, send chan Packet) {
	for i := 0; i <= N; i++ {
		send <- Packet{id, i, "msg"}
	}
}

func receiver(id int, rec chan Packet) {
	for {
		p := <-rec
		println("Pacote recebido : ",
		id, p.sender, p.receiver, p.message)
	}
}

func main() {
	var chanRing [N]chan Token
	var chanSend [N]chan Packet
	var chanRec [N]chan Packet
	for i:=0; i<N; i++{
		chanRing[i] = make(chan Token)
		chanSend[i] = make(chan Packet)
		chanRec[i] = make(chan Packet)
	}
	for i:=0; i<(N-1); i++{
		go node(i, false, chanSend[i],
				chanRec[i], chanRing[i], chanRing[i+1])
	}
	go node(N-1, true, chanSend[N-1],
			chanRec[N-1], chanRing[N-1], chanRing[0])
	fin := make(chan struct{})
	<-fin
}

func node(id int, hasToken bool, send chan Packet,
		  receive chan Packet, ringMy chan Token, ringNext chan Token){
			  println("node ",id)
			  go sender(id, send)
			  go receiver(id, receive)
			  for {
				  if hasToken{
					  p := Packet{0,0,"no"}
					  t := Token{true,p}
					  select {//has a token
					  case p:= <-send: { //want to send
						if p.receiver == id{
							receive <- p
							t := Token{true,p}
							ringNext <- t
						} else {
							t := Token{false,p}
							ringNext <- t
						}
						hasToken = false
					  }
					case ringNext <- t: //dont want to send
						hasToken = false
					}
				  } else { //dont have token
					t := <- ringMy
					if t.empty { //received token
						if t.packet.sender == id{
							ringNext <- Token{true,t.packet}
						} else {
							if t.packet.receiver == id {
								receive <- t.packet
							}
							ringNext <- t
						}
					}
				}
			  }
		  }
