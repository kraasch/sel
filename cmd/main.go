package main

import (
  "bufio"
  "fmt"
  "os"
  "flag"
  tea "github.com/charmbracelet/bubbletea"
)

var (
      result string
      isVerbose bool
)

type model struct {
  choices  []string         // List of items.
  cursor   int              // Cursor on the list.
  selected map[int]struct{} // List of selected list items.
}

func initialModel() model {
  initialChoices :=  []string{}
  scan := bufio.NewScanner(os.Stdin)
  scan.Split(bufio.ScanLines)
  for scan.Scan() {
    line := scan.Text()
    initialChoices = append(initialChoices, line)
  }
  return model{
    choices:  initialChoices,
    selected: make(map[int]struct{}),
  }
}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  decr := func() {
    if m.cursor > 0 {
      m.cursor--
    }
  }
  incr := func() {
    if m.cursor < len(m.choices)-1 {
      m.cursor++
    }
  }
  switch msg := msg.(type) {
    case tea.KeyMsg: // Is it a key press?
    switch msg.String() {
      case "ctrl-c": // Should exit the program without print.
      return m, tea.Quit
      case "q": // Should exit the program with print.
      for i, choice := range m.choices {
        if _, ok := m.selected[i]; ok {
          result += fmt.Sprintf("%s\n", choice)
        }
      }
      return m, tea.Quit
      case "up", "k", "ctrl-p": // Move the cursor up.
      decr()
      case "down", "j", "ctrl-n": // Move the cursor down.
      incr()
      case "n", "x": // De-select under cursor.
      delete(m.selected, m.cursor)
      incr()
      case "y", ".": // Select under cursor.
      m.selected[m.cursor] = struct{}{}
      incr()
      case "h", "?": // Toggle help.
      isVerbose = !isVerbose
      case "enter", " ": // Toggle selection.
      _, ok := m.selected[m.cursor]
      if ok {
        delete(m.selected, m.cursor)
      } else {
        m.selected[m.cursor] = struct{}{}
      }
    }
  }
  return m, nil
}

func (m model) View() string {
  s := ""
  if isVerbose {
    s = "Toggle (enter, space), select (y, .), unselect (n, x):\n\n" 
  }
  for i, choice := range m.choices { // Iterate over choices.
    cursor := " " // no cursor.
    if m.cursor == i {
      cursor = ">" // cursor.
    }
    checked := " " // not selected.
    if _, ok := m.selected[i]; ok {
      checked = "x" // selected.
    }
    s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
  }
  if isVerbose {
    s += "\nFinish (q), abort (ctrl-c), toggle help (h, ?).\n"
  }
  return s
}

func main() {
  verboseLong  := flag.Bool("verbose", false, "Print key info.")
  verboseShort := flag.Bool("v", false, "Print key info.")
  flag.Parse()
  isVerbose = *verboseLong || *verboseShort
  p := tea.NewProgram(initialModel())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error: %v", err)
    os.Exit(1)
  }
  fmt.Print(result)
}
