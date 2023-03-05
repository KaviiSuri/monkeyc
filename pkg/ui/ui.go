package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	width = 60.
	// charsPerWord is the average characters per word used by most typing tests
	// to calculate your WPM score.
	charsPerWord = 5.
)

type tickMsg time.Time

var (
	wpms []float64
)

var f, _ = tea.LogToFile("debug.log", "debug")

type Ui struct {
	Text     []rune
	Typed    []rune
	Start    time.Time
	Mismatch int
	Score    float64
	Progress progress.Model
}

func (m Ui) updateProgress() (tea.Model, tea.Cmd) {
	cmd := m.Progress.SetPercent(float64(len(m.Typed)) / float64(len(m.Text)))

	if m.Progress.Percent() >= 1.0 {
		return m, tea.Quit
	}

	return m, cmd
}

// Init implements tea.Model
func (Ui) Init() tea.Cmd {
	return tickCmd()
}

// Update implements tea.Model
func (m Ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Start.IsZero() {
			m.Start = time.Now()
		}

		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		if msg.Type == tea.KeyBackspace && len(m.Typed) > 0 {
			m.Typed = m.Typed[:len(m.Typed)-1]
			return m.updateProgress()
		}

		if msg.Type != tea.KeyRunes && msg.Type != tea.KeySpace {
			return m, nil
		}

		char := msg.Runes[0]
		next := rune(m.Text[len(m.Typed)])

		// A newline character should be entered regarless of the user typing somethign else
		// so as to not break the interface
		if next == '\n' {
			m.Typed = append(m.Typed, next)

			// if the user enters a space character, it should be ignored and treated as a newline itself
			if char == ' ' {
				return m, nil
			}
		}

		m.Typed = append(m.Typed, msg.Runes...)

		if char == next {
			m.Score += 1
		}

		return m.updateProgress()
	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - 4
		if m.Progress.Width > width {
			m.Progress.Width = width
		}
		return m, nil

	case tickMsg:

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		m, cmd := m.updateProgress()
		return m, tea.Batch(tickCmd(), cmd)
	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

// View implements tea.Model
func (m Ui) View() string {
	remaining := m.Text[len(m.Typed):]

	var typed string
	for i, c := range m.Typed {
		if c == rune(m.Text[i]) {
			typed += TypedStyle.Render(string(c))
		} else {
			typed += ErrorStyle.Render(string(m.Text[i]))
		}
	}

	s := fmt.Sprintf("\n %s\n\n%s", m.Progress.View(), typed)

	if len(remaining) > 0 {
		s += CurrentStyle.Render(string(remaining[:1]))
		s += UnTypedStyle.Render(string(remaining[1:]))
	}

	var wpm float64
	if len(m.Typed) > 1 {
		wpm = (m.Score / charsPerWord) / (float64(time.Since(m.Start).Minutes()))
	}

	s += fmt.Sprintf("\n\nWPM: %v", int(wpm))

	if len(m.Typed) > charsPerWord {
		wpms = append(wpms, wpm)
	}

	wpmsCount := wpms
	if len(wpmsCount) <= 0 {
		wpmsCount = []float64{0}
	}

	return s
}

var _ tea.Model = Ui{}

func New(text string) Ui {
	progress := progress.New(progress.WithDefaultGradient())
	return Ui{Progress: progress, Text: []rune(text)}
}
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
