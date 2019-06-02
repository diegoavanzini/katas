package anagrams

import (
	"strconv"
	"strings"
	"testing"
)

func Test_bi_anagram_should_contain_ib(t *testing.T) {
	if !contains(anagram("bi"), "ib") {
		t.Error("anagram(\"bi\") doesn't contains \"ib\" as expected.")
	}
}

func Test_length_bi_anagram_should_be_2(t *testing.T) {
	if len(anagram("bi")) != 2 {
		t.Error("anagram(\"bi\") != 2 as expected.")
	}
}

func Test_bir_anagram_should_contain_ibr(t *testing.T) {
	anag := anagram("bir")
	if !contains(anag, "ibr") {
		t.Error("anagram(\"bir\") doesn't contains \"ibr\" as expected.")
	}
	if len(anag) != 6 {
		t.Error("anagram(\"bir\") doesn't contains 6 elements as expected.")
	}
}

func Test_biro_anagram_should_contain_ibro(t *testing.T) {
	anag := anagram("biro")
	if !contains(anag, "ibro") {
		t.Error("anagram(\"biro\") doesn't contains \"ibr\" as expected.")
	}
	cont := strings.Split("biro bior brio broi boir bori ibro ibor irbo irob iobr iorb rbio rboi ribo riob roib robi obir obri oibr oirb orbi orib", " ")
	for _, value := range cont {
		if !contains(anag, value) {
			t.Errorf("anagram(\"biro\") doesn't contains \"%s\" as expected.", value)
		}
	}
}

func Test_biro_anagram_should_contains_24_elements(t *testing.T) {
	anag := anagram("biro")
	if len(anag) != 24 {
		t.Error("anagram(\"bir\") doesn't contains 24 elements as expected." + strconv.Itoa(len(anag)))
	}
}

func Test_biro_anagram_should_contains_all_elements(t *testing.T) {
	anag := anagram("biro")
	cont := strings.Split("biro bior brio broi boir bori ibro ibor irbo irob iobr iorb rbio rboi ribo riob roib robi obir obri oibr oirb orbi orib", " ")
	for _, value := range cont {
		if !contains(anag, value) {
			t.Errorf("anagram(\"biro\") doesn't contains \"%s\" as expected.", value)
		}
	}
}


func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

