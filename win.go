package faithtop

import (
	"github.com/andlabs/ui"
)

type FWin struct {
	v         *ui.Window
	showAfter bool
	onClosing func() bool
}

func Win(w, h int) *FWin {
	v := &FWin{}
	v.v = ui.NewWindow("", w, h, true)
	v.v.OnClosing(func(*ui.Window) bool {
		if v.onClosing != nil {
			return v.onClosing()
		}
		return true
	})
	return v
}
func (v *FWin) Title(t string) *FWin {
	v.v.SetTitle(t)
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
func (v *FWin) OnClose(f func() bool) *FWin {
	v.onClosing = f
	return v
}
func (v *FWin) MainQuitOnClose() *FWin {
	v.onClosing = func() bool {
		MainQuit()
		return true
	}
	return v
}
func (v *FWin) VBox(is ...IView) *FWin {
	v.Add(VBox().Append(is...))
	return v
}

func (v *FWin) HBox(is ...IView) *FWin {
	v.Add(HBox().Append(is...))
	return v
}
