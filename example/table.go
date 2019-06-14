package main

import (
	"fmt"

	. "gitee.com/gofaith/faithtop.libui"
)

func main() {
	Main(setup)
}
func setup() {
	table := Table(onGetRows,
		TableColumnText("text", onGetText),
	)
	table.Expand()
	WinDefault().DeferShow().QuitOnClose().VBox(
		table,
	)
}
func onGetRows() int {
	return 28
}
func onGetText(r int) string {
	return fmt.Sprint(r)
}
