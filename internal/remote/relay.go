package remote

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/slides/internal/model"
)

// CommandRelay is meant to expose slide interaction to external
// processes that can work as a remote for the slides.
type CommandRelay struct {
	*tea.Program
}

func NewCommandRelay(p *tea.Program) *CommandRelay {
	return &CommandRelay{
		Program: p,
	}
}

func (r *CommandRelay) SlideNext() {
	r.Send(model.RemoteMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'n'},
	})
}

func (r *CommandRelay) SlidePrev() {
	r.Send(model.RemoteMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'p'},
	})
}

func (r *CommandRelay) SlideFirst() {
	// Requires 2 keystrokes to actually
	// move to first slide
	r.Send(model.RemoteMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'g'},
	})
	r.Send(model.RemoteMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'g'},
	})
}

func (r *CommandRelay) SlideLast() {
	r.Send(model.RemoteMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'G'},
	})
}

func (r *CommandRelay) CodeExecute() {
	r.Send(model.RemoteMsg{
		Type: tea.KeyCtrlE,
	})
}

func (r *CommandRelay) CodeCopy() {
	r.Send(model.RemoteMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'y'},
	})
}

func (r *CommandRelay) Quit() {
	r.Send(model.RemoteMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'q'},
	})
}
