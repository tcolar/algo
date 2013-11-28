// History: Nov 26 13 tcolar Creation

package algo

import (
	"log"
	"math/rand"
	"time"
)

type Goalie struct {
	SavePct int
	// Channel that receives pucks from players
	// "handling" one shot at a time
	PuckChannel  chan string
	GoalsAgainst int
}

type Team struct {
	Goalie  Goalie
	Skaters [5]string
}

// Simuate 2 hockey teams of 5 skaters ecah shooting 10 pucks each (concurrently)
// aganst the an opposite goaie
func GameSim() {
	sharks_channel := make(chan string)
	penguins_channel := make(chan string)
	sharks := Team{Goalie{92, sharks_channel, 0}, [5]string{"Boyle", "Burns", "Couture", "Marleau", "Hertl"}}
	penguins := Team{Goalie{90, penguins_channel, 0}, [5]string{"Letang", "Malkin", "Dupuis", "Kunitz", "Neal"}}

	log.Print("Game started.")
	// Each skater starts shhoting
	for i := 0; i != 5; i++ {
		go StartShooting(sharks.Skaters[i], &penguins.Goalie)
		go StartShooting(penguins.Skaters[i], &sharks.Goalie)
	}
	go ReceiveShots(&sharks.Goalie)
	go ReceiveShots(&penguins.Goalie)
	time.Sleep(3 * time.Second)
	log.Print("Game completed.")
	log.Printf("Final score - Sharks: %d, Penguins: %d",
		penguins.Goalie.GoalsAgainst, sharks.Goalie.GoalsAgainst)
}

func ReceiveShots(goalie *Goalie) {
	var skater string
	ok := true
	for ok {
		// golaie recives shots
		if skater, ok = <-goalie.PuckChannel; ok {
			// randomly makes saves based on save percentage
			if rand.Int()%100 >= goalie.SavePct {
				goalie.GoalsAgainst++
				log.Printf(" ! GOAL from %s", skater)
			} else {
				//log.Print("Saved shot from %s", skater)
			}
		}
	}
}

func StartShooting(skater string, goalie *Goalie) {
	// shoot 10 pucks each
	for i := 1; i <= 10; i++ {
		// "random" delay between shots
		val := 12 + time.Duration(rand.Int()%100)
		time.Sleep(time.Millisecond * val)
		log.Printf("Shot %d from %s", i, skater)
		goalie.PuckChannel <- skater
	}
}
