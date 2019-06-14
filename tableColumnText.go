package faithtop

import "github.com/andlabs/ui"

// FTableColumnText text column
type FTableColumnText struct {
	name      string
	onGetText func(r int) string
}

func (f *FTableColumnText) getName() string {
	return f.name
}

func (f *FTableColumnText) getTableValue(model *ui.TableModel, r int) ui.TableValue {
	return ui.TableString(f.onGetText(r))
}

func (f *FTableColumnText) setTableValue(model *ui.TableModel, r int, v ui.TableValue) {
}

func TableColumnText(name string, onGetText func(row int) string) *FTableColumnText {
	if onGetText == nil {
		panic("onGetText cannot be nil")
	}
	return &FTableColumnText{
		name:      name,
		onGetText: onGetText,
	}
}
