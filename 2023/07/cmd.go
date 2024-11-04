package adventOfCode_2023_07

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "07",
	Short: "Day 07 of Advent of Code 2023",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent string, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`%s/%s/input.txt`, parent, command))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 01: %d\n", part1(string(b)))
	fmt.Printf("Part 02: %d\n", part2(string(b)))
}

type HandBid struct {
	hand string
	bid  int
}

func part1(input string) int64 {
	list := ParseFile(input)

	sort.Slice(list, func(i, j int) bool {
		return HandBidSortFunc(list[i], list[j], false)
	})

	var result int64 = 0
	for i, handBid := range list {
		result += int64(handBid.bid * (i + 1))
	}

	return result
}

func part2(input string) int64 {
	list := ParseFile(input)

	sort.Slice(list, func(i, j int) bool {
		return HandBidSortFunc(list[i], list[j], true)
	})

	var result int64 = 0
	for i, handBid := range list {
		result += int64(handBid.bid * (i + 1))
	}

	return result
}

func ParseFile(input string) []HandBid {
	lines := strings.Split(input, "\n") // Each line represents a hand and its bid

	list := make([]HandBid, 0)

	for _, line := range lines {
		if regexp.MustCompile(`^$`).MatchString(line) {
			continue
		}
		var hand string
		var bid int
		if _, err := fmt.Sscanf(line, "%s %d", &hand, &bid); err != nil {
			panic(err)
		}
		list = append(list, HandBid{hand, bid})
	}

	return list
}

// CARD_VALUE uses the index of the card to determine its strength.
// 'T' is at index 10, so it's stronger than '4' at index 3 
var CARD_VALUE = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func GetCardValue(card rune) int {
	return slices.Index(CARD_VALUE, card)
}

// HandBidSortFunc is a function used to sort a slice of HandBid by hand
// strength. If both hands have the same type (eg: one pair), then their
// strength is calculated using the first card, in each hand, that has the
// highest value (eg: AAJ23 has more value than AJA23)
func HandBidSortFunc(h1, h2 HandBid, handleJoker bool) bool {
	var s1, s2 int
	if handleJoker {
		s1 = GetHandStrengthWithJokers(h1.hand)
		s2 = GetHandStrengthWithJokers(h2.hand)
	} else {
		s1 = GetHandStrengthWithoutJokers(h1.hand)
		s2 = GetHandStrengthWithoutJokers(h2.hand)
	}

	switch {
	case s1 < s2:
		return true
	case s1 > s2:
		return false
	}

	for i := 0; i < 5; i++ {
		diff := GetCardValue(rune(h1.hand[i])) - GetCardValue(rune(h2.hand[i]))
		if diff != 0 {
			return diff < 1
		}
	}

	return false
}

// Hand strengths
const (
	HIGH_CARD       = 0
	ONE_PAIR        = 1
	TWO_PAIRS       = 2
	THREE_OF_A_KIND = 3
	FULL_HOUSE      = 4
	FOUR_OF_A_KIND  = 5
	FIVE_OF_A_KIND  = 6
)

// GetHandStrengthWithoutJokers calculates the strength of the provided poker
// hand without processing jokers as such.
//
// eg: AAJ23 is considered a pair
func GetHandStrengthWithoutJokers(hand string) int {
	charOccurences := map[rune]int{}

	for _, c := range hand {
		charOccurences[c]++
	}

	_, highestOccurence := GetMostOccurringCard(charOccurences)

	switch {
	case len(charOccurences) == 5:
		return HIGH_CARD
	case len(charOccurences) == 4:
		return ONE_PAIR
	case len(charOccurences) == 3 && highestOccurence == 2:
		return TWO_PAIRS
	case len(charOccurences) == 3 && highestOccurence == 3:
		return THREE_OF_A_KIND
	case len(charOccurences) == 2 && highestOccurence == 3:
		return FULL_HOUSE
	case len(charOccurences) == 2 && highestOccurence == 4:
		return FOUR_OF_A_KIND
	case len(charOccurences) == 1:
		return FIVE_OF_A_KIND
	default:
		return 0
	}
}

// GetHandStrengthWithJokers calculates the strength of the provided poker hand
// with jokers processed as a placeholder to create the strongest hand.
//
// Eg: AAJ23 is considered a three of a kind.
func GetHandStrengthWithJokers(hand string) int {
	charOccurences := map[rune]int{}

	for _, c := range hand {
		charOccurences[c]++
	}

	uniqueCards := len(charOccurences)
	jokerAmount, ok := charOccurences['J']
	if ok {
		uniqueCards -= 1
	}
  delete(charOccurences, 'J')

	_, highestOccurence := GetMostOccurringCard(charOccurences)

	switch {
	case uniqueCards == 5:
		return HIGH_CARD
	case uniqueCards == 4:
		return ONE_PAIR
	case uniqueCards == 3 && highestOccurence+jokerAmount == 2:
		return TWO_PAIRS
	case uniqueCards == 3 && highestOccurence+jokerAmount == 3:
		return THREE_OF_A_KIND
	case uniqueCards == 2 && highestOccurence+jokerAmount == 3:
		return FULL_HOUSE
	case uniqueCards == 2 && highestOccurence+jokerAmount == 4:
		return FOUR_OF_A_KIND
	case uniqueCards <= 1: // if 0, then the hand is 5 jokers
		return FIVE_OF_A_KIND
	default:
		return 0
	}
}

// GetMostOccurringCard returns the card that has the most occurrences in the
// hand, along with its amount of occurrences.
//
// eg: AAJ23 returns (A, 2)
func GetMostOccurringCard(cardOccurences map[rune]int) (card rune, occurrences int) {
	for k, v := range cardOccurences {
		if v > occurrences {
			occurrences = v
			card = k
		}
	}

	return
}
