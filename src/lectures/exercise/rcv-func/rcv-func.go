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
	health, maxHealth uint
	energy, maxEnergy uint
	name              string
}

func (player *Player) modifyHealth(health uint) {
	player.health += health
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}

	fmt.Println("Health changed to", player.health)
}

func (player *Player) modifyEnergy(energy uint) {
	player.energy += energy
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println("Energy changed to", player.energy)
}

func main() {
	player := Player{
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
		name:      "Kobe",
	}

	player.modifyHealth(10)
	player.modifyEnergy(10)
}
