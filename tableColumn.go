package faithtop

import (
	"github.com/andlabs/ui"
)

type ITableColumn interface {
	getName() string
	getTableValue(model *ui.TableModel, r int) ui.TableValue
	setTableValue(model *ui.TableModel, r int, v ui.TableValue)
}
