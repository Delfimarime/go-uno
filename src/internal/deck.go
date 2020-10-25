package internal

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type Color string

const (
	Red    Color = "R"
	Green  Color = "G"
	Blue   Color = "B"
	Yellow Color = "Y"
	Black  Color = "L"
)

type Symbol string

const (
	Zero   Symbol = "0"
	One    Symbol = "1"
	Two    Symbol = "2"
	Three  Symbol = "3"
	Four   Symbol = "4"
	Five   Symbol = "5"
	Six    Symbol = "6"
	Seven  Symbol = "7"
	Eight  Symbol = "8"
	Nine   Symbol = "9"
	Skip   Symbol = "S"
	Revert Symbol = "R"
	Wild   Symbol = "W"
	Draw2  Symbol = "+2"
	Draw4  Symbol = "+4"
)

type Card interface {
	GetId() string

	GetColor() Color

	GetSymbol() Symbol
}

type Deck struct {
	in    []Card
	out   []Card
	ready bool
}

func (instance *Deck) Draw(size int) []Card {
	return nil
}

func (instance *Deck) AfterPropertiesSet() {
	if !instance.ready {
		content := make([]Card, 108)

		content[0] = newCard(Zero, Red)
		content[1] = newCard(Zero, Green)
		content[2] = newCard(Zero, Blue)
		content[3] = newCard(Zero, Yellow)

		content[4] = newCard(Wild, Black)
		content[5] = newCard(Wild, Black)
		content[6] = newCard(Wild, Black)
		content[7] = newCard(Wild, Black)

		content[8] = newCard(Draw4, Black)
		content[9] = newCard(Draw4, Black)
		content[10] = newCard(Draw4, Black)
		content[11] = newCard(Draw4, Black)

		withColor := []Symbol{One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Skip, Revert, Draw2}
		colors := []Color{Red, Green, Blue, Yellow}

		var yndex int = 12

		for position := 0; position < len(withColor); position++ {
			for _, color := range colors {
				for index := 0; index < 2; index++ {
					content[yndex] = newCard(withColor[position], color)
					yndex += 1
				}
			}
		}

		shuffle(content)

		instance.in = content
		instance.out = make([]Card, 0)
		instance.ready = true
	}
}

func newCard(symbol Symbol, color Color) Card {
	return CardModel{id: uuid.New().String(), symbol: symbol, color: color}
}

func (instance *Deck) Size() int {
	return len(instance.in)
}

func (instance *Deck) IsSizeGreaterOrEqualTo(size int) bool {
	return instance.Size() >= size
}

func (instance *Deck) IsSizeLesserOrEqualTo(size int) bool {
	return instance.Size() <= size
}

func shuffle(content []Card) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(content) > 0 {
		n := len(content)
		randIndex := r.Intn(n)
		content[n-1], content[randIndex] = content[randIndex], content[n-1]
		content = content[:n-1]
	}
}

type CardModel struct {
	id     string
	color  Color
	symbol Symbol
}

func (instance CardModel) GetId() string {
	return instance.id
}

func (instance CardModel) GetSymbol() Symbol {
	return instance.symbol
}

func (instance CardModel) GetColor() Color {
	return instance.color
}
