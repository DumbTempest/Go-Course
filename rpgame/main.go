package main

import (
	"fmt"
	"slices"
)

type Player struct {
	Name      string
	Inventory []string
}

type item struct {
	Name string
}

func (p1 *Player) AddItem(itemName string) {
	p1.Inventory = append(p1.Inventory, itemName)
}

func (p1 *Player) RemoveItem(itemName string) {
	p1.Inventory = slices.Delete(p1.Inventory, slices.Index(p1.Inventory, itemName), slices.Index(p1.Inventory, itemName)+1)
}

func main() {
	player := Player{
		Name:      "Hero",
		Inventory: []string{"Sword"},
	}

	fmt.Println("Before adding item:", player.Inventory)
	player.AddItem("Shield")
	fmt.Println("After adding item:", player.Inventory)

}
