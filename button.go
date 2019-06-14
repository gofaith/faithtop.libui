package faithtop

import (
	"github.com/andlabs/ui"
)

type FButton struct {
	FBaseView
	v *ui.Button
}

func Button() *FButton {
	v := &FButton{}
	v.v = ui.NewButton("")
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FButton) Expand() *FButton {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (v *FButton) Text(t string) *FButton {
	v.v.SetText(t)
	return v
}

func (f *FButton) GetText() string {
	return f.v.Text()
}

func (v *FButton) OnClick(f func()) *FButton {
	v.v.OnClicked(func(*ui.Button) {
		f()
	})
	return v
}

func (v *FButton) Enabled(b bool) *FButton {
	if b {
		v.v.Enable()
	} else {
		v.v.Disable()
	}
	return v
}
