package days

import (
	"os"
	"fmt"
	"strings"
	"math"
)

func ScoreCard(c string) int {
	card := strings.Split(strings.Split(c, ": ")[1], " | ")
	winning := strings.Split(card[0], " ")
	actual := strings.Split(card[1], " ")
	matches := 0
	for _,w := range winning {
		for _,a := range actual {
			if w == a {
				matches += 1
			}
		}
	}

	return matches
}



func Day04() bool {
	b,_ := os.ReadFile("input/day04.txt")
	input := string(b)
	input = strings.Replace(input,"  ", " ", -1)
	input_list := strings.Split(input[:len(input)-1], "\n") 

	max_cards := len(input_list)
	scores := make([]int,max_cards)
	copies := make([]int,max_cards)

	p1 := 0
	p2 := 0

	for i,card := range input_list {
		scores[i] =  ScoreCard(card)
		if scores[i] > 0 {p1 += int(math.Pow(2, float64(scores[i]-1)))}

		copies[i] += 1
		for j := i+1; j <= i+scores[i] && j < max_cards; j++ {
			copies[j] += copies[i]
		}
		p2 += copies[i]
	}
	
	fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
	return true
}
