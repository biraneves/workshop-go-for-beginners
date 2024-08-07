package main

import (
	"flag"
	"fmt"
)

func main() {
	userCsvFile := flag.String("file", "testdata/users_001.csv", "Path to user file")
	nWinners := flag.Int("winners", 3, "Number of winners")

	flag.Parse()

	records, err := ParseCsvFile(*userCsvFile)
	if err != nil {
		return
	}

	selectedWinners := SelectWinners(records, *nWinners)

	for i, winner := range selectedWinners {
		fmt.Printf("%d: %s (%s)\n", i+1, winner[1], winner[0])
	}
}
