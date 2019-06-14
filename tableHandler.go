package faithtop

import (
	"github.com/andlabs/ui"
)

type tableHandler struct {
	fnOnGetRows      func() int
	columnTypes []ui.TableValue
	fnCellValue        func(model *ui.TableModel,row, column int) ui.TableValue
	fnSetCellValue     func(model *ui.TableModel,row, column int, value ui.TableValue)
}

func (t *tableHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return t.columnTypes
}

func (t *tableHandler) NumRows(m *ui.TableModel) int {
	return t.fnOnGetRows()
}

func (t *tableHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	return t.fnCellValue(m,row, column)
}

func (t *tableHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	t.fnSetCellValue(m,row, column, value)
}
