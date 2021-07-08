package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/google/uuid"
	"encoding/json"
)

type Deck struct {
	ID 		  string `json: "deck_id"`
	Shuffled  bool 	 `json: "shuffled"`
	Remaining int 	 `json: "remaining"`
	Cards 	  []
}

type Card struct {
	Value string `json: "value"`
	Suit  string `json: "suit"`
	Code  string `json: "code"`
}

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string("Diamonds", "Hearts", "Clubs", "Spades")
	cardValues := []string("Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King")

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			code := fmt.Sprintf("%v of %s", val, suit)
			card := Card{Value: val, Suit: suit, Code: code}
			cards = append(cards, card)
		}
	}
	return cards
}

func openDeck(w http.ResponseWriter, r *http.Request){


	fmt.Fprintf(w, "Opened a deck!")
}

func drawCard(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Draw a card")
}
 
func createDeck(w http.ResponseWriter, r *http.Request){
	id := uuid.New()
	deck := Deck{ID: id.String(), Shuffled: false, Remaining: 52}
	json.NewEncoder(w).Encode(deck)
}

func main(){
	http.HandleFunc("/", createDeck)
	http.HandleFunc("/open", openDeck)
	http.HandleFunc("/draw", drawCard)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

