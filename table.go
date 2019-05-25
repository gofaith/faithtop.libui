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
}

func Table(onGetRows func() int, columns ...*FTableColumn) *FTable {
	v := &FTable{}
	v.handler = &tableHandler{}
	v.model = ui.NewTableModel(v.handler)
	v.v = ui.NewTable(&ui.TableParams{
		Model:                         v.model,
		RowBackgroundColorModelColumn: 1,
	})
	v.view = v.v
	return v
}

func GetTableById(id string) *FTable {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FTable); ok {
			return b
		}
	}
	return nil
}
