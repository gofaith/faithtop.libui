package faithtop

import "github.com/andlabs/ui"

// FTableColumnButton button column
type FTableColumnButton struct {
	name      string
	onGetText func(r int) string
	onClick   func(r int)
}

func (f *FTableColumnButton) getName() string {
	return f.name
}
func (f *FTableColumnButton) setTableValue(model *ui.TableModel, r int, v ui.TableValue) {
	f.onClick(r)
}

func (f *FTableColumnButton) getTableValue(model *ui.TableModel, r int) ui.TableValue {
	return ui.TableString(f.onGetText(r))
}

// TableColumnButton TableColumnButton
func TableColumnButton(name string, onGetText func(row int) string, onClick func(row int)) *FTableColumnButton {
	if onGetText == nil || onClick == nil {
		panic("onGetText or onClick cannot be nil")
	}
	return &FTableColumnButton{
		name:      name,
		onGetText: onGetText,
		onClick:   onClick,
	}
}
