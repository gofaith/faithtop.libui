package faithtop

import (
	"github.com/andlabs/ui"
)

type FRadios struct {
	FBaseView
	v *ui.RadioButtons
}

func Radios() *FRadios {
	v := &FRadios{}
	v.v = ui.NewRadioButtons()
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FRadios) Expand() *FRadios {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (f *FRadios) Append(texts ...string) *FRadios {
	for _, text := range texts {
		f.v.Append(text)
	}
	return f
}

func (f *FRadios) OnChange(fn func(i int)) *FRadios {
	f.v.OnSelected(func(*ui.RadioButtons) {
		fn(f.v.Selected())
	})
	return f
}

func (f *FRadios) Select(i int) *FRadios {
	f.v.SetSelected(i)
	return f
}

func (f *FRadios) GetSelected() int {
	return f.v.Selected()
}
