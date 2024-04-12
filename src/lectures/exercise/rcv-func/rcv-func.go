//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	health int
	energy int
	name   string
}

func (player *Player) modifyHealth(health int) {
	player.health += health
	fmt.Println("Health changed to", player.health)
}

func (player *Player) modifyEnergy(energy int) {
	player.energy += energy
	fmt.Println("Energy changed to", player.energy)
}

func main() {
	player := Player{70, 100, "Player1"}

	player.modifyHealth(10)
	player.modifyEnergy(-10)
}
