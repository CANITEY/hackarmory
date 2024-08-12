package tools

import (
	"fmt"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var defaultTools = []string{
	"subfinder",
	"sublist3r",
	"amass",
	"nmap",
	"wireshark",
	"postman",
	"burpsuite",
	"fuff",
	"dirsearch",
	"dirb",
	"seclists",
	"gobuster",
	"zap",
	"wfuzz",
	"sublime",
	"firefox",
	"waybackurls",
	"httpx",
	"httpprobe",
	"gxss",
	"sqlmap",
	"subzy",
}

// model initialization
func NewModel(width int) *Model {
	slices.Sort(defaultTools)
	return &Model{
		Tools: defaultTools,
		Selected: make(map[string]struct{}),
		Width: width,
	}
}

type Model struct {
	Tools    []string
	Selected map[string]struct{}
	Cursor   int
	Width    int
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "Q", "ctrl+c":
			return m, tea.Quit
		case "k", "K", "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
			return m, nil
		case "j", "J", "down":
			if m.Cursor < len(m.Tools)-1 {
				m.Cursor++
			}
			return m, nil
		case " ":
			if _, ok := m.Selected[m.Tools[m.Cursor]]; ok {
				delete(m.Selected, m.Tools[m.Cursor])
			} else {
				m.Selected[m.Tools[m.Cursor]] = struct{}{}
			}
			return m, nil
		case "enter":
			return m, nil
		default:
			return m, nil
		}
	default:
		return m, nil
	}
	return m, nil
}

func (m *Model) View() string {
	list := strings.Builder{}
	for index, item := range m.Tools {
		cursor := " "
		if m.Cursor == index {
			cursor = ">"
		}

		selected := " "
		if _, ok := m.Selected[m.Tools[index]]; ok {
			selected = "✔"
		}

		list.WriteString(fmt.Sprintf("%v[%v] %v\n", cursor, selected, item))
	}
	listMenu := lipgloss.NewStyle().
		Width(m.Width - 3).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Render(strings.Trim(list.String(), "\n"))
	help := lipgloss.NewStyle().
		Width(m.Width - 3).
		BorderStyle(lipgloss.NormalBorder()).
		PaddingLeft(1).
		BorderForeground(lipgloss.Color("63")).
		Render("j, J, down to go down, k, K, up to go up; space to toggle selection; enter to go to installation step; Q, q, ctrl+c to quit")

	return lipgloss.JoinVertical(lipgloss.Center, listMenu, help)
}
