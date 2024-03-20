package main

import (
	"fmt"
	"strings"
)

const (
	Markdown int = iota
	Html
)

type ListStrategy interface {
	Prefix(builder *strings.Builder)
	Suffix(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

// Markdown strategy
type MarkdownList struct{}

func (m *MarkdownList) Prefix(builder *strings.Builder) {}

func (m *MarkdownList) Suffix(builder *strings.Builder) {}

func (m *MarkdownList) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("- " + item + "\n")
}

// HTML Strategy
type HTMLList struct{}

func (h *HTMLList) Prefix(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HTMLList) Suffix(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HTMLList) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("  <li>" + item + "</li>\n")
}

// TextProcessor
type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

// Constructor
func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetStrategy(format int) {
	switch format {
	case Markdown:
		t.listStrategy = &MarkdownList{}
	case Html:
		t.listStrategy = &HTMLList{}
	}
}

// Add an item to the list
func (t *TextProcessor) AppendList(items ...string) {
	s := t.listStrategy
	s.Prefix(&t.builder)
	for _, item := range items {
		s.AddListItem(&t.builder, item)
	}
	s.Suffix(&t.builder)
}

// Reset the string builder
func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

// Default formatting option
func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(&MarkdownList{})
	tp.AppendList("Apple", "Bat", "Cat", "Dog", "Elephant")
	fmt.Println(tp)
	/*
		- Apple
		- Bat
		- Cat
		- Dog
		- Elephant
	*/

	tp.Reset()
	tp.SetStrategy(Html)
	tp.AppendList("Apple", "Bat", "Cat", "Dog", "Elephant")
	fmt.Println(tp)
	/*
		<ul>
		  <li>Apple</li>
		  <li>Bat</li>
		  <li>Cat</li>
		  <li>Dog</li>
		  <li>Elephant</li>
		</ul>
	*/
}
