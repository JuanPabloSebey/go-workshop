package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ball := make(chan int)
	winner := make(chan bool)

	go player(ball, winner, "Delpo", .95)
	go player(ball, winner, "Djokovic", .75)

	ball <- 0
	<- winner

}

func player(ball chan int, winner chan bool, name string, skill float32) {
	for hit :=range ball {

		shot := getShot(hit)

		if rand.Float32() > skill {
			fmt.Printf("%v throw the ball out of the court with a %v\n", name, shot)
			winner <- true
			return
		}

		fmt.Printf("%v pass the ball with a %v\n", name, shot)

		time.Sleep(time.Second)
		ball <- hit + 1
	}
}

func getShot(hitNumber int) string{
	if hitNumber == 0 {
		return 	"serve"
	}

	switch hitChance := rand.Float32(); {
	case hitChance < .1:
		return "gan willy"
	case hitChance < .4:
		return "volley shot"
	case hitChance < .7:
		return "forehand shot"
	default:
		return "backhand shot"
	}
}