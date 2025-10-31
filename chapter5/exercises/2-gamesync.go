package main

import (
	"fmt"
	"sync"
	"time"
)

func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int, timerExpired *bool) {
	cond.L.Lock()
	fmt.Println(playerId, ": Connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}

	for *playersRemaining > 0 && !*timerExpired {
		fmt.Println(playerId, ": Waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	if *timerExpired {
		fmt.Println(playerId, ": Game cancelled")
	} else {
		fmt.Println("All players connected. Ready player", playerId)
	}
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	timerExpired := false
	time.AfterFunc(3*time.Second, func() {
		cond.L.Lock()
		timerExpired = true
		cond.Broadcast()
		cond.L.Unlock()
	})

	playersInGame := 4
	for playerId := 0; playerId < 4; playerId++ {
		go playerHandler(cond, &playersInGame, playerId, &timerExpired)
		time.Sleep(1 * time.Second)
	}
}
