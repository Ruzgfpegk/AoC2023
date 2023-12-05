package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("4_card.txt") // 15455663
	//file, err := os.Open("4_example2.txt") // 30

	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()

	sumOfCards := 0

	numberOfCardsPerSeries := make([]int, 2) // Index 0 inutilisé, index 1 pour la carte 1

	reCardNumberAndWinningsAndDrawn := regexp.MustCompile(`^Card\s+(\d+):\s+([\d\s]+?)\s+\|\s+([\d\s]+)$`)
	reSplitBySpaces := regexp.MustCompile(`\s+`)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := reCardNumberAndWinningsAndDrawn.FindStringSubmatch(line)

		if matches != nil {
			//fmt.Println("Card", matches[1], ", winning", matches[2], ", drawn", matches[3])

			cardNumberInt, _ := strconv.Atoi(matches[1])
			numberOfCardsPerSeries[cardNumberInt]++

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

			// Préparation de l'index suivant, quoi qu'il arrive
			if cardNumberInt < len(numberOfCardsPerSeries) {
				numberOfCardsPerSeries = append(numberOfCardsPerSeries, 0)
			}

			// Ajout du nombre de cartes suivantes
			for i := 1; i <= localWins; i++ {
				// Préparation des index des autres cartes à suivre
				if cardNumberInt+i < len(numberOfCardsPerSeries) {
					numberOfCardsPerSeries = append(numberOfCardsPerSeries, 0)
				}

				numberOfCardsPerSeries[cardNumberInt+i] += numberOfCardsPerSeries[cardNumberInt]
			}

			// Ajout des cartes de la série en cours à la somme
			sumOfCards += numberOfCardsPerSeries[cardNumberInt]
		}
	}

	fmt.Println("Total :", sumOfCards)
}
