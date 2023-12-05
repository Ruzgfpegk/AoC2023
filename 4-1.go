package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("4_card.txt") // 23678
	//file, err := os.Open("4_example1.txt") // 13

	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	sumOfPoints := 0

	reCardNumberAndWinningsAndDrawn := regexp.MustCompile(`^Card\s+(\d+):\s+([\d\s]+?)\s+\|\s+([\d\s]+)$`)
	reSplitBySpaces := regexp.MustCompile(`\s+`)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := reCardNumberAndWinningsAndDrawn.FindStringSubmatch(line)

		if matches != nil {
			//fmt.Println("Card", matches[1], ", winning", matches[2], ", drawn", matches[3])

			var winning [100]bool
			localWins := 0

			winningList := reSplitBySpaces.Split(matches[2], -1)
			drawnList := reSplitBySpaces.Split(matches[3], -1)

			// Population du tableau de vérification
			for _, winningElement := range winningList {
				winningElementInt, _ := strconv.Atoi(winningElement)
				winning[winningElementInt] = true
			}

			// Test sur le tableau de vérification
			for _, drawnListElement := range drawnList {
				drawnListElementInt, _ := strconv.Atoi(drawnListElement)

				if winning[drawnListElementInt] {
					localWins++
				}
			}

			// Décompte des points de la carte
			if localWins > 0 {
				sumOfPoints += powInt(2, localWins-1)
			}
		}

		//fmt.Println(powInt(2, 0), powInt(2, 1), powInt(2, 2), powInt(2, 3)) // Test rapide fonction powInt
	}

	fmt.Println("Total :", sumOfPoints)
}

// Le math.Pow de Go balance du float. Autant rester en int tout du long.
// https://simple.wikipedia.org/wiki/Exponentiation_by_squaring
func powInt(base, exponent int) int {
	result := 1

	for exponent > 0 {
		if exponent%2 == 1 {
			result *= base
		}
		base *= base
		exponent /= 2
	}

	return result
}
