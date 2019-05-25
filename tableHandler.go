package faithtop

import (
	"github.com/andlabs/ui"
)

type tableHandler struct {
	types        []ui.TableValue
	numRows      int
	cellValue    func(row, column int) ui.TableValue
	setCellValue func(row, column int, value ui.TableValue)
}

func (t *tableHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	if t.types == nil {
		return []ui.TableValue{
			ui.TableString(""),
		}
	}
	return t.types
}

func (t *tableHandler) NumRows(m *ui.TableModel) int {
	return t.numRows
}

func (t *tableHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if t.cellValue == nil {
		panic("cellValue")
	}
	return t.cellValue(row, column)
}

func (t *tableHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	if t.setCellValue == nil {
		panic("setCellValue")
	}
	t.setCellValue(row, column, value)
}
