package main

import (
	"fmt"
	"strings"

	"github.com/CANITEY/hackarmory/internal/checks"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/CANITEY/hackarmory/internal/ui/dependency"
)

type UiInterface int

const (
	DependencyInterface UiInterface = iota
	ToolsInterface
)


func main() {
	p := tea.NewProgram()

	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
