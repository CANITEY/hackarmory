package installer

import (
	"strings"
	"time"

	"github.com/CANITEY/hackarmory/internal/helpers"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

// tools status 0 not installed, 1 installing, 2 installed
type Model struct {
	tools map[string]int
	progress progress.Model
	queue []string
	index int
}

func (m Model) Init() tea.Cmd {
	return m.tickCmd(m.index)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding * 2
		if m.progress.Width > maxWidth {
			m.progress.Width = msg.Width - 2 * padding
		}
		return m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		cmd := m.progress.IncrPercent(0.25)
		if m.index == len(m.queue) - 1 {

		} else {
			m.index++
		}
		return m, tea.Batch(m.tickCmd(m.index), cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m Model) View() string {
	buf := strings.Builder{}
	pad := strings.Repeat(" ", padding)
	progress := "\n" +
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")

	toolsLogger := helpers.NewToolLogger(m.tools)
	buf.WriteString(toolsLogger.Log())
	buf.WriteString(progress)
	return buf.String()
}

// user defined functions
func (m *Model) tickCmd(index int) tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		m.tools[m.queue[index]] = 1
		return tickMsg(t)
	})
}

func NewModel(tools []string) *Model {
	toolsMap := helpers.StoMInt(tools, -1)
	return &Model{
		progress: progress.New(progress.WithDefaultGradient()),
		tools:    toolsMap,
		queue: tools,
	}
}
