package faithtop

import (
	"github.com/andlabs/ui"
)

type FWin struct {
	v              *ui.Window
	showAfter      bool
	isClosedByCode bool
	fnsOnClose     []func() bool
}

func WinDefault() *FWin {
	return Win(400, 300)
}

func Win(w, h int) *FWin {
	v := &FWin{}
	v.v = ui.NewWindow("", w, h, true)
	v.v.OnClosing(func(*ui.Window) bool {
		v.executeCloseFns()
		return true
	})
	return v
}

func ShowWin(i IView) *FWin {
	return WinDefault().DeferShow().Add(i)
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

func (w *FWin) Close() *FWin {
	w.isClosedByCode = true
	w.executeCloseFns()
	w.v.Destroy()
	w.isClosedByCode = false
	return w
}

func (f *FWin) executeCloseFns() {
	for _, fn := range f.fnsOnClose {
		if fn != nil {
			fn()
		}
	}
}

func (v *FWin) OnClose(f func() bool) *FWin {
	v.fnsOnClose = append(v.fnsOnClose, f)
	return v
}
func (v *FWin) QuitOnCloseClicked() *FWin {
	fn := func() bool {
		if v.isClosedByCode {
			return true
		}
		MainQuit()
		return true
	}
	v.fnsOnClose = append(v.fnsOnClose, fn)
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
func (v *FWin) Margin() *FWin {
	v.v.SetMargined(true)
	return v
}
func (v *FWin) Bordless() *FWin {
	v.v.SetBorderless(true)
	return v
}

// popup
func PopupDefault() *FWin {
	return Popup(200, 100)
}

func Popup(w, h int) *FWin {
	return Win(w, h).Bordless()
}

// msg window

func ShowInfo(w *FWin, title, content string) {
	ui.MsgBox(w.v, title, content)
}

func ShowErr(w *FWin, title, content string) {
	ui.MsgBoxError(w.v, title, content)
}
