package main

import "fmt"

// Subsystem: Buffer
type Buffer struct {
	content string
}

func (b *Buffer) Write(text string) {
	b.content += text
}

// Subsystem: Viewport
type Viewport struct {
	buffer *Buffer
	width  int
	height int
}

func NewViewport(buffer *Buffer, width, height int) *Viewport {
	return &Viewport{
		buffer: buffer,
		width:  width,
		height: height,
	}
}

func (v *Viewport) Render() string {
	return fmt.Sprintf("Viewport (%dx%d):\n%s", v.width, v.height, v.buffer.content)
}

// Facade: Console
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	active    int
}

func NewConsole() *Console {
	// Initialize with a single buffer and viewport
	buffer := &Buffer{}
	viewport := NewViewport(buffer, 80, 40)

	return &Console{
		buffers:   []*Buffer{buffer},
		viewports: []*Viewport{viewport},
		active:    0,
	}
}

func (c *Console) CreateBuffer() {
	buffer := &Buffer{}
	c.buffers = append(c.buffers, buffer)
}

func (c *Console) CreateViewport(width, height int) {
	buffer := c.buffers[len(c.buffers)-1]
	viewport := NewViewport(buffer, width, height)
	c.viewports = append(c.viewports, viewport)
	c.active = len(c.viewports) - 1
}

func (c *Console) SwitchViewport(index int) {
	if index >= 0 && index < len(c.viewports) {
		c.active = index
	}
}

func (c *Console) Write(text string) {
	buffer := c.buffers[c.active]
	buffer.Write(text)
}

func (c *Console) Render() string {
	viewport := c.viewports[c.active]
	return viewport.Render()
}

// Client code
func main() {
	console := NewConsole()

	// Writing to the initial viewport
	console.Write("Hello from the initial viewport!\n")

	// Creating a new buffer and viewport
	console.CreateBuffer()
	console.CreateViewport(120, 60)

	// Writing to the new viewport
	console.SwitchViewport(1)
	console.Write("Hello from the new viewport!\n")

	// Rendering the current viewport
	result := console.Render()
	fmt.Println(result)
	/*
		Viewport (120x60):
		Hello from the new viewport!
	*/
}
