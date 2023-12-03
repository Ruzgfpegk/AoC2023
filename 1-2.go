package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/1
// Avec les substitutions (partie 2) : 57345

func main() {
	file, err := os.Open("1_calibration.txt")
	//file, err := os.Open("1_example2.txt") // 281, avec substitutions
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier: ", err)
		return
	}
	defer file.Close()

	// Faits de telle sorte que le "remplacement" d'un mot n'en coupe par un autre
	numsLong := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numsShort := [9]string{"1e", "2o", "3e", "4", "5e", "6", "7n", "8t", "9e"}

	reFinalMatch := regexp.MustCompile(`(\d).*?(\d)?[^0-9]*$`)

	scanner := bufio.NewScanner(file)

	var total = 0
	var lineNum = 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Remplacement de texte
		for i := 0; i < len(line); i++ { // Itération dans la ligne
			for j := 0; j < 9; j++ { // Itération dans les nombres
				if strings.Index(line, numsLong[j]) == i {
					line = strings.ReplaceAll(line, numsLong[j], numsShort[j]) // Injection du nombre raccourci
				}
			}
		}

		matches := reFinalMatch.FindStringSubmatch(line)

		if matches != nil {
			// S'il n'y a qu'un seul nombre, il est à la fois le premier et le dernier
			if matches[2] == "" {
				matches[2] = matches[1]
			}

			lineFinal, _ := strconv.Atoi(matches[1] + matches[2])
			total += lineFinal
			//fmt.Println(lineNum, " - ", total, " - ", lineFinal)
		} else {
			fmt.Println("Erreur ligne ", lineNum)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier: ", err)
		return
	}

	fmt.Println("Total: ", total)
}
