package std

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*************************************************************************************************/
// String Operations
/*************************************************************************************************/
func StringOperationsDemo() {
	// Strings to work with
	str := "The quick brown fox jumps over the lazy dog."
	strSlice := []string{
		"The quick brown fox jumps over the lazy dog.",
		"Pack my box with five dozen liquor jugs.",
		"Mr. Jock, TV quiz PhD, bags few lynx.",
		"Cwm fjord bank glyphs vext quiz.",
		"Blowzy night-frumps vex'd Jack Q.",
		"Waltz, nymph, for quick jigs vex Bud.",
		"Big fjords vex quick waltz nymph.",
		"Sphinx of black quartz, judge my vow.",
		"Jump frog, vex bad luck, quip why?",
		"Vexed nymphs go for quick jigs, Waltz.",
		"Quick wafting zephyrs vex bold Jim.",
	}

	/********************************************************************************************************************/
	// Join: Joining an array of strings into a single string with a separator
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.Join" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Yellow + strings.Join(strSlice, "\n") + Reset)

	/********************************************************************************************************************/
	// Split: Slices string into substrings separated by seperator. Returns a slice of the substrings.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.Split" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	splitStr := strings.Split(str, " ")
	fmt.Printf(Green+"SPLIT 'str':"+Reset+" %v%s%v of length: %v%d%v and type %v%T%v\n", Yellow, splitStr, Reset, Yellow, len(splitStr), Reset, Yellow, splitStr, Reset)

	/********************************************************************************************************************/
	// Replace: Takes a string, replaces the old with the new string, n number of times, left to right.
	// If n < 0, there is no limit on the number of replacements. Returns a copy of the original string.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.Replace" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	// Replace only one instance of "e"
	fmt.Println(Green+"REPLACE ONLY A SINGLE INSTANCE:"+Reset, Yellow+strings.Replace(str, "e", "_REPLACED_", 1)+Reset)
	// Replace all instances of "e". n < 0 (e.g. -1)
	fmt.Println(Green+"REPLACE ALL INSTANCES:"+Reset, Yellow+strings.Replace(str, "e", "_REPLACED_", -1)+Reset)
	// Proof that the original string is unmodified
	fmt.Println(Green+"UNMODIFIED ORIGINAL STRING:"+Reset, Yellow+str+Reset)

	/********************************************************************************************************************/
	// ReplaceAll: Takes a string, replaces all occurences of the old with the new string.
	// Same as using strings. Replace with n < 0
	// Returns a copy of the original string.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.ReplaceAll" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	// Replace all occurences of a space with no space
	fmt.Println(Green+"REPLACE ALL INSTANCES:"+Reset, Yellow+strings.ReplaceAll(str, " ", "")+Reset)

	/********************************************************************************************************************/
	// ToUpper: Takes a string, returns with all Unicode letters mapped to their upper case.
	// ToLower: Takes a string, returns with all Unicode letters mapped to their lower case.
	// ToTitle: Takes a string, returns with all Unicode letters mapped to their title case.
	// Doesn't mutate the original string.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.ToUpper, strings.ToLower, strings.Title (deprecated)" + Reset)
	fmt.Println(Purple + "\tDemo: Using external packages: 'golang.org/x/text/cases', 'golang.org/x/text/languages'" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Green+"UPPERCASE:"+Reset, Yellow+strings.ToUpper(str)+Reset)
	fmt.Println(Green+"LOWERCASE:"+Reset, Yellow+strings.ToLower(str)+Reset)
	fmt.Println(Green+"TITLECASE (deprecated):"+Reset, Yellow+strings.Title(str)+Reset)
	titleCaser := cases.Title(language.English)
	fmt.Println(Green+"TITLECASE (using 'golang.org/x/text/cases'):"+Reset, Yellow+titleCaser.String(str)+Reset)
	fmt.Println(Green+"UNMODIFIED ORIGINAL STRING:"+Reset, Yellow+str+Reset)

	/********************************************************************************************************************/
	// Index: Returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	// LastIndex: Returns the index of the last instance of substr in s, or -1 if substr is not present in s.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.Index, strings.LastIndex" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	// Index of existing string returns the index
	fmt.Printf("%vINDEX OF 'quick':%v %v%d%v\n", Green, Reset, Yellow, strings.Index(str, "quick"), Reset)  // 4
	fmt.Printf("%vLAST INDEX OF 'e':%v %v%d%v\n", Green, Reset, Yellow, strings.LastIndex(str, "e"), Reset) // 33
	// Index of non-existent substring returns -1
	fmt.Printf("%vINDEX OF NON-EXISTENT 'apple':%v %v%d%v\n", Green, Reset, Yellow, strings.Index(str, "apple"), Reset)          // -1
	fmt.Printf("%vLAST INDEX OF NON-EXISTENT 'apple':%v %v%d%v\n", Green, Reset, Yellow, strings.LastIndex(str, "apple"), Reset) // -1

	/********************************************************************************************************************/
	// Slicing a String: Slicing a String returns a new string with the substring
	// Accessing a character by index: Converted into bytes which need to be typecasted back into string to print
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Slicing Strings, Accessing a character by Index" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	// Slicing a substring non-destructively
	the := str[:3]
	dog := string(str[len(str)-4 : len(str)-1])

	// Accessing a character. Gets converted into bytes, which need to be typecasted back to string
	q := string(str[4])
	g := string(str[len(str)-2])

	fmt.Println(Green+"SLICE THE FIRST 3 CHARS:"+Reset, Yellow+the+Reset) // The
	fmt.Println(Green+"THE FOURTH CHARACTER:"+Reset, Yellow+q+Reset)      // q
	fmt.Println(Green+"THE SECOND LAST CHARACTER:"+Reset, Yellow+g+Reset) // g
	fmt.Println(Green+"EXTRACT THE WORD 'dog':"+Reset, Yellow+dog+Reset)  // dog
	fmt.Println(Green+"UNMODIFIED ORIGINAL STRING:"+Reset, Yellow+str+Reset)

	/********************************************************************************************************************/
	// Contains: Returns a boolean whether substr is within s.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.Contains" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	fmt.Printf(Green+"CONTAINS SUBSTRING 'quick' (TRUE):"+Reset+" %v%t%v\n", Yellow, strings.Contains(str, "quick"), Reset) // true
	fmt.Printf(Green+"CONTAINS SUBSTRING 'dino' (FALSE):"+Reset+" %v%t%v\n", Yellow, strings.Contains(str, "dino"), Reset)  // false

	/********************************************************************************************************************/
	// HasPrefix: Returns a boolean whether the string s begins with prefix.
	// HasSuffix: Returns a boolean whether the string s ends with suffix.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.HasPrefix, strings.HasSuffix" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	// Case: true
	fmt.Printf(Green+"HAS PREFIX 'The' (TRUE):"+Reset+" %v%t%v\n", Yellow, strings.HasPrefix(str, "The"), Reset)   // true
	fmt.Printf(Green+"HAS SUFFIX 'dog.' (TRUE):"+Reset+" %v%t%v\n", Yellow, strings.HasSuffix(str, "dog."), Reset) // true

	// Case: false
	fmt.Printf(Green+"HAS PREFIX 'Dino' (FALSE):"+Reset+" %v%t%v\n", Yellow, strings.HasPrefix(str, "Dino"), Reset)   // false
	fmt.Printf(Green+"HAS SUFFIX 'dino.' (FALSE):"+Reset+" %v%t%v\n", Yellow, strings.HasSuffix(str, "dino."), Reset) // false
	// fmt.Println(PrintC(Green, "HAS SUFFIX 'dino.' (FALSE):"), PrintC(Yellow, fmt.Sprintf("%t", strings.HasSuffix(str, "dino.")))) // Prints the same

	/********************************************************************************************************************/
	// Trim: Returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.Trim" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Green+"TRIM ('T' and '.'):"+Reset, Yellow+strings.Trim(str, "T.")+Reset)
	fmt.Println(Green+"TRIM ('quick') DOESN'T WORK AS NOT LEADING OR TRAILING:"+Reset, Yellow+strings.Trim(str, "quick")+Reset)

	/********************************************************************************************************************/
	// Map: Returns a copy of the string s with all its characters modified according to the mapping function.
	// If mapping returns a negative value, the character is dropped from the string with no replacement.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strings.Map" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	// Mapping Function to Flip cases
	// Alphabets range: A-Z: 65-90, a-z: 97-122. There's a difference of ASCII 32 characters in between.
	// Reference: https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html
	// Note: We are dealing with runes, therefore single quotes must be used. Runes are true `int32` type.
	flipCaseMappingFunc := func(r rune) rune {
		switch {
		// If LowerCase, convert to UpperCase
		case r >= 97 && r <= 122:
			return r - 32

		// If UpperCase, convert to LowerCase (exception is first character)
		case r >= 65 && r <= 90:
			return r + 32
		}
		return r
	}
	fmt.Println(Green+"MAP TO FLIP CASES:"+Reset, Yellow+strings.Map(flipCaseMappingFunc, str)+Reset)
}
