package faithtop

import (
	"github.com/andlabs/ui"
)

type FTextArea struct {
	FBaseView
	v *ui.MultilineEntry
}

func TextArea() *FTextArea {
	v := &FTextArea{}
	v.v = ui.NewMultilineEntry()
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FTextArea) Expand() *FTextArea {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (v *FTextArea) Text(t string) *FTextArea {
	v.v.SetText(t)
	return v
}

func (f *FTextArea) ReadOnly(b bool) *FTextArea {
	f.v.SetReadOnly(b)
	return f
}

func (f *FTextArea) GetText() string {
	return f.v.Text()
}
