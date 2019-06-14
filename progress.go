package faithtop

import (
	"github.com/andlabs/ui"
)

type FProgress struct {
	FBaseView
	v *ui.ProgressBar
}

func Progress() *FProgress {
	v := &FProgress{}
	v.v = ui.NewProgressBar()
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FProgress) Expand() *FProgress {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (f *FProgress) SetValue(i int) *FProgress {
	f.v.SetValue(i)
	return f
}

func (f *FProgress) SetBussy() *FProgress {
	f.v.SetValue(-1)
	return f
}
