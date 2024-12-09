// This file contains a generic Button component.

// Import necessary packages
import (
	// BubbbleTea (terminal UI) and LipGloss (styling)
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Sets the package to "Components"
package components

// Returns a RGB lipgloss.Color from a [3]int array (RGB)
func chcol(col: [3]int) lipgloss.Color {
	return lipgloss.Color(col[0], col[1], col[2])
}

// Interface for a button component
type Button struct {
	label string
	color [3]int  // RGB color in a [3]int array
	onClick func()
	width int
	height int
}

// Returns a new Button component
func NewButton(label string, color [3]int, onClick func(), width int, height int) *Button {
	return &Button {
		label: label,
		color: color,
		onClick: onClick,
		width: width,
		height: height,
	}
}

// Renders the button component
func (b *Button) Render() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, b.label, lipgloss.NewStyle().Foreground(chcol(b.color)).Render(b.label))
}
