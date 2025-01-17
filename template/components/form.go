package components

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type FormAttribute struct {
	Name            string
	Header          template.HTML
	Content         []types.FormField
	TabContents     [][]types.FormField
	TabHeaders      []string
	Footer          template.HTML
	Url             string
	Method          string
	PrimaryKey      string
	InfoUrl         string
	CSRFToken       string
	Title           template.HTML
	OperationFooter template.HTML
	Prefix          string
	CdnUrl          string
	types.Attribute
}

func (compo *FormAttribute) SetHeader(value template.HTML) types.FormAttribute {
	compo.Header = value
	return compo
}

func (compo *FormAttribute) SetPrimaryKey(value string) types.FormAttribute {
	compo.PrimaryKey = value
	return compo
}

func (compo *FormAttribute) SetContent(value []types.FormField) types.FormAttribute {
	compo.Content = value
	return compo
}

func (compo *FormAttribute) SetTabContents(value [][]types.FormField) types.FormAttribute {
	compo.TabContents = value
	return compo
}

func (compo *FormAttribute) SetTabHeaders(value []string) types.FormAttribute {
	compo.TabHeaders = value
	return compo
}

func (compo *FormAttribute) SetFooter(value template.HTML) types.FormAttribute {
	compo.Footer = value
	return compo
}

func (compo *FormAttribute) SetPrefix(value string) types.FormAttribute {
	compo.Prefix = value
	return compo
}

func (compo *FormAttribute) SetUrl(value string) types.FormAttribute {
	compo.Url = value
	return compo
}

func (compo *FormAttribute) SetInfoUrl(value string) types.FormAttribute {
	compo.InfoUrl = value
	return compo
}

func (compo *FormAttribute) SetMethod(value string) types.FormAttribute {
	compo.Method = value
	return compo
}

func (compo *FormAttribute) SetTitle(value template.HTML) types.FormAttribute {
	compo.Title = value
	return compo
}

func (compo *FormAttribute) SetToken(value string) types.FormAttribute {
	compo.CSRFToken = value
	return compo
}

func (compo *FormAttribute) GetBoxHeader() template.HTML {
	return template.HTML(fmt.Sprintf(`<h3 class="box-title">%s</h3>
            <div class="box-tools">
                <div class="btn-group pull-right" style="margin-right: 10px">
                    <a href='%s' class="btn btn-sm btn-default form-history-back"><i
                                class="fa fa-arrow-left"></i> %s</a>
                </div>
            </div>`, language.GetFromHtml(compo.Title), compo.InfoUrl, language.Get("Back")))
}

func (compo *FormAttribute) SetOperationFooter(value template.HTML) types.FormAttribute {
	compo.OperationFooter = value
	return compo
}

func (compo *FormAttribute) GetContent() template.HTML {
	compo.CdnUrl = config.Get().AssetUrl
	return ComposeHtml(compo.TemplateList, *compo, "form",
		"form/default", "form/file", "form/textarea", "form/custom",
		"form/selectbox", "form/text", "form/radio", "form/switch",
		"form/password", "form/select", "form/singleselect",
		"form/richtext", "form/iconpicker", "form/datetime", "form/number",
		"form/email", "form/url", "form/ip", "form/color", "form/currency", "form_components")
}
