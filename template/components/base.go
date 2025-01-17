package components

import (
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type Base struct {
	Attribute types.Attribute
}

func (b Base) Box() types.BoxAttribute {
	return &BoxAttribute{
		Name:       "box",
		Header:     template.HTML(""),
		Body:       template.HTML(""),
		Footer:     template.HTML(""),
		Title:      "",
		HeadBorder: "",
		Attribute:  b.Attribute,
	}
}

func (b Base) Col() types.ColAttribute {
	return &ColAttribute{
		Name:      "col",
		Size:      "col-md-2",
		Content:   "",
		Attribute: b.Attribute,
	}
}

func (b Base) Form() types.FormAttribute {
	return &FormAttribute{
		Name:      "form",
		Content:   []types.FormField{},
		Url:       "/",
		Method:    "post",
		InfoUrl:   "",
		Title:     "edit",
		Attribute: b.Attribute,
	}
}

func (b Base) Image() types.ImgAttribute {
	return &ImgAttribute{
		Name:      "image",
		Width:     "50",
		Height:    "50",
		Src:       "",
		Attribute: b.Attribute,
	}
}

func (b Base) Tabs() types.TabsAttribute {
	return &TabsAttribute{
		Name:      "tabs",
		Attribute: b.Attribute,
	}
}

func (b Base) Alert() types.AlertAttribute {
	return &AlertAttribute{
		Name:      "alert",
		Attribute: b.Attribute,
	}
}

func (b Base) Label() types.LabelAttribute {
	return &LabelAttribute{
		Name:      "label",
		Color:     "success",
		Content:   "",
		Attribute: b.Attribute,
	}
}

func (b Base) Popup() types.PopupAttribute {
	return &PopupAttribute{
		Name:      "popup",
		Attribute: b.Attribute,
	}
}

func (b Base) Paginator() types.PaginatorAttribute {
	return &PaginatorAttribute{
		Name:      "paginator",
		Attribute: b.Attribute,
	}
}

func (b Base) Row() types.RowAttribute {
	return &RowAttribute{
		Name:      "row",
		Content:   "",
		Attribute: b.Attribute,
	}
}

func (b Base) Button() types.ButtonAttribute {
	return &ButtonAttribute{
		Name:      "button",
		Content:   "",
		Href:      "",
		Attribute: b.Attribute,
	}
}

func (b Base) Table() types.TableAttribute {
	return &TableAttribute{
		Name:      "table",
		Thead:     []map[string]string{},
		InfoList:  []map[string]template.HTML{},
		Type:      "normal",
		Attribute: b.Attribute,
	}
}

func (b Base) DataTable() types.DataTableAttribute {
	return &DataTableAttribute{
		TableAttribute: *(b.Table().SetType("data-table").(*TableAttribute)),
		EditUrl:        "",
		NewUrl:         "",
		Attribute:      b.Attribute,
	}
}

func (b Base) Tree() types.TreeAttribute {
	return &TreeAttribute{
		Name:      "tree",
		Tree:      []menu.Item{},
		Attribute: b.Attribute,
	}
}
