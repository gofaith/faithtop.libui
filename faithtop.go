package faithtop

import (
	"github.com/andlabs/ui"
)

var (
	idMap = make(map[string]interface{})
)

func MainQuit() {
	ui.Quit()
}
func Main(f func()) {
	ui.Main(f)
}
func RunOnUIThread(f func()) {
	ui.QueueMain(f)
}
