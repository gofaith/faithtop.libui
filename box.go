package faithtop

import (
	"github.com/andlabs/ui"
)

type FBox struct {
	FBaseView
	v *ui.Box
}

func VBox() *FBox {
	v := &FBox{}
	v.v = ui.NewVerticalBox()
	v.view = v.v
	return v
}
func HBox() *FBox {
	v := &FBox{}
	v.v = ui.NewHorizontalBox()
	v.view = v.v
	return v
}

func (v *FBox) Append(is ...IView) *FBox {
	for _, i := range is {
		v.v.Append(i.getBaseView().view, i.getBaseView().stretch)
	}
	return v
}
