package std

import (
	"fmt"
	"strings"
)

// Terminal Colours
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func PrintC(color string, strs ...string) string {
	return color + strings.Join(strs, " ") + Reset
}

func PrintHeader(s string) {
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\t" + s + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
}
