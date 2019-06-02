package anagrams

import (
	"fmt"
	"strings"
)

func anagram(toAnagram string) []string {
	res := []string{}
	if len(toAnagram) == 2 {
		res = append(res, fmt.Sprintf("%s%s", string(toAnagram[1]), string(toAnagram[0])))
		res = append(res, toAnagram)
		return res
	}
	for _, value := range toAnagram {
		characterRemoved := removeCharacters(toAnagram, string(value))
		anagramWithoutChar := anagram(characterRemoved)
		for _, v := range anagramWithoutChar {
			res = append(res, fmt.Sprintf("%s%s", string(value), v))
		}
	}
	return res
}

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}
	return strings.Map(filter, input)
}
