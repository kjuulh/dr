package pages

import tea "github.com/charmbracelet/bubbletea"

type ChangePage interface {
	Page() string
}

type changePage struct {
	page string
}

func NewChangePage(page string) tea.Cmd {
	return func() tea.Msg {
		return &changePage{
			page: page,
		}
	}
}

func (changePage *changePage) Page() string {
	return changePage.page
}

var _ tea.Msg = &changePage{}
var _ ChangePage = &changePage{}
