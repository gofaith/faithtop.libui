package faithtop

import (
	"github.com/andlabs/ui"
)

type FCounter struct {
	FBaseView
	v *ui.Spinbox
}

func Counter(min, max int) *FCounter {
	v := &FCounter{}
	v.v = ui.NewSpinbox(min, max)
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FCounter) Expand() *FCounter {
	v.FBaseView.Expand()
	return v
}

// ==========================================

func (f *FCounter) GetNum() int {
	return f.v.Value()
}

func (f *FCounter) Num(i int) *FCounter {
	f.v.SetValue(i)
	return f
}

func (f *FCounter) OnChange(fn func(int)) *FCounter {
	f.v.OnChanged(func(*ui.Spinbox) {
		fn(f.v.Value())
	})
	return f
}
