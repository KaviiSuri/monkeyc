package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var TypedStyle = lipgloss.
	NewStyle().
	Foreground(lipgloss.Color("#fff"))

var UnTypedStyle = lipgloss.
	NewStyle().
	Foreground(lipgloss.Color("#555")).
	Faint(true)

var CurrentStyle = UnTypedStyle.
	Underline(true)

var ErrorStyle = lipgloss.
	NewStyle().
	Foreground(lipgloss.Color("#fff")).
	Background(lipgloss.Color("#f33"))
