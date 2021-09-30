package main

import (
	"fmt"
	"golang-design-patterns/abstract_factory"
)

func main() {
	// abstract factory method
	err := abstractFactoryMethod()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func abstractFactoryMethod() error {
	adidasFactory, err := abstract_factory.GetSportsFactory("adidas")
	if err != nil {
		return err
	}
	nikeFactory, err := abstract_factory.GetSportsFactory("nike")
	if err != nil {
		return err
	}
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
	return nil
}

func printShoeDetails(s abstract_factory.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

//
func printShirtDetails(s abstract_factory.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
