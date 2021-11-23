package main

import (
	"fmt"
	"log"
)

// make abstract factory
type iSportFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportFactory,error) {
	switch {
	case brand == "addidas":
		return &addidas{},nil
	case brand == "nike":
		return &nike{},nil
	default:
		return nil,fmt.Errorf("Wrong brand type passed")
	}
}

// make addidas concrete factory

type addidas struct {

}

func (a *addidas)makeShoe() iShoe{
	return &addidasShoe{
		shoe:	shoe{
			logo:"addidas",
			size: 14,
		},
	}
}

func (a *addidas)makeShirt() iShirt {
	return &addidasShirt{
		shirt:	shirt{
			logo:"addidas",
			size:14,
		},
	}
}

// make nike concrete factory

type nike struct {

}

func (n *nike)makeShoe() iShoe{
	return &nikeShoe{
		shoe:	shoe{
			logo:"nike",
			size:14,
		},
	}
}

func (n *nike)makeShirt() iShirt {
	return &nikeShirt{
		shirt:	shirt {
			logo:"nike",
			size:14,
		},
	}
}

// make abstract product

type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo()	string
	getSize()	int
}

type shoe struct {
	logo	string
	size 	int
}

func (s *shoe)setLogo(logo string) {
	s.logo = logo
}

func (s *shoe)setSize(size int) {
	s.size = size
}

func (s *shoe)getLogo() string {
	return s.logo
}

func (s *shoe)getSize() int {
	return s.size
}

type iShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo()	string
	getSize()	int
}

type shirt struct {
	logo	string
	size	int
}

func (t *shirt)setLogo(logo string) {
	t.logo = logo
}

func (t *shirt)setSize(size int) {
	t.size = size
}

func (t *shirt)getLogo() string{
	return t.logo
}

func (t *shirt)getSize() int{
	return t.size
}

// make addidas concrete product

type addidasShoe struct {
	shoe
}

type addidasShirt struct {
	shirt
}

// make nike concrete product

type nikeShoe struct {
	shoe
}

type nikeShirt struct {
	shirt
}


func main() {
	// for nike factory
	// just have to fill setter methods for nike factory
	nike,err := getSportsFactory("nike")
	if err != nil {
		log.Panic(err.Error())
	}
	nikeShoe := nike.makeShoe()
	nikeShirt := nike.makeShirt()
	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)

	// for addidas factory
	// after that have to fill setter methods for addidas factory
	addidas,err := getSportsFactory("addidas")
	if err != nil {
		log.Panic(err.Error())
	}
	addidasShoe := addidas.makeShoe()
	addidasShirt := addidas.makeShirt()
	printShirtDetails(addidasShirt)
	printShoeDetails(addidasShoe)
}


func printShirtDetails(shirt iShirt) {
	fmt.Printf("logo : %s",shirt.getLogo())
	fmt.Println()
	fmt.Printf("size : %d",shirt.getSize())
	fmt.Println()
}

func printShoeDetails(shoe iShoe) {
	fmt.Printf("logo : %s",shoe.getLogo())
	fmt.Println()
	fmt.Printf("size : %d",shoe.getSize())
	fmt.Println()
}


