package main

import (
	"fmt"
	"math/rand"
	"time"
)

func player() chan string {
	output := make(chan string)
	count := rand.Intn(100)
	move := []string{"UP", "DOWN", "LEFT", "RIGHT"}

	go func() {
		defer close(output)
		for i := 0; i < count; i++ {
			output <- move[rand.Intn(4)]
			d := time.Duration(rand.Intn(200))
			time.Sleep(d * time.Millisecond)
		}
	}()
	return output
}

func play(player int, movement string, moreData bool, activePlayers *int) {
	if moreData {
		fmt.Printf("Player %d: %s\n", player, movement)
	} else {
		*activePlayers--
		fmt.Printf("Player %d left the game. Remaining players: %d\n", player, *activePlayers)
	}
}

func main() {
	player1Ch := player()
	player2Ch := player()
	player3Ch := player()
	player4Ch := player()
	activePlayers := 4

	for activePlayers > 1 {
		select {
		case m1, moreData := <-player1Ch:
			play(1, m1, moreData, &activePlayers)
		case m2, moreData := <-player2Ch:
			play(2, m2, moreData, &activePlayers)
		case m3, moreData := <-player3Ch:
			play(3, m3, moreData, &activePlayers)
		case m4, moreData := <-player4Ch:
			play(4, m4, moreData, &activePlayers)
		}
	}

	fmt.Println("Game finished")
}
