package main

import (
	"github.com/CANITEY/hackarmory/internal/ui/tools"
	tea "github.com/charmbracelet/bubbletea"
)

type UiInterface int

const (
	DependencyInterface UiInterface = iota
	ToolsInterface
)


func main() {
	p := tea.NewProgram(tools.NewModel())
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
