package lib

import (
	"fmt"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

// Add an entry to the Journal
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// Remove an entry from the Journal
func (j *Journal) RemoveEntry(index int) bool {
	if index >= j.Count() {
		return false
	} else if index == j.Count()-1 {
		// Remove last index
		j.entries = append(j.entries[:index])
	} else {
		j.entries = append(j.entries[:index], j.entries[:index+1]...)
	}
	entryCount = j.Count()
	return true
}

// Display the count
func (j *Journal) Count() int {
	return len(j.entries)
}

// Stringify
func (j *Journal) String(sep string) string {
	return strings.Join(j.entries, sep)
}
