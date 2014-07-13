package main

import (
	"code.google.com/p/go.text/unicode/norm"
	"math/rand"
	"strings"
	"unicode"
)

type language uint8

const (
	hawaiian language = iota
	japanese
)

type name struct {
	origin language
	min    int // minimum length (syllables)
	max    int // maximum length (syllables)
}

func toTitleCase(s string) string {

	rs := []rune(s)

	i := 0
	for ; !unicode.IsLetter(rs[i]); i++ {
	}

	rs[i] = unicode.ToTitle(rs[i])

	return string(rs)
}

func (n name) String(r *rand.Rand) (s string) {

	syllableCount := n.min + r.Intn(n.max-n.min)

	switch n.origin {
	case hawaiian:
		vowels := []string{
			"a", "e", "i", "o", "u",
			"ā", "ē", "ī", "ō", "ū",
			"ai", "ae", "ao", "au", "ei",
			"eu", "iu", "oi", "ou", "ui",
		}
		consonants := []string{
			"'", "h", "k", "l",
			"m", "n", "p", "w",
		}

		for i := 0; i < syllableCount; i++ {
			if r.Intn(1) == 0 {
				i_consonant := r.Intn(len(consonants))
				s += consonants[i_consonant]
			}
			i_vowel := r.Intn(len(vowels))
			s += vowels[i_vowel]
		}

	case japanese:
		// order preserved from omniglot for maintainability
		// http://www.omniglot.com/writing/japanese_katakana.htm

		// standard syllabary
		standard := []string{
			"ka", "ki", "ku", "ke", "ko",
			"sa", "shi", "su", "se", "so",
			"ta", "chi", "tsu", "te", "to",
			"na", "ni", "nu", "ne", "no",
			"ha", "hi", "hu", "he", "ho",
			"ma", "mi", "mu", "me", "mo",
			"ya", "yu", "yo",
			"ra", "ri", "ru", "re", "ro",
			"wa", "wo",
		}

		// these come up a lot
		// (japan likes its vowels)
		bareVowels := []string{
			"a", "i", "u", "e", "o",
		}

		// syllabary plus dakuten
		additional := []string{
			"ga", "gi", "gu", "ge", "go",
			"za", "ji", "zu", "ze", "zo",
			"da", "ji" /*"zu",*/, "de", "do",
			"ba", "bi", "bu", "be", "bo",
			"pa", "pi", "pu", "pe", "po",
		}

		// extended syllabary
		extended := []string{
			"kya", "kyu", "kyo",
			"gya", "gyu", "gyo",
			"nya", "nyu", "nyo",
			"hya", "hyu", "hyo",
			"bya", "byu", "byo",
			"pya", "pyu", "pyo",
			"mya", "myu", "myo",
			"rya", "ryu", "ryo",
			"ja", "ju", "je", "jo",
			"cha", "chu", "che", "cho",
			"sha", "shu", "she", "sho",
		}

		for i := 0; i < syllableCount; i++ {
			where := r.Intn(9)
			switch {
			case where == 8:
				i_extended := r.Intn(len(extended))
				s += extended[i_extended]
			case where > 5:
				i_additional := r.Intn(len(additional))
				s += additional[i_additional]
			case where > 2:
				i_standard := r.Intn(len(standard))
				s += standard[i_standard]
			default:
				i_vowel := r.Intn(len(bareVowels))
				vowel := bareVowels[i_vowel]

				// If we're adding a bare vowel, we need to add a macron iff
				// the last vowel is the same
				if strings.HasSuffix(s, vowel) {

					// append a "combining macron above" to the string
					s += "\u0304"

					// NFC-normalize:
					// combines the macron with the previous character
					s = string(norm.NFC.Bytes([]byte(s)))

					// Otherwise, just add the vowel
				} else {
					s += bareVowels[i_vowel]
				}
			}

			if r.Intn(3) == 0 {
				s += "n"
			}
		}
	}

	return toTitleCase(s)
}
