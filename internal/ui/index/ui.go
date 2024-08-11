package index

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UiInterface int

const (
	DependencyInterface UiInterface = iota
	ToolsInterface
	InstallerInterface
)


type Model struct{
	width int
	version string
	ui UiInterface
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg :=msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tea.WindowSizeMsg:
		m.width = msg.Width
		return m, nil
	}
	return m, nil
}

func (m *Model) View() string {
	logo := `
.##.....##....###.....######..##....##....###....########..##.....##..#######..########..##....##
.##.....##...##.##...##....##.##...##....##.##...##.....##.###...###.##.....##.##.....##..##..##.
.##.....##..##...##..##.......##..##....##...##..##.....##.####.####.##.....##.##.....##...####..
.#########.##.....##.##.......#####....##.....##.########..##.###.##.##.....##.########.....##...
.##.....##.#########.##.......##..##...#########.##...##...##.....##.##.....##.##...##......##...
.##.....##.##.....##.##....##.##...##..##.....##.##....##..##.....##.##.....##.##....##.....##...
.##.....##.##.....##..######..##....##.##.....##.##.....##.##.....##..#######..##.....##....##...
`


	buf := strings.Builder{}
	buf.WriteString(logo)
	footer := fmt.Sprintf("\nMade By CANITEY, Version %v", m.version)
	buf.WriteString(footer)
	styledFrame := lipgloss.NewStyle().
		Width(m.width - 3).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		AlignHorizontal(lipgloss.Center).
		Render(buf.String())

	view := strings.Builder{}
	view.WriteString(styledFrame)
	view.WriteString("\nDUMMY DATA")
	return view.String()
}

func NewModel(version string) *Model {
	return &Model{
		version: version,
	}
}
