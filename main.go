package main

import "fmt"

type Player struct {
	health int
}

// Value receiver (doesn't modify original)
func (p Player) getHealth() int {
	return p.health
}

// Pointer receiver (can modify original)
func (p *Player) takeDamage(damage int) {
	p.health -= damage
}

func takeDamage(p *Player, damage int) {
	fmt.Println("Playes is taking damage from explosion")
	p.health -= damage
}

func main() {
	player := Player{health: 100}

	fmt.Println(player.getHealth()) // 100

	player.takeDamage(30)
	fmt.Println(player.getHealth()) // 70

	player.takeDamage(&player, 20)
}
