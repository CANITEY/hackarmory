package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/CANITEY/hackarmory/internal/ui/dependencies"
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
