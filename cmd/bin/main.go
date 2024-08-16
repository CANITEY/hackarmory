package main

import (
	"github.com/CANITEY/hackarmory/internal/installs"
	// "github.com/CANITEY/hackarmory/internal/ui/index"
	// tea "github.com/charmbracelet/bubbletea"
)

// link all UIs together
func main() {
	// p := tea.NewProgram(index.NewModel("1.0.0"), tea.WithAltScreen())
	// _, err := p.Run()
	// if err != nil {
	// 	panic(err)
	// }
	installs.CreateToolsDir()
}
