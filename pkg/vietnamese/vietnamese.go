package vietnamese

import (
	"strings"
	"unicode/utf8"
)

var sourceCharacters = stringToRune(
	`ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚÝàáâãèéêìíòóôõùúý
ĂăĐđĨĩŨũƠơƯưẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọỎỏỐốỒồỔổỖỗỘộỚớỜờỞởỠỡỢợỤụỦủỨứỪừỬửỮữỰự`)

var destinationCharacters = stringToRune(
	`AAAAEEEIIOOOOUUYaaaaeeeiioooouuy
AaDdIiUuOoUuAaAaAaAaAaAaAaAaAaAaAaAaEeEeEeEeEeEeEeEeIiIiOoOoOoOoOoOoOoOoOoOoOoOoUuUuUuUuUuUuUu`)

func stringToRune(s string) []rune {
	result := make([]rune, 0, utf8.RuneCountInString(s))
	for _, c := range s {
		result = append(result, c)
	}
	return result
}

func binarySearch(sortedArray []rune, key rune) int {
	first := 0
	last := len(sortedArray) - 1
	for first <= last {
		mid := (first + last) / 2
		val := sortedArray[mid]
		if val == key {
			return mid
		} else if val > key {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return -1
}

func removeAccentChar(ch rune) rune {
	index := binarySearch(sourceCharacters, ch)
	if index >= 0 {
		return destinationCharacters[index]
	}
	return ch
}

// RemoveAccent removes Vietnamese accents
func RemoveAccent(s string) string {
	var b strings.Builder

	for _, runeValue := range s {
		_, _ = b.WriteRune(removeAccentChar(runeValue))
	}

	return b.String()
}
