package faithtop

import (
	"github.com/andlabs/ui"
)

type FTab struct {
	FBaseView
	v *ui.Tab
}

func Tab() *FTab {
	v := &FTab{}
	v.v = ui.NewTab()
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FTab) Expand() *FTab {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (f *FTab) Append(title string, i IView) *FTab {
	f.v.Append(title, i.getBaseView().view)
	return f
}
