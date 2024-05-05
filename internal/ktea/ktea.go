package ktea

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var selection string = ""

type model struct {
	table  table.Model
	width  int
	height int
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, tea.Quit
		case "enter":
			selection = m.table.SelectedRow()[0]
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	//    return baseStyle.Render(m.table.View()) + "\n"
	if m.width == 0 {
		return ""
	}

	table := baseStyle.Render(m.table.View())
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, table)
}

func center(s string, w int) string {
	if len(s) >= w {
		return s
	}
	n := w - len(s)
	div := n / 2
	return strings.Repeat(" ", div) + s + strings.Repeat(" ", div)
}

func Ktea(strFlag string, myDir string) {
	// files, err := os.ReadDir(os.Getenv("HOME") + "/.kube")
	// files, err := os.ReadDir(myDir)
	// if err != nil {
	//		log.Fatal(err)
	//	}

	centeredTitle := center("Kube Configs", 30)

	columns := []table.Column{
		{Title: centeredTitle, Width: 30},
	}

	rows := []table.Row{}

	//	for _, file := range files {
	//		if file.Type().IsRegular() {
	//			rows = append(rows, table.Row{file.Name()})
	//		}
	//	}

	dirs, err := os.ReadDir(myDir)
	if err != nil {
		log.Println(err)
	}

	for _, e := range dirs {
		if e.IsDir() {
			rows = append(rows, table.Row{e.Name()})
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("240")).
		Background(lipgloss.Color("255")).
		Bold(false)
	t.SetStyles(s)

	m := model{t, 0, 0}
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if selection != "" {
		if strFlag == "env" {
			fmt.Printf("Setting ${KUBECONFIG} to ${HOME}/.kube/%s\n", selection)
			os.Setenv("KUBECONFIG", os.Getenv("HOME")+"/.kube"+selection)
			fmt.Println("KUBECONFIG:", os.Getenv("KUBECONFIG"))
		} else {
			if _, err := os.Lstat(os.Getenv("HOME") + "/.kube/config"); err == nil {
				os.Remove(os.Getenv("HOME") + "/.kube/config")
			}
			fmt.Printf("Linking ${HOME}/.kube/conifg -> %s\n", selection)
			err := os.Symlink(os.Getenv("HOME")+"/.kube/"+selection, os.Getenv("HOME")+"/.kube/config")
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		fmt.Println("No changes made.")
	}
}
