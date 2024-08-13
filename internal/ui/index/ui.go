package index

import (
	"fmt"
	"strings"

	"github.com/CANITEY/hackarmory/internal/messages"
	"github.com/CANITEY/hackarmory/internal/ui/dependencies"
	"github.com/CANITEY/hackarmory/internal/ui/installer"
	"github.com/CANITEY/hackarmory/internal/ui/tools"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width    int
	version  string
	external tea.Model
}

func (m *Model) Init() tea.Cmd {
	return m.external.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd
		m.external, cmd = m.external.Update(msg)
		return m, cmd
	case messages.Next:
		m.external = tools.NewModel(m.width)
		return m, nil
	case messages.Install:
		m.external = installer.NewModel(msg, m.width)
		cmd := m.external.Init()
		return m, cmd
	case installer.TickMsg, progress.FrameMsg:
		var cmd tea.Cmd
		m.external, cmd = m.external.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.width = msg.Width
		var cmd tea.Cmd
		m.external, cmd = m.external.Update(msg)
		return m, cmd
	default:
		var cmd tea.Cmd
		m.external, cmd = m.external.Update(msg)
		return m, cmd
	}
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


	return lipgloss.JoinVertical(lipgloss.Center, styledFrame, m.external.View())
}

func NewModel(version string) *Model {
	return &Model{
		version:  version,
		external: dependencies.NewModel(),
	}
}
