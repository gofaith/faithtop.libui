package faithtop

import (
	"github.com/andlabs/ui"
)

type FListView struct {
	FBaseView
	v               *ui.Area
	lb              *ui.Box
	vhs             []ViewHolder
	currentCreation int
	createView      func(*FListView) IView
	bindData        func(*ViewHolder, int)
	getCount        func() int
}

type ViewHolder struct {
	vlist map[string]IView
}



func VlistView(createView func(*FListView) IView, bindData func(*ViewHolder, int), getCount func() int) *FListView {
	fb := &FListView{}
	v:=ui.NewScrollingArea(nil, 40, 40)
	setupWidget(&v.Widget)
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	fb.createView = createView
	fb.bindData = bindData
	fb.getCount = getCount
	fb.v.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	fb.lb, _ = gtk.ListBoxNew()
	fb.v.Add(fb.lb)
	for i := 0; i < getCount(); i++ {
		fb.currentCreation = i
		fb.vhs = append(fb.vhs, ViewHolder{vlist: make(map[string]IView)})
		fb.lb.Add(createView(fb).getBaseView().view)
	}
	fb.execBindData()
	return fb.Size(-2, -2)
}

func GetListViewById(id string) *FListView {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FListView); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetListViewByItemId(id string) *FListView {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FListView); ok {
			return bt
		}
	}
	return nil
}