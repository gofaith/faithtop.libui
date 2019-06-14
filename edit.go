package faithtop

import (
	"github.com/andlabs/ui"
)

type FEdit struct {
	FBaseView
	v *ui.Entry
}

func Edit() *FEdit {
	v := &FEdit{}
	v.v = ui.NewEntry()
	v.view = v.v
	return v
}

func PasswordEdit() *FEdit {
	v := &FEdit{}
	v.v = ui.NewPasswordEntry()
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FEdit) Expand() *FEdit {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (v *FEdit) Text(t string) *FEdit {
	v.v.SetText(t)
	return v
}

func (f *FEdit) GetText() string {
	return f.v.Text()
}
