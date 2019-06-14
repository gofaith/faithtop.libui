package faithtop

import "github.com/andlabs/ui"

// FTableColumnEdit text column
type FTableColumnEdit struct {
	name      string
	data      map[int]string
	onGetText func(r int) string
	onSetText func(r int, text string)
}

func (f *FTableColumnEdit) getName() string {
	return f.name
}

func (f *FTableColumnEdit) getTableValue(model *ui.TableModel, r int) ui.TableValue {
	if v, ok := f.data[r]; ok {
		return ui.TableString(v)
	}
	return ui.TableString(f.onGetText(r))
}

func (f *FTableColumnEdit) setTableValue(model *ui.TableModel, r int, v ui.TableValue) {
	str := string(v.(ui.TableString))
	f.data[r] = str
	model.RowChanged(r)
	f.onSetText(r, str)
}

func TableColumnEdit(name string, onGetText func(row int) string, onSetText func(row int, text string)) *FTableColumnEdit {
	if onGetText == nil {
		panic("onGetText cannot be nil")
	}
	return &FTableColumnEdit{
		name:      name,
		onGetText: onGetText,
		onSetText: onSetText,
		data:      make(map[int]string),
	}
}
