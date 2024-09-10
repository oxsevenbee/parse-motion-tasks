package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Read from stdin and skip first to lines. Then per line print the full line to stdout.

	// Create a new scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Skip the first two lines
	scanner.Scan()
	scanner.Scan()

	linkRe := regexp.MustCompile(`^\[([^|]*)].*`)
	//linkRe := regexp.MustCompile(`.*`)

	// Print the rest of the lines
	for scanner.Scan() {
		line := scanner.Text()

		// split line into words
		words := strings.Split(line, "|")
		match := linkRe.FindStringSubmatch(strings.TrimSpace(words[1]))

		duration := parseDuration(words[2])

		fmt.Printf("%s %s\n", match[1], duration)
	}
}

func parseDuration(duration string) string {
	// Parse the duration into minutes. Format: "1h", "30m", "1h 30m"
	minutes := 0
	words := strings.Split(duration, " ")
	for _, word := range words {
		if strings.HasSuffix(word, "h") {
			m, _ := strconv.Atoi(word[:len(word)-1])
			minutes += 60 * m
		} else if strings.HasSuffix(word, "m") {
			m, _ := strconv.Atoi(word[:len(word)-1])
			minutes += m
		}
	}
	return strconv.Itoa(minutes)
}
