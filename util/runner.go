package util

import (
	"fmt"
	"log"
	"os"
)

func RunDailyProblemSet(
	part1 func(string),
	part2 func(string),
) {
	if len(os.Args) != 3 {
		fmt.Println("invalid arguments for problem set")
		return
	}

	part := os.Args[1]
	inputFilePath := os.Args[2]
	fmt.Println(fmt.Sprintf("Running Part %s, input file: %s", part, inputFilePath))

	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	stringData := string(data)

	switch part {
	case "1":
		part1(stringData)
	case "2":
		part2(stringData)
	default:
		log.Fatal("unknown problem part")
	}
}
