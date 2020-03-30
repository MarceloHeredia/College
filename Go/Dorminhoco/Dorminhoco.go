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

type Move struct {
	yourTurn bool
	player Player
	pass Card
}

var cards = [TOTAL_CARDS]Card { //starts with joker
	Card{
		number: 0,
		suit: "Joker",
	},
}

//create the card numbers 1-4 and 3 suits and add to the deck that already has the joker
func createCards() {
	var cards_pos = 1
	for _, j:= range suits{
		for i:=1; i<=BGC_NUMBER; i++{
			var card = Card{i,j}
			cards[cards_pos] = card
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
		for j:= 0; j<MAX_CARDS-1; j++{
			this_hand[j] = cards[countCards]
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

func fillMove (mov chan Move, player Player, isFirst bool){
	var mv = Move{
		yourTurn:isFirst,
		player:player,
	}
	mov <- mv
}

func main(){
	createCards()
	shuffleCards()
	for i:=0; i<TOTAL_CARDS; i++{
		fmt.Println(cards[i])
	}
	var players = initializePlayers()
	var chanRing [NUM_PLAYERS]chan Move
	var chanPass chan Pass
	for i:=0;i<NUM_PLAYERS; i++{
		var fst = false
		if (i==0){
			fst = true
		}
		fillMove(chanRing[i],players[i],fst)
	}

	go play(true, chanPass, chanRing[0], chanRing[1])

	for i:=1; i<(NUM_PLAYERS-1); i++{
		go play(false, chanPass, chanRing[i], chanRing[i+1])
	}

	go play(false,chanPass, chanRing[NUM_PLAYERS-1], chanRing[0])

	fin := make(chan struct{})
	<-fin
}

func play(yourTurn bool, pass chan Pass,
		  ringMy chan Move, ringNext chan Move){
			  for {
				  if yourTurn{

				  }else {//not your turn
					receiveCard(pass, ringMy)
				}
			  }
}

//auto selects a card to discard from your hand and then remove it from the hand array then re-sorts the array
func discardCard(chan Move) chan Card{

}

func receiveCard(pass chan Pass, move chan Move){
	p := <- pass
}