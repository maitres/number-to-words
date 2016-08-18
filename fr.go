package ntn

import (
	"fmt"
	"log"
	"strings"
)

var frenchMegas = []string{"", "mille", "million", "milliard", "billiard", "trillion", "trilliard", "quadrillion", "quadrilliard", "quintillion", "quintilliard"}
var frenchUnits = []string{"", "un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf"}
var frenchTens = []string{"", "dix", "vingt", "trente", "quarante", "cinquante", "soixante", "soixante", "quatre-ving", "quatre-vingt"}
var frenchTeens = []string{"dix", "onze", "douze", "treize", "quatorze", "quinze", "seize", "dix-sept", "dix-huit", "dix-neuf"}

// IntegerToFrench converts an integer to French words
func IntegerToFrench(input int) string {
	log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "moins")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)
	log.Printf("Triplets: %v\n", triplets)

	// zero is a special case
	if len(triplets) == 0 {
		return "zéro"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		log.Printf("Triplet: %d (idx=%d)\n", triplet, idx)

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		log.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
		if hundreds > 0 {
			if hundreds == 1 {
				words = append(words, "cent")
			} else {
				words = append(words, frenchUnits[hundreds], "cents")
			}
		}

		switch tens {
		case 0:
			words = append(words, frenchUnits[units])
		case 1:
			words = append(words, frenchTeens[units])
			break
		case 7, 9:
			if units == 1 {
				words = append(words, frenchTens[tens], "et", frenchTeens[units])
			} else {
				word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchTeens[units])
				words = append(words, word)
			}
			break
		case 8:
			words = append(words, frenchTens[tens], frenchUnits[units])
			break
		default:
			if units == 1 {
				words = append(words, frenchTens[tens], "et", frenchUnits[units])
			} else {
				word := fmt.Sprintf("%s-%s", frenchTens[tens], frenchUnits[units])
				words = append(words, word)
			}
			break
		}

		// mega
		mega := frenchMegas[idx]
		if mega != "" && triplet > 1 {
			mega += "s"
		}
		if mega != "" {
			words = append(words, mega)
		}
	}

	log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}
