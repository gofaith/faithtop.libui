package faithtop

import (
	"github.com/andlabs/ui"
)

type IView interface {
	getBaseView() *FBaseView
}
type FBaseView struct {
	view    ui.Control
	stretch bool
}

func (v *FBaseView) getBaseView() *FBaseView {
	return v
}
func (v *FBaseView) Expand() {
	v.stretch = true
}
