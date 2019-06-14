package faithtop

import (
	"github.com/andlabs/ui"
)

type FCheck struct {
	FBaseView
	v *ui.Checkbox
}

func Check() *FCheck {
	v := &FCheck{}
	v.v = ui.NewCheckbox("")
	v.view = v.v
	return v
}

// ------------------------------------------
func (v *FCheck) Expand() *FCheck {
	v.FBaseView.Expand()
	return v
}

// ==========================================
func (v *FCheck) Text(t string) *FCheck {
	v.v.SetText(t)
	return v
}
func (v *FCheck) OnCheck(fn func(bool)) *FCheck {
	v.v.OnToggled(func(*ui.Checkbox) {
		fn(v.v.Checked())
	})
	return v
}

func (f *FCheck) Checked(b bool) *FCheck {
	f.v.SetChecked(b)
	return f
}

func (f *FCheck) GetChecked() bool {
	return f.v.Checked()
}