package main

import "log"

// This function is called when you register your Battlesnake on play.battlesnake.com
// It controls your Battlesnake appearance and author permissions.
func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "kavereon",
		Color:      "#8888888",
		Head:       "default",
		Tail:       "default",
	}
}
