package faithtop

import (
	"fmt"

	"github.com/andlabs/ui"
)

// FTableColumnCheck check box column
type FTableColumnCheck struct {
	name         string
	data         map[int]bool
	onGetChecked func(r int) bool
	onSetChecked func(r int, b bool)
}

func (f *FTableColumnCheck) getName() string {
	return f.name
}

func (f *FTableColumnCheck) getTableValue(model *ui.TableModel, r int) ui.TableValue {
	if v, ok := f.data[r]; ok {
		return tranformBool(v)
	}
	return tranformBool(f.onGetChecked(r))
}

func (f *FTableColumnCheck) setTableValue(model *ui.TableModel, r int, value ui.TableValue) {
	fmt.Println(value.(ui.TableInt))
	b := value.(ui.TableInt) == ui.TableTrue
	f.data[r] = b
	model.RowChanged(r)
	f.onSetChecked(r, b)
}

func tranformBool(b bool) ui.TableValue {
	if b {
		return ui.TableTrue
	}
	return ui.TableFalse
}

func TableColumnCheck(name string, onGetCheck func(row int) bool, onSetChecked func(row int, b bool)) *FTableColumnCheck {
	if onGetCheck == nil || onSetChecked == nil {
		panic("nil")
	}
	return &FTableColumnCheck{
		name:         name,
		onGetChecked: onGetCheck,
		onSetChecked: onSetChecked,
		data:         make(map[int]bool),
	}
}
