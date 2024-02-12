package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func chooseFont(fontChoice int) func(byte) []string {
	switch fontChoice {
	case 1:
		return standard
	case 2:
		return shadow
	case 3:
		return thinkertoy
	default:
		return standard
	}
}

func generateAsciiArt(text string, fontChoice int) string {
	var m1 []string

	for i := 0; i < len(text); i++ {
		var m2 []string
		switch fontChoice {
		case 1:
			m2 = standard(text[i])
		case 2:
			m2 = shadow(text[i])
		case 3:
			m2 = thinkertoy(text[i])
		}

		if i == 0 {
			m1 = m2
		} else {
			for j := 0; j < len(m1); j++ {
				m1[j] += m2[j]
			}
		}
	}
	return strings.Join(m1, "\n")
}

func processFlags() (string, string, int) {
	var inputText string
	var outputFileName string
	var fontChoice int

	flag.StringVar(&inputText, "text", "", "Input text for ASCII art")
	flag.StringVar(&outputFileName, "output", "", "Output file for ASCII art")
	flag.IntVar(&fontChoice, "font", 1, "Choose a font style (1: Standard, 2: Shadow, 3: Thinkertoy)")

	flag.Parse()

	if inputText == "" || outputFileName == "" {
		fmt.Println("Usage: go run . -text=\"{YourText}\" -output=banner.txt -font=1")
		os.Exit(1)
	}

	return inputText, outputFileName, fontChoice
}

func processFile(inputText, outputFileName string, fontChoice int) error {
	modifiedText := generateAsciiArt(inputText, fontChoice)

	err := os.WriteFile(outputFileName, []byte(modifiedText), 0644)
	if err != nil {
		return fmt.Errorf("error writing to output file: %v", err)
	}

	return nil
}

func main() {
	inputText, outputFileName, fontChoice := processFlags()

	err := processFile(inputText, outputFileName, fontChoice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Check the result file:", outputFileName)
}

// go run . --output=banner.txt --text="hello" --font=3