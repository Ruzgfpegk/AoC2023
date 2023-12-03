package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2023/day/1
// Sans les substitutions (partie 1) : 57346

func main() {
	file, err := os.Open("1_calibration.txt")
	//file, err := os.Open("1_example1.txt") // 142, sans substitutions
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier: ", err)
		return
	}
	defer file.Close()

	reFinalMatch := regexp.MustCompile(`(\d).*?(\d)?[^0-9]*$`)

	scanner := bufio.NewScanner(file)

	var total = 0
	var lineNum = 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

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
