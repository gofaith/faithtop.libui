package faithtop

import (
	"github.com/andlabs/ui"
)

type FTable struct {
	FBaseView
	v         *ui.Table
	handler   *tableHandler
	model     *ui.TableModel
	onGetRows func() int
	columns   []ITableColumn
	rowsNum   int
}

func Table(onGetRows func() int, columns ...ITableColumn) *FTable {
	v := &FTable{}
	v.handler = &tableHandler{}
	v.handler.fnOnGetRows = onGetRows
	v.handler.fnCellValue = func(model *ui.TableModel, r, c int) ui.TableValue {
		return columns[c].getTableValue(model, r)
	}
	v.handler.fnSetCellValue = func(model *ui.TableModel, r, c int, value ui.TableValue) {
		columns[c].setTableValue(model, r, value)
	}

	v.model = ui.NewTableModel(v.handler)
	v.v = ui.NewTable(&ui.TableParams{
		Model: v.model,
	})
	v.view = v.v
	// table
	if onGetRows == nil {
		panic("onGetRows cannot be nil")
	}
	v.onGetRows = onGetRows
	for index, c := range columns {
		if c == nil {
			panic("table column cannot be nil")
		}
		v.columns = append(v.columns, c)
		if _, ok := c.(*FTableColumnText); ok {
			v.handler.columnTypes = append(v.handler.columnTypes, ui.TableString(""))
			v.v.AppendTextColumn(c.getName(), index, ui.TableModelColumnNeverEditable, nil)
		} else if _, ok := c.(*FTableColumnButton); ok {
			v.handler.columnTypes = append(v.handler.columnTypes, ui.TableString(""))
			v.v.AppendButtonColumn(c.getName(), index, ui.TableModelColumnAlwaysEditable)
		} else if _, ok := c.(*FTableColumnEdit); ok {
			v.handler.columnTypes = append(v.handler.columnTypes, ui.TableString(""))
			v.v.AppendTextColumn(c.getName(), index, ui.TableModelColumnAlwaysEditable, nil)
		} else if _, ok := c.(*FTableColumnCheck); ok {
			v.handler.columnTypes = append(v.handler.columnTypes, ui.TableFalse)
			v.v.AppendCheckboxColumn(c.getName(), index, ui.TableModelColumnNeverEditable)
		} else if _, ok := c.(*FTableColumnProgress); ok {
			v.handler.columnTypes = append(v.handler.columnTypes, ui.TableInt(0))
			v.v.AppendProgressBarColumn(c.getName(), index)
		} else {
			panic("column type invalid")
		}
	}

	//initialize value
	v.rowsNum = onGetRows()
	return v
}

func (f *FTable) SetID(id string) *FTable {
	idMap[id] = f
	return f
}
func GetTableById(id string) *FTable {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FTable); ok {
			return b
		}
	}
	return nil
}

func (f *FTable) NotifyChanged() *FTable {
	newNum := f.onGetRows()
	if f.rowsNum > newNum {
		for i := 0; i < newNum; i++ {
			f.model.RowChanged(i)
		}
		for i := 0; i < f.rowsNum-newNum; i++ {
			f.model.RowDeleted(f.rowsNum - 1 - i)
		}
	} else {
		for i := 0; i < f.rowsNum; i++ {
			f.model.RowChanged(i)
		}
		for i := 0; i < newNum-f.rowsNum; i++ {
			f.model.RowInserted(f.rowsNum + i)
		}
	}
	f.rowsNum = newNum
	return f
}
