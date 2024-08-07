package main

import (
	"github.com/CANITEY/hackarmory/internal/ui/dependencies"
	tea "github.com/charmbracelet/bubbletea"
)

type UiInterface int

const (
	DependencyInterface UiInterface = iota
	ToolsInterface
)


func main() {
	p := tea.NewProgram(dependencies.NewModel())
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
