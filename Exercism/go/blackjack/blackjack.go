package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10

	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Val := ParseCard(card1)
	card2Val := ParseCard(card2)
	dealerCardVal := ParseCard(dealerCard)
	mySum := card1Val + card2Val

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case mySum == 21:
		if dealerCardVal != 11 && dealerCardVal != 10 {
			return "W"
		}
		return "S"
	case mySum >= 17 && mySum <= 20:
		return "S"
	case mySum >= 12 && mySum <= 16:
		if dealerCardVal < 7 {
			return "S"
		}
		return "H"
	default:
		return "H"
	}
}
