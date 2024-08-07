package main

import "math/rand"

func SelectWinners(records [][]string, nWinners int) [][]string {
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	return records[:nWinners]
}
