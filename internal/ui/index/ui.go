package index

import tea "github.com/charmbracelet/bubbletea"

type Model struct{}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch /* msg :=  */msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	}
	return m, nil
}

func (m *Model) View() string {
	return "Hello world"
}
