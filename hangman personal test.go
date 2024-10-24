package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func loadWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.ToLower(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

func displayWord(word string, revealed []bool) string {
	display := ""
	for i, letter := range word {
		if revealed[i] {
			display += string(letter)
		} else {
			display += "_"
		}
	}
	return display
}

func revealRandomLetters(word string, revealed []bool, n int) {
	rand.Seed(time.Now().UnixNano())
	lettersRevealed := 0

	for lettersRevealed < n {
		index := rand.Intn(len(word))
		if !revealed[index] {
			revealed[index] = true
			lettersRevealed++
		}
	}
}

func displayHangman(attempts int) {
	hangmanStages := []string{
		"=========\n\n      |  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========",
		"  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========",
		"  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========",
		"  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========",
		"  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========",
		"  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========",
		"  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========",
		"  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========",
		"  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========",
	}

	fmt.Println(hangmanStages[10-attempts])
}

func hangman(filename string) {
	words, err := loadWords(filename)
	if err != nil {
		fmt.Println("Erreur lors du chargement des mots:", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]

	revealed := make([]bool, len(word))
	n := len(word)/2 - 1
	revealRandomLetters(word, revealed, n)
	attempts := 10

	fmt.Println("Bienvenue dans le jeu Hangman!")
	fmt.Printf("Vous avez %d tentatives pour deviner le mot.\n", attempts)

	for attempts > 0 {
		displayHangman(attempts)
		fmt.Printf("\nMot à deviner: %s\n", displayWord(word, revealed))
		fmt.Printf("Tentatives restantes: %d\n", attempts)

		fmt.Print("Devinez une lettre: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		guess := strings.ToLower(scanner.Text())

		if len(guess) != 1 {
			fmt.Println("Veuillez entrer une seule lettre.")
			continue
		}

		letterFound := false
		for i, letter := range word {
			if string(letter) == guess && !revealed[i] {
				revealed[i] = true
				letterFound = true
			}
		}

		if letterFound {
			fmt.Printf("Bonne lettre: '%s'!\n", guess)
		} else {
			fmt.Printf("Mauvaise lettre: '%s'.\n", guess)
			attempts--
		}

		wordGuessed := true
		for _, revealedLetter := range revealed {
			if !revealedLetter {
				wordGuessed = false
				break
			}
		}

		if wordGuessed {
			fmt.Printf("Félicitations! Vous avez deviné le mot: %s\n", word)
			break
		}
	}

	if attempts == 0 {
		fmt.Printf("Vous avez perdu! Le mot était: %s\n", word)
	}
}

func main() {
	hangman("words.txt")
}
