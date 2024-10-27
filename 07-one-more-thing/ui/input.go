package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel(prompt string) model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 255
	ti.Width = 80

	ti.Prompt = prompt

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// Handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

var promptStyle lipgloss.Style
func (m model) View() string {
	return promptStyle.Render(m.textInput.View()+"\n") 
}


func Input(color, prompt string) (string, error) {
	promptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
	p := tea.NewProgram(initialModel(prompt))
	m, err := p.Run()
	if err != nil {
		return "", err
	}
	if m, ok := m.(model); ok {
		return strings.TrimSpace(m.textInput.Value()), nil
	}
	return "", fmt.Errorf("😡 unable to get input")
}

// Println prints the given strings with a specific style and adds a newline at the end.
func Println(color string, strs ...interface{}) {
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
	
	// Convert all arguments to strings
	strSlice := make([]string, len(strs))
	for i, v := range strs {
		strSlice[i] = fmt.Sprint(v)
	}
	
	// Join all strings and render with the style
	renderedString := textStyle.Render(strings.Join(strSlice, " "))
	
	// Print the rendered string with a newline
	fmt.Println(renderedString)
}

func Print(color string, strs ...interface{}) {
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
	
	// Convert all arguments to strings
	strSlice := make([]string, len(strs))
	for i, v := range strs {
		strSlice[i] = fmt.Sprint(v)
	}
	
	// Join all strings and render with the style
	renderedString := textStyle.Render(strings.Join(strSlice, " "))
	
	// Print the rendered string with a newline
	fmt.Print(renderedString)
}
