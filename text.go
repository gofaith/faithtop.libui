package faithtop

import (
	"github.com/andlabs/ui"
)

type FText struct {
	FBaseView
	v *ui.Label
}

func Text(s string) *FText {
	v := &FText{}
	v.v = ui.NewLabel(s)
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FText) Expand() *FText {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (v *FText) Text(t string) *FText {
	v.v.SetText(t)
	return v
}

func (v *FText) GetText() string {
	return v.v.Text()
}
