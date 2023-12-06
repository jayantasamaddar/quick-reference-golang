package codewars

/* GetCount */
// Return the number (count) of vowels in the given string.
// We will consider a, e, i, o, u as vowels for this Kata (but not y).
// The input string will only consist of lower case letters and/or spaces.
func GetCount(str string) (count int) {
	for _, s := range str {
		switch s {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
