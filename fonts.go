package main

import "unicode"

func replaceRunes(in_raw string, m map[rune]rune) string {
	in_runes := []rune(in_raw)
	var out_runes []rune

	for _, key := range in_runes {
		value := m[key]
		if value == rune(0) {
			out_runes = append(out_runes, key)
		} else {
			out_runes = append(out_runes, value)
		}
	}

	return string(out_runes)
}

func smallText(raw string) string {
	m := map[rune]rune{
		'b': 'ʙ',
		'd': 'ᴅ',
		'f': 'ғ',
		'g': 'ɢ',
		'h': 'ʜ',
		'i': 'ı',
		'j': 'ᴊ',
		'k': 'ᴋ',
		'l': 'ʟ',
		'p': 'ᴩ',
		'q': 'ϙ',
		't': 'ᴛ',
		'y': 'ʏ',
	}

	raw_runes := []rune(raw)
	for i, r := range raw_runes {
		raw_runes[i] = unicode.ToLower(r)
	}
	raw = string(raw_runes)

	return replaceRunes(raw, m)
}
