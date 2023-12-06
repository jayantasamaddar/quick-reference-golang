package std

import (
	"fmt"
	"strings"
)

func SliceOperationsDemo() {
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
	}

	/********************************************************************************************************************/
	// Length: Find the length of the slice
	/********************************************************************************************************************/
	fmt.Printf("%s: %d\n", PrintC(Green, "LENGTH OF SLICE (`strSlice`)"), len(strSlice))

	/********************************************************************************************************************/
	// Append: Versatile function that can be used to:
	// (1) Push items to an existing slice
	// (2) Shift item to an existing slice
	// (3) Concatenate multiple slices into one
	// (4) Splice the slice, removing certain elements
	/********************************************************************************************************************/

	/********************************************************************************************************************/
	// (1) Push: Add an element to the end of the slice
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Push Operation" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	strSlice = append(strSlice, "The five boxing wizards jump quickly.")

	fmt.Println(PrintC(Yellow, strings.Join(strSlice, "\n")))
	fmt.Printf("\n%s: %d\n", PrintC(Green, "LENGTH OF SLICE (`strSlice`)"), len(strSlice))

	/********************************************************************************************************************/
	// (2) Shift: Add an element to the front of the slice
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Shift Operation" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	strSlice = append([]string{"Jackdaws love my big sphinx of quartz."}, strSlice...)

	fmt.Println(PrintC(Yellow, strings.Join(strSlice, "\n")))
	fmt.Printf("\n%s: %d\n", PrintC(Green, "LENGTH OF SLICE (`strSlice`)"), len(strSlice))

	/********************************************************************************************************************/
	// (3) Concatenate: Take one []string and append into another.
	// The result maybe assigned to either of the slices used for concatenation or an entirely new slice.
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Concatenate Operation" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	strSlice2 := []string{
		"Quick wafting zephyrs vex bold Jim.",
		"How vexingly quick daft zebras jump!",
		"Sphinx-like, they giggle at jumbled text.",
		"The jay, pig, fox, zebra, and my wolves quack!",
		"Brawny gods just flocked up to quiz and vex him.",
		"J.Q. Vandz struck my big fox whelp.",
		"Five or six big juicy steaks sizzled in a pan.",
		"Mix Zapf with Veljovic and get quirky Beziers.",
		"Sympathizing would fix Quaker objectives.",
	}

	// Concatenate to a new Slice
	concatenatedSlice := append(strSlice, strSlice2...)

	fmt.Println(PrintC(Cyan, strings.Join(concatenatedSlice, "\n")))
	fmt.Printf("\n%s: %d\n", PrintC(Green, "LENGTH OF SLICE (`concatenatedSlice`)"), len(concatenatedSlice))

	// Old slices are unchanged
	fmt.Println()
	fmt.Println(PrintC(Yellow, strings.Join(strSlice, "\n")))
	fmt.Printf("\n%s: %d\n", PrintC(Green, "LENGTH OF SLICE (`strSlice`)"), len(strSlice))
	fmt.Println()
	fmt.Println(PrintC(Blue, strings.Join(strSlice2, "\n")))
	fmt.Printf("\n%s: %d\n", PrintC(Green, "LENGTH OF SLICE (`strSlice2`)"), len(strSlice2))

	/********************************************************************************************************************/
	// (4) Splice: Remove elements from a slice
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: Splice Operation" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	// Case 1: From `strSlice2` remove the 4th (index: 3) and 5th (index: 4) element
	fmt.Printf("%s: %d\n", PrintC(Green, "LENGTH OF SLICE (`strSlice2`) BEFORE"), len(strSlice2))

	strSlice2 = append(strSlice2[:3], strSlice2[5:]...)

	fmt.Printf("%s: %d\n", PrintC(Green, "LENGTH OF SLICE (`strSlice2`) AFTER"), len(strSlice2))
	fmt.Println()
	fmt.Println(PrintC(Blue, strings.Join(strSlice2, "\n")))

}
