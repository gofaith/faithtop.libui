package faithtop

import (
	"github.com/andlabs/ui"
)

type FCombo struct {
	FBaseView
	ss []string
	v  *ui.Combobox
}

func Combo() *FCombo {
	v := &FCombo{}
	v.v = ui.NewCombobox()
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FCombo) Expand() *FCombo {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (f *FCombo) Append(s ...string) *FCombo {
	for _, i := range s {
		f.v.Append(i)
		f.ss = append(f.ss, i)
	}
	return f
}

func (f *FCombo) OnSelected(fn func(int, string)) *FCombo {
	f.v.OnSelected(func(*ui.Combobox) {
		fn(f.v.Selected(), f.ss[f.v.Selected()])
	})
	return f
}

func (f *FCombo) Select(i int) *FCombo {
	f.v.SetSelected(i)
	return f
}

func (f *FCombo) GetSelected() int {
	return f.v.Selected()
}

func (f *FCombo) GetSelectedString() string {
	return f.ss[f.v.Selected()]
}
