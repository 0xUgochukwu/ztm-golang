// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package main

import "fmt"
import "testing"

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

func TestModifyHealth(t *testing.T) {
	player := Player{
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
		name:      "Kobe",
	}
	player.modifyHealth(10)

	if player.health > player.maxHealth {
		t.Errorf("Health should not exceed maxHealth")
	}
}

func TestModifyEnergy(t *testing.T) {
	player := Player{
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
		name:      "Kobe",
	}
	player.modifyEnergy(10)
	if player.energy > player.maxEnergy {
		t.Errorf("Energy should not exceed maxEnergy")
	}
}
