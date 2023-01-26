package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	//"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}
}

type keymap struct {
	start key.Binding 
	input key.Binding
	table key.Binding
	quit  key.Binding
}

type model struct {
	// Screens include start, timer, input, bye
	screenSelected string      	   // The selected screen
	stopwatch      stopwatch.Model
	keymap         keymap
	help 		   help.Model
}

func initialModel() model {
	m := model {
		screenSelected: "start",
		stopwatch: stopwatch.NewWithInterval(time.Second),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys(" ", "enter"),
				key.WithHelp("space/enter", "start/pause"),
			),
			input: key.NewBinding(
				key.WithKeys("i"),
				key.WithHelp("i", "input entry"),
			),
			table: key.NewBinding(
				key.WithKeys("v"),
				key.WithHelp("v", "view weekly table"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q/ctrl+c", "quit"),
			),
		},
		help: help.NewModel(),
	}

	m.keymap.start.SetEnabled(false)

	return m
}

func (m model) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.screenSelected = "quit"
			return m, tea.Quit
		case key.Matches(msg, m.keymap.start):
			m.screenSelected = "timer"
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.screenSelected
}
