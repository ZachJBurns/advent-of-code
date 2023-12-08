package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	file, _ := os.ReadFile("day7.input")
	text := string(file)
	split_text := strings.Split(text, "\n")
	hands_and_bet := map[string]int{}
	hands := []string{}

	for _, line := range split_text {
		if line == "" {
			continue
		}
		split_play := strings.Split(line, " ")
		hand, bet_str := split_play[0], split_play[1]
		bet, _ := strconv.Atoi(bet_str)
		hands_and_bet[hand] = bet
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(a, b int) bool {
		hand_a := hands[a]
		hand_b := hands[b]

		hand_a_score, hand_b_score := score(hand_a), score(hand_b)
		if hand_a_score > hand_b_score {
			return true
		} else if hand_a_score < hand_b_score {
			return false
		} else {
			for i := 0; i < len(hand_a); i++ {
				if hand_a[i] != hand_b[i] {
					if hand_a[i] == 'A' {
						return true
					} else if hand_b[i] == 'A' {
						return false
					} else if hand_a[i] == 'K' {
						return true
					} else if hand_b[i] == 'K' {
						return false
					} else if hand_a[i] == 'Q' {
						return true
					} else if hand_b[i] == 'Q' {
						return false
					} else if hand_a[i] == 'J' {
						return true
					} else if hand_b[i] == 'J' {
						return false
					} else if hand_a[i] == 'T' {
						return true
					} else if hand_b[i] == 'T' {
						return false
					} else {
						return hand_a[i] > hand_b[i]
					}

				}
			}
			// never hits this case
			return true
		}
	})

	total := 0
	for idx, hand := range hands {
		total += hands_and_bet[hand] * (len(hands) - idx)
	}
	fmt.Println(total)

}

func score(hand string) int {
	hand_count := map[string]int{}
	score_map := [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 2}, {1, 2, 2}, {1, 1, 3}, {2, 3}, {1, 4}, {5}}

	for i := 0; i < len(hand); i++ {
		hand_count[string(hand[i])] += 1
	}

	hand_arr := []int{}
	for _, v := range hand_count {
		hand_arr = append(hand_arr, v)
	}

	sort.Ints(hand_arr)

	hand_score := 0
	for idx, score := range score_map {
		if reflect.DeepEqual(hand_arr, score) {
			hand_score = idx + 1
		}
	}
	return hand_score
}

func part2() {
	file, _ := os.ReadFile("day7.input")
	text := string(file)
	split_text := strings.Split(text, "\n")
	hands_and_bet := map[string]int{}
	hands := []string{}

	for _, line := range split_text {
		if line == "" {
			continue
		}
		split_play := strings.Split(line, " ")
		hand, bet_str := split_play[0], split_play[1]
		bet, _ := strconv.Atoi(bet_str)
		hands_and_bet[hand] = bet
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(a, b int) bool {
		hand_a := hands[a]
		hand_b := hands[b]

		replace_order := []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
		max_hand_a_score := -1
		max_hand_b_score := -1

		// Lets just brute force this...
		for _, replace := range replace_order {
			a := strings.ReplaceAll(hand_a, "J", replace)
			new_a_score := score(a)
			if new_a_score > max_hand_a_score {
				max_hand_a_score = new_a_score
			}

			b := strings.ReplaceAll(hand_b, "J", replace)
			new_b_score := score(b)
			if new_b_score > max_hand_b_score {
				max_hand_b_score = new_b_score
			}
		}

		if max_hand_a_score > max_hand_b_score {
			return true
		} else if max_hand_a_score < max_hand_b_score {
			return false
		} else {
			for i := 0; i < len(hand_a); i++ {
				if hand_a[i] != hand_b[i] {
					if hand_a[i] == 'J' {
						return false
					} else if hand_b[i] == 'J' {
						return true
					}
					if hand_a[i] == 'A' {
						return true
					} else if hand_b[i] == 'A' {
						return false
					} else if hand_a[i] == 'K' {
						return true
					} else if hand_b[i] == 'K' {
						return false
					} else if hand_a[i] == 'Q' {
						return true
					} else if hand_b[i] == 'Q' {
						return false
					} else if hand_a[i] == 'J' {
						return true
					} else if hand_b[i] == 'J' {
						return false
					} else if hand_a[i] == 'T' {
						return true
					} else if hand_b[i] == 'T' {
						return false
					} else {
						return hand_a[i] > hand_b[i]
					}

				}
			}
			return true
		}
	})

	total := 0
	for idx, hand := range hands {
		total += hands_and_bet[hand] * (len(hands) - idx)
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
