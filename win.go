package faithtop

import (
	"github.com/andlabs/ui"
)

type FWin struct {
	v         *ui.Window
	showAfter bool
}

func Win(title string, w, h int, hasMenubar bool) *FWin {
	v := &FWin{}
	v.v = ui.NewWindow(title, w, h, hasMenubar)
	return v
}
func (w *FWin) DeferShow() *FWin {
	w.showAfter = true
	return w
}
func (w *FWin) Add(i IView) *FWin {
	w.v.SetChild(i.getBaseView().view)
	if w.showAfter {
		w.Show()
	}
	return w
}
func (w *FWin) Show() *FWin {
	w.v.Show()
	return w
}
