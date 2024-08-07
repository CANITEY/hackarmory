package dependencies

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/CANITEY/hackarmory/internal/checks"
)

type Style struct {
	Success lipgloss.Style
	Failure lipgloss.Style
}

func NewModel() *Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))

	style := Style{
		Success: lipgloss.NewStyle().Foreground(lipgloss.Color("2")),
		Failure: lipgloss.NewStyle().Foreground(lipgloss.Color("1")),
	}

	dependencies := []string{
		"python",
		"go",
		"java",
		"gcc",
		"g++",
	}

	return &Model {
		Spinner: s,
		Dependencies: dependencies,
		Styles: style,
	}
}

// model function
type Model struct {
	Styles Style
	Spinner spinner.Model
	Index int
	Dependencies []string
	SuccededDep []string
	FailedDep []string
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(m.CheckDep(m.Index), m.Spinner.Tick)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "Q", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}
	case DepMessage:
		if m.Index < len(m.Dependencies) - 1 {
			m.Index++
		} else {
			break
		}
		return m, m.CheckDep(m.Index)
	default:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *Model) View() string {
	buf := strings.Builder{}
	buf.WriteString(m.FormatDeps())
	buf.WriteString(fmt.Sprintf("%v Checking depenedencies", m.Spinner.View()))
	return buf.String()
}

// user's custom functions
type DepMessage string

func (m *Model) CheckDep(index int) tea.Cmd {
	return func() tea.Msg {
		dependency := m.Dependencies[m.Index]
		success, fail := checks.CheckDependency(dependency)
		if fail != nil {
			m.FailedDep = append(m.FailedDep, dependency)
			return DepMessage(dependency)
		} else if success != "" {
			m.SuccededDep = append(m.SuccededDep, dependency)
			return DepMessage(dependency)
		}
		return DepMessage(dependency)
	}
}

func (m *Model) FormatDeps() string {
	buf := strings.Builder{}
	for _, dep := range m.SuccededDep {
		buf.WriteString(fmt.Sprintf("%v %v\n", m.Styles.Success.Render("[✔]"), dep))
	}
	for _, dep := range m.FailedDep {
		buf.WriteString(fmt.Sprintf("%v %v\n", m.Styles.Failure.Render("[x]"), dep))
	}

	return buf.String()
}
