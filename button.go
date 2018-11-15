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
func (v *FButton) OnClick(f func()) *FButton {
	v.v.OnClicked(func(*ui.Button) {
		f()
	})
	return v
}
