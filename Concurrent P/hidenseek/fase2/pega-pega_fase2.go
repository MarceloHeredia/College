//Marcelo Heredia
package main

import (
	"math/rand"
	"time"
	"fmt"
)

//declaracao global
type GridPosition struct{ 
	x,y int
}

const N = 10 //area NxN do jogo
const F = 2 //numero de jogadores fugitivos
var end chan bool 
var limPassosPega = GridPosition{2,2} //limite de passos no eixo x e no eixo y
var limPassosFoge = GridPosition{1,1} 
var distanciaFim = GridPosition{2,2} //distancia minima no eixo x e no eixo y para fim do jogo

func procPega(minhaPos GridPosition, ultMovFugitivos []chan GridPosition, meuUltMov chan GridPosition){
	var posFugitivos []GridPosition
	fugPegos := 0
	for {
		for i:=0; i<F; i++{
			posFugitivos[i] = <- ultMovFugitivos[i]	//fica tentando ler do canal de movimentos do fugitivo
		}
	
		//ao ler, eh sua vez.

		minhaPos = movimentoPegador(minhaPos, posFugitivos) //executa o movimento

		printLocalizacoes(minhaPos, posFugitivos)

		for i:=0; i<F-fugPegos; i++{
			if isItOver(distanciaJogadores(minhaPos,posFugitivos[i])){
				fmt.Println("Pegador conseguiu alcancar um fugitivo ")
				fugPegos++
			}
		}

		for i:=0; i<F-fugPegos;  i++{
			meuUltMov <- minhaPos
		}
	}
//encerraPega: end <- true
}

func procFoge(id int, minhaPos GridPosition, ultMovPegador chan GridPosition, ultMovFugitivos []chan GridPosition){
	for{
		posPegador := <- ultMovPegador  //fica tentando ler do canal de movimentos do pegador

		//ao ler, eh sua vez
		if isItOver(distanciaJogadores(minhaPos,posPegador)){
			fmt.Println("Fugitivo percebeu que foi pego.")
			goto encerraFuga
		}
		
		minhaPos = movimentoFugitivo(minhaPos, posPegador) //executa o movimento

		ultMovFugitivos[id] <- minhaPos
	}
encerraFuga: end <- true
}

func main(){
	rand.Seed(time.Now().UnixNano())

	end = make(chan bool,F)
	var posPega = randLocal() //posicao do jogador que foge

	var posFoge []GridPosition
	for i:= range posFoge{
		posFoge[i] = randLocal2(posPega)
	}
	
	printLocalizacoes(posPega, posFoge)
	
	movPega := make(chan GridPosition,1) //contem a ultima posicao do jogador que pega
	var chFoge []chan GridPosition
	for i:=0; i<F; i++{
		chFoge[i] = make(chan GridPosition, F)
	}

	//Variante A
	go procPega(posPega, chFoge, movPega)//envia ao processo do jogador que pega o canal com movimentacoes do Fugitivo
	
	for i:=0; i<F; i++{
		go procFoge(i, posFoge[i], movPega, chFoge)// envia ao processo do jogador que foge o canal com movimentacoes do Pegador
	}

	for i:=0; i<F; i++{
		chFoge[i] <- posFoge[i]
	}

	for i:=0; i<F; i++{
		<-end
	}

	fmt.Println("Fim de Jogo!")
}

//executa um movimento do jogador fugitivo para longe do jogador que pega
func movimentoFugitivo(minhaPos, posPegador GridPosition)GridPosition{
	var movX,movY int

	//se ambos eixos se aproximam do limite ao mesmo tempo, escolher o mais distante do pegador
	// e retornar nessa direcao, sem mudar o outro eixo
	if minhaPos.x >= N-1 &&
		minhaPos.y >= N-1{
		if abs(minhaPos.x-posPegador.x) > abs(minhaPos.y-posPegador.y){//se o eixo mais distante for o X
			movX = minhaPos.x-limPassosFoge.x
			movY = minhaPos.y
		}else{
			movX = minhaPos.x
			movY = minhaPos.y-limPassosFoge.y
		}
	}else if minhaPos.x >= N-1 ||
			minhaPos.x <=1{// se apenas minha posicao no eixo X se aproxima de um dos limite
		movX = minhaPos.x
		if posPegador.y < minhaPos.y {// se o pegador estiver para tras no eixo Y
			if minhaPos.y + limPassosFoge.y <= N-1{//se eu posso dar todos passos possiveis sem chegar ao limite
				movY = minhaPos.y + limPassosFoge.y
			}else{//se nao posso dar todos passos possiveis sem chegar ao limite
				movY = N-1
			}
		}else{//se ele estiver na minha frente no eixo Y
			if minhaPos.y - limPassosFoge.y >= 1{ //se posso dar todos passos para tras sem chegar ao limite
				movY = minhaPos.y - limPassosFoge.y
			}
		}
	}else if minhaPos.y >= N-1||
			minhaPos.y <=1{//se apenas minha posicao no eixo Y se aproxima de um dos limite
		movY = minhaPos.y
		if posPegador.x < minhaPos.x {// se o pegador estiver para tras no eixo X
			if minhaPos.x + limPassosFoge.x <= N-1{//se eu posso dar todos passos possiveis sem chegar ao limite
				movX = minhaPos.x + limPassosFoge.x
			}else{//se nao posso dar todos passos possiveis sem chegar ao limite
				movX = N-1
			}
		}else{//se ele estiver na minha frente no eixo X
			if minhaPos.x - limPassosFoge.x >= 1{ //se posso dar todos passos para tras sem chegar ao limite
				movX = minhaPos.x - limPassosFoge.x
			}
		}
	}else if minhaPos.x <= 1 && minhaPos.y <=1{//estou me aproximando dos limites inferiores em ambos eixos
		if abs(minhaPos.x-posPegador.x) > abs(minhaPos.y-posPegador.y){//se o eixo mais distante for o X
			movX = minhaPos.y+limPassosFoge.x
			movY = minhaPos.y
		}else{
			movX = minhaPos.x
			movY = minhaPos.y+limPassosFoge.y
		}
	}else{//se nada se aproxima de limite algum
		if minhaPos.x < posPegador.x{
			if minhaPos.x - limPassosFoge.x >= 1{
				movX = minhaPos.x - limPassosFoge.x
			}else{
				movX = 1
			}
		}else{
			if minhaPos.x + limPassosFoge.x <= N-1{
				movX = minhaPos.x + limPassosFoge.x
			}else{
				movX = N-1
			}
		}
		if minhaPos.y < posPegador.y{
			if minhaPos.y - limPassosFoge.y >=1{
				movY = minhaPos.y - limPassosFoge.y
			}else{
				movY = 1
			}
		}else{
			if minhaPos.y + limPassosFoge.y <= N-1{
				movY = minhaPos.y + limPassosFoge.y
			}else{
				movY = N-1
			}
		}
	}

	return GridPosition{movX, movY}
}

//executa um movimento do jogador que pega na direcao do fugitivo
func movimentoPegador (minhaPos GridPosition, posFugitivos []GridPosition)GridPosition{
	//calcula numero de passos no eixo X
	var movX,movY int
	posFugitivo := posFugitivos[whoIsCloser(minhaPos, posFugitivos)]
	//se a distancia no eixo X for maior que o limite de passos
	//utiliza o limite de passos
	if abs(minhaPos.x - posFugitivo.x) > limPassosPega.x{ 
		movX = limPassosPega.x
	}else{//caso contrario anda a distancia necessaria apenas
		movX = abs(minhaPos.x-posFugitivo.x)
	}
	//calcula numero de passos no eixo Y
	//mesma regra do eixo X se aplica no eixo Y
	if abs(minhaPos.y - posFugitivo.y) > limPassosPega.y{
		movY = limPassosPega.y
	}else{
		movY = abs(minhaPos.y-posFugitivo.y)
	}
	
	//agora calcula a direcao desses passos no eixo X
	//se eu estiver além do fugitivo em algum dos eixos, ando para tras nesse eixo
	if minhaPos.x > posFugitivo.x{
		movX = minhaPos.x-movX
	}else{
		movX = minhaPos.x+movX
	}
	if minhaPos.y > posFugitivo.y{
		movY = minhaPos.y-movY
	}else{
		movY = minhaPos.y+movY
	}

	return GridPosition{movX,movY}
}

//define posicoes aleatorias para um jogador
func randLocal()GridPosition{

	x := rand.Intn(N)
	y := rand.Intn(N)

	return GridPosition{x,y}
}

//define posicoes aleatorias para um jogador
//garante q o jogador 2 nao saira na posicao do jogador 1
func randLocal2(fst GridPosition)GridPosition{
	gd := GridPosition{rand.Intn(N), rand.Intn(N)}

	for isItOver(distanciaJogadores(fst,gd)){
		gd = GridPosition{rand.Intn(N), rand.Intn(N)}
	}
	return gd;
}

//faz um print simples das localizacoes formatado para o fim de cada turno
func printLocalizacoes(posPega GridPosition, posFoge []GridPosition){
	var posPegaFormt = posPega.Posicao()

	var posFogeFormt []string
	for i:=0; i<F; i++{
		posFogeFormt[i] = posFoge[i].Posicao()
	}

	fmt.Println("Posicao jogador que pega: ",posPegaFormt)
	for i:=0; i<F; i++{
		fmt.Println("Posicao jogador que foge ",i,": ",posFogeFormt[i],
				"  Distância [x,y]:", distanciaJogadores(posPega,posFoge[i]).Posicao())
	}
}

//retorna a posicao de um jogador formatado [x,y]
func (p GridPosition) Posicao() string{
	return fmt.Sprintf("[%d,%d]",p.x,p.y)
}

//compara a GridPosition que representa a distancia entre os jogadores 
//com a GridPosition declarada nas constantes cujos numeros representam a distancia minima para pegar o fugitivo
func isItOver(test GridPosition) bool{
	if test.x <= distanciaFim.x && 
	   test.y <= distanciaFim.y{
		   return true
	   }
	return false
}

func whoIsCloser(me GridPosition, targets []GridPosition)int{
	search := 0
	xpos := abs(me.x - targets[0].x)
	ypos := abs(me.y - targets[0].y)

	for i:=1; i<F; i++{
		if abs(me.x - targets[i].x) < xpos &&
		   abs(me.y - targets[i].y) < ypos{
			   search = i
			   xpos = abs(me.x - targets[i].x)
			   ypos = abs(me.y - targets[i].y)
		   }
	}

	return search
}

//retorna uma GridPosition com o valor absoluto da diferenca entre as posicoes dos dois jogadores
func distanciaJogadores(p,q GridPosition) GridPosition{
	x := abs(p.x - q.x)
	y := abs(p.y - q.y)
	return GridPosition{x,y}
}

func abs(x int)int{
	if x<0{
		return -x
	}
	return x
}