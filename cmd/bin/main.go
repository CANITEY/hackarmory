package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UiInterface int

const (
	DependencyInterface UiInterface = iota
	ToolsInterface
)

type Style struct {
	Success lipgloss.Style
	Failure lipgloss.Style
}

func NewModel() *Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return &Model {
		Spinner: s,
	}
}

// model function
type Model struct {
	Interface UiInterface
	Styles Style
	Spinner spinner.Model
	Index int
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(m.CheckDep(m.Index), m.Spinner.Tick)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "Q", "ctrl+c":
			fmt.Println("q")
			return m, tea.Quit
		default:
			return m, nil
		}
	case SuccessMsg:
		return m, nil
	case FailureMsg:
		return m, nil
	default:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}

func (m *Model) View() string {
	return m.Spinner.View()
}

// user's custom functions
type SuccessMsg string
type FailureMsg string

func (m *Model) CheckDep(index int) tea.Cmd {
	return func() tea.Msg {
		return SuccessMsg("XSXSXS")
	}
}

func main() {
	p := tea.NewProgram(NewModel())

	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
