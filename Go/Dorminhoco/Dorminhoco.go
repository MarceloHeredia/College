//Marcelo Heredia
/*
* for this problem i'm considering 0 and empty string as joker
* there will be 3 suits of cards on the game(hearts, diamonds, spades) for the simplification 
* the card number will be one, two and three and four only (0 will be the joker)
*/
package main

import (
	"fmt"
	"math/rand"
	//"sync"
	"time"
)

//constants
const NUM_PLAYERS = 4
const MAX_CARDS = 5
const BGC_NUMBER = 4
const TOTAL_CARDS = 13
var suits = []string{"hearts", "diamonds", "spades"}
var names = []string{"John", "Emma", "Lou", "Mia"}

//structs
type Card struct {
	number int
	suit string
}

type Pass struct{
	card Card
	target Player
}

type Player struct {
	name string
	hand []Card
	lockJoker bool
}

func (p Player) HasJocker() bool {
    for i:=0; i<len(p.hand); i++{
		if p.hand[i].number == 0{
			return true
		}
	}
	return false
}

type Move struct {
	yourTurn bool
	player Player
	pass Card
}

var cards = []Card { //starts with joker
	Card{
		number: 0,
		suit: "Joker",
	},
}

//create the card numbers 1-4 and 3 suits and add to the deck that already has the joker
func createCards() {
	var cards_pos = 0
	for _, j:= range suits{
		for i:=1; i<=BGC_NUMBER; i++{
			var card = Card{i,j}
			cards = append(cards, card)
			cards_pos++
		}
	}
}

//shuffles the cards for more randomness
func shuffleCards(){
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<TOTAL_CARDS; i++{
		var rd = rand.Intn(TOTAL_CARDS-1)+1
		var aux = cards[i]
		cards[i] = cards[rd]
		cards[rd] = aux
	}
}

//initializes an array of players
func initializePlayers() [NUM_PLAYERS]Player{
	var players [NUM_PLAYERS]Player
	var countCards = 0
	for i:=0; i<NUM_PLAYERS; i++{
		var hasJoker = false
		var this_hand []Card
		for j:= 0; j<MAX_CARDS-2; j++{
			this_hand = append(this_hand, cards[countCards])
			if cards[countCards].number == 0{
				hasJoker = true
			}
			countCards++
		}
		players[i] = Player{
			name: names[i],
			hand: this_hand,
			lockJoker: hasJoker,
		}
	}
	//giving the first player of the circle the extra card and allowing him to pass one card
	players[0].hand = append(players[0].hand, cards[countCards])

	return players
}

func fillMove (mov chan Move, player Player, fstPlayer bool){
	var mv = Move{
		yourTurn:fstPlayer,
		player:player,
	}
	mov <- mv
}

func main(){
	createCards()
	shuffleCards()
	var players = initializePlayers()
	var chanRing [NUM_PLAYERS]chan Move
	var chanPass [NUM_PLAYERS]chan Pass
	for i:=0;i<NUM_PLAYERS; i++{
		var fst = false
		if (i==0){
			fst = true
		}
		fillMove(chanRing[i],players[i], fst)
		chanRing[i] = make (chan Move)
		chanPass[i] = make (chan Pass)
	}

	go play(chanPass[0], chanRing[0], chanRing[1])

	for i:=1; i<(NUM_PLAYERS-1); i++{
		go play(chanPass[i], chanRing[i], chanRing[i+1])
	}

	go play(chanPass[NUM_PLAYERS-1], chanRing[NUM_PLAYERS-1], chanRing[0])

	fin := make(chan struct{})
	<-fin
}

func play(pass chan Pass,
		  ringMy chan Move, ringNext chan Move){
			  for {
				  tstTurn := <- ringMy
				  if tstTurn.yourTurn{
				  	ringMy <- tstTurn
					sendCard(ringMy,pass,ringNext)
				  }else {//not your turn apparently
					ringMy <- tstTurn
					receiveCard(pass, ringMy)
				}
			  }
}

//auto selects a card to discard from your hand and then remove it from the hand array then re-sorts the array
func sendCard(move chan Move, pass chan Pass, nextMove chan Move){
	myMv := <- move
	var cardToDiscard = discard(myMv.player.hand, myMv.player.lockJoker)
	fmt.Println("Jogador(a): ",myMv.player.name, "\n Passou adiante a carta: ",cardToDiscard)
	move <- myMv
	nxtMov := <- nextMove
	var snd = Pass{cardToDiscard,nxtMov.player}
	nxtMov.yourTurn = true
	nextMove <- nxtMov
	pass <- snd
}

func discard(hand []Card, canDiscardJoker bool) Card{
	var nmbrs = []int{0,0,0,0,0}
	for i:=0; i<len(hand); i++{
		nmbrs[hand[i].number]++
	}
	var indexRemoval int 
	var cardRemoved Card
	var hasRemoved = false
	if nmbrs[0] == 1 && canDiscardJoker{ //if are going to pass the joker
		for i:=0; i<len(hand);i++{
			if hand[i].number == 0{
				indexRemoval = i
				cardRemoved = hand[i]
				hasRemoved = true
			}
		}
	}else { //are going to pass the less frequent card
		var numForRemov int
		for i:=1; i<len(nmbrs); i++{//start by position 1 because u cant pass joker when didnt fall on first if
			if (nmbrs[i] == 1 ){
				numForRemov = i
			}
		}
		for i:=0; i<len(hand); i++{ 
			if(hand[i].number == numForRemov){
				indexRemoval = i
				cardRemoved = hand[i]
				hasRemoved = true
			}
		}
	}
	if !hasRemoved{//there are no less frequent card, for now i'll use a random removal
		indexRemoval = rand.Intn(len(hand)-1)+1
		cardRemoved = hand[indexRemoval]
		hasRemoved = true
	}
	hand = append(hand[:indexRemoval] , hand[indexRemoval+1:]...)
	return cardRemoved
}

func receiveCard(pass chan Pass, move chan Move){
	p := <- pass
	mv := <- move
	var person = mv.player
	if p.target.name == person.name{
		fmt.Println("Jogador(a): ",person.name,"\n Recebeu a carta: ", p.card)
		if person.HasJocker() {
			person.lockJoker = false
		}
		person.hand = append(person.hand, p.card)
		if p.card.number == 0{
			person.lockJoker = true
		}

	}
	move <- mv
}