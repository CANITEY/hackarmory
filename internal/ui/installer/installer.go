package installer

import (
	"github.com/charmbracelet/bubbletea"
)

type Model struct {}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m *Model) View() string {
	return "Hello world"
}
