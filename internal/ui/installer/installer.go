package installer

import (
	"slices"
	"strings"
	"time"

	"github.com/CANITEY/hackarmory/internal/helpers"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TickMsg time.Time

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

// tools status 0 not installed, 1 installing, 2 installed
type Model struct {
	tools    map[string]int
	progress progress.Model
	queue    []string
	index    int
	width    int
}

func (m Model) Init() tea.Cmd {
	return m.tickCmd(m.index)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String(){
			case "Q", "q", "ctrl+c":
			return m, tea.Quit
		}
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.progress.Width = msg.Width - padding*2
		if m.progress.Width > maxWidth {
			m.progress.Width = msg.Width - 2*padding
		}
		return m, nil

	case TickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		cmd := m.progress.SetPercent((float64(m.index) + 1) / float64(len(m.queue)))
		if m.index == len(m.queue)-1 {
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
	pad := strings.Repeat(" ", padding)

	progress := lipgloss.NewStyle().
		Width(m.width - 3).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Render("\n" + " " + m.progress.View() + "\n\n" + pad + helpStyle("Press any key to quit"))

	toolsLogger := helpers.NewToolLogger(m.tools)
	log := lipgloss.NewStyle().
		Width(m.width - 3).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Render(toolsLogger.Log())

	return lipgloss.JoinVertical(lipgloss.Center, log, progress)
}

// user defined functions
func (m *Model) tickCmd(index int) tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		m.tools[m.queue[index]] = 1
		return TickMsg(t)
	})
}

func NewModel(tools []string, width int) *Model {
	toolsMap := helpers.StoMInt(tools, -1)
	m := &Model{}
	m.width = width
	m.progress = progress.New(progress.WithDefaultGradient())
	m.progress.Width = width - padding*2
	if m.progress.Width > maxWidth {
		m.progress.Width = width - 2*padding
	}

	m.tools = toolsMap
	slices.Sort(tools)
	m.queue = tools

	return m
}
