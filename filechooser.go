package faithtop

import (
	"github.com/andlabs/ui"
)

func OpenFile(w *FWin) string {
	return ui.OpenFile(w.v)
}

func SaveFile(w *FWin) string {
	return ui.SaveFile(w.v)
}
