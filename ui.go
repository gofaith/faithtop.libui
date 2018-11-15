package faithtop

import (
	"github.com/andlabs/ui"
)

type IView interface {
	getBaseView() *FBaseView
}
type FBaseView struct {
	view ui.Control
}

func (v *FBaseView) getBaseView() *FBaseView {
	return v
}
