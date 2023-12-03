package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/2

func main() {
	file, err := os.Open("2_games.txt") // 63307
	//file, err := os.Open("2_example1.txt") // 2286

	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	sumOfPowers := 0

	reGameNumberAndData := regexp.MustCompile(`^Game (\d+): (.*)$`)
	reNumberAndColor := regexp.MustCompile(`^(\d+) (\w+)$`)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := reGameNumberAndData.FindStringSubmatch(line)

		if matches != nil {
			//gameNumber, _ := strconv.Atoi(matches[1])
			reachList := strings.Split(matches[2], "; ")

			localMax := map[string]int{"red": 0, "green": 0, "blue": 0}

			for _, reach := range reachList {
				colorsInReach := strings.Split(reach, ", ")

				for _, colorAndNumber := range colorsInReach {
					numberAndColorMatch := reNumberAndColor.FindStringSubmatch(colorAndNumber)

					if numberAndColorMatch != nil {
						number, _ := strconv.Atoi(numberAndColorMatch[1])

						if localMax[numberAndColorMatch[2]] < number {
							localMax[numberAndColorMatch[2]] = number
						}
					}
				}
			}

			//fmt.Println(gameNumber, "-", localMax)
			// Ajout de la puissance du jeu
			localPower := localMax["red"] * localMax["green"] * localMax["blue"]
			sumOfPowers += localPower
			//fmt.Println(gameNumber, "-", localPower, "=> Total", sumOfPowers)
		} // Après matching ligne
	} // Fin scan ligne par ligne
	fmt.Println("Total:", sumOfPowers)
}
