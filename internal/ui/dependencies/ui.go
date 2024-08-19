package dependencies

import (
	"fmt"
	"strings"

	"github.com/CANITEY/hackarmory/internal/checks"
	"github.com/CANITEY/hackarmory/internal/messages"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
		"gem",
	}

	return &Model{
		Spinner:      s,
		Dependencies: dependencies,
		Styles:       style,
	}
}

// model function
type Model struct {
	Styles       Style
	Spinner      spinner.Model
	Index        int
	Dependencies []string
	SuccededDep  []string
	FailedDep    []string
	width        int
	endMsg string
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(m.CheckDep(m.Index), m.Spinner.Tick)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "Q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.IsComplete() == ValidationMsg(true) {
				return m, m.Next
			} else {
				return m, nil
			}
		default:
			return m, nil
		}
	case DepMessage:
		if m.Index < len(m.Dependencies)-1 {
			m.Index++
		} else {
			return m, m.IsComplete
		}
		return m, m.CheckDep(m.Index)
	case ValidationMsg:
		if ok := msg; !ok {
			m.endMsg = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render("There are some missing dependencies, fix them then restart me (press Q to quit)")
			return m, nil
		} else {
			m.endMsg = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render("All is good press 'enter' to continue")
			return m, nil
		}
	default:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}

func (m *Model) View() string {
	depsFormat := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		PaddingLeft(1).
		Width(m.width - 3).
		BorderForeground(lipgloss.Color("63")).
		Render(strings.TrimSpace(m.FormatDeps()))

	spinnerFormat := lipgloss.NewStyle().
		Width(m.width - 3).
		BorderStyle(lipgloss.NormalBorder()).
		PaddingLeft(1).
		BorderForeground(lipgloss.Color("63")).
		Render(fmt.Sprintf("%v Checking depenedencies", m.Spinner.View()))

	endMsg := lipgloss.NewStyle().
		Width(m.width - 3).
		BorderStyle(lipgloss.NormalBorder()).
		PaddingLeft(1).
		BorderForeground(lipgloss.Color("63")).
		Render(m.endMsg)

	if m.endMsg == "" {
		return lipgloss.JoinVertical(lipgloss.Center, depsFormat, spinnerFormat)
	} else {
		return lipgloss.JoinVertical(lipgloss.Center, depsFormat, endMsg)
	}
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

type ValidationMsg bool

func (m *Model) IsComplete() tea.Msg {
	if len(m.FailedDep) > 0 {
		return ValidationMsg(false)
	}
	return ValidationMsg(true)
}


func (m *Model) Next() tea.Msg {
	return messages.Next(true)
}
