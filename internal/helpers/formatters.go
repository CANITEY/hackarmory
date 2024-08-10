package helpers

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func NewToolLogger(tools map[string]int) *ToolLogger {
	waiting := 		lipgloss.NewStyle().Foreground(lipgloss.Color("25")).PaddingTop(1).Render
	installing := 	lipgloss.NewStyle().Background(lipgloss.Color("12")).PaddingTop(1).Render
	done := 		lipgloss.NewStyle().Foreground(lipgloss.Color("2")).PaddingTop(1).Render
	error := 		lipgloss.NewStyle().Foreground(lipgloss.Color("1")).PaddingTop(1).Render
	style := style{
		waiting:    waiting,
		done:       done,
		installing: installing,
		err:        error,
	}
	return &ToolLogger{
		Tools: tools,
		style: style,
	}
}

type style struct {
	waiting, installing, done, err func(...string) string
}

type ToolLogger struct {
	Tools map[string]int
	style style
}

func (t *ToolLogger) Log() string {
	buf := strings.Builder{}

	tools := []string{}

	for tool := range t.Tools {
		tools = append(tools, tool)
	}

	slices.Sort(tools)

	for _, tool := range tools {
		switch t.Tools[tool] {
		case -1:
			buf.WriteString(t.style.waiting(fmt.Sprintf("[W] %v\n", tool)))
		case 0:
			buf.WriteString(t.style.installing(fmt.Sprintf("[I] %v\n", tool)))
		case 1:
			buf.WriteString(t.style.done(fmt.Sprintf("[D] %v\n", tool)))
		case 2:
			buf.WriteString(t.style.err(fmt.Sprintf("[E] %v\n", tool)))
		}
	}
	return buf.String()
}
