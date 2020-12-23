package main

import (
	"fmt"
	"time"
)

/*
source: https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
Here is a program that mines ore. The functions in this example perform: finding ore, mining ore, and smelting ore.
In our example, the mine and ore are represented as an array of strings,
with each function taking in and returning a “processed” array of strings.
finder() ---arrayOfOre--> miner() ---arrayOfMinedOre-> smelter()
*/

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	// channels for go routines to communicate
	// <- points the direction data is going in
	// channel <- "hello" (send)
	// someVar := <- channel (receive)
	oreChan := make(chan string)
	minedOreChan := make(chan string)

	// go routines (independent working lightweight threads)
	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			oreChan <- item // send to oreChan
		}
	}(theMine)

	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan //receive from oreChan
			fmt.Println("Miner: Received " + foundOre + " from finder")
			minedOreChan <- "minedOre" // send to minedOreChan
		}
	}()

	// Smelter
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan // receive from minedOreChan
			fmt.Println("From Miner :", minedOre)
			fmt.Println("From Smelter : Ore smelted")
		}
	}()

	// Receiving from channel after 5 sec
	// As main() is a go routine
	// go routine block on a read until something is sent.
	// That’s exactly what is happening to the main routine by adding this code
	// The main routine will block, giving our other go routines 5 seconds of additional life to run.
	<-time.After(time.Second * 1)
}
