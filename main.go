package main

import "fmt"
import (
	"math/rand"
	"time"
)


const(
	door1 = 1 + iota
	door2
	door3
)

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	// Generate a new seed each time

	stay_correct_count := 0
	change_correct_count := 0

	N_trys := 1000000
	for ii:=0 ; ii<N_trys; ii++ {
		rand.Seed(time.Now().UTC().UnixNano())

		available_doors := []int{door1, door2, door3}

		// Set the winning door
		winning_door := available_doors[rand.Intn(len(available_doors))]
		// Contestant picks one of the doors
		picked_door := available_doors[rand.Intn(len(available_doors))]

		// Open the door that isn't winning and not picked by the contestant
		opened_door := 0
		for d:=0; d <= door3; d++ {
			if (d != winning_door && d != picked_door) {
				opened_door = d
			}
		}

		remaining_doors := make([]int, 0)
		for d:=door1; d <= door3; d++ {
			if d != opened_door {
				remaining_doors = append(remaining_doors, d)
			}
		}

		// Simulate staying with the first pick
		if (picked_door == winning_door) {
			stay_correct_count ++
		}

		// Simulate picking the other door
		for _, d := range remaining_doors {
			if (d != picked_door) {
				picked_door = d
				break
			}
		}

		if (picked_door == winning_door) {
			change_correct_count ++
		}

	}


	fmt.Printf("Staying with the initally picked door worked out %v of %v\n", stay_correct_count, N_trys)
	fmt.Printf("Changing doors worked out %v of %v\n", change_correct_count, N_trys)

}