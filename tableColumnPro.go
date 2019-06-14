package faithtop

import "github.com/andlabs/ui"

// FTableColumnProgress text column
type FTableColumnProgress struct {
	name          string
	onGetProgress func(r int) int
}

func (f *FTableColumnProgress) getName() string {
	return f.name
}

func (f *FTableColumnProgress) getTableValue(model *ui.TableModel, r int) ui.TableValue {
	i := f.onGetProgress(r)
	if i == -1 {
		return ui.TableInt(-1)
	}
	return ui.TableInt(i % 100)
}

func (f *FTableColumnProgress) setTableValue(model *ui.TableModel, r int, v ui.TableValue) {
}

func TableColumnProgress(name string, onGetProgress func(row int) int) *FTableColumnProgress {
	if onGetProgress == nil {
		panic("onGetProgress cannot be nil")
	}
	return &FTableColumnProgress{
		name:          name,
		onGetProgress: onGetProgress,
	}
}
