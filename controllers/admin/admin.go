package admin

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
	options           map[string]string
	right             string
}

func (this *IndexController) Prepare() {
	this.right = "right.html"
}

func (this *IndexController) Index() {
	this.display("index")
}

func (this *IndexController) User() {
	this.display("user")
}




func (this *IndexController) display(tpl string) {
	theme := "admin"
	if v, ok := this.options["theme"]; ok && v != "" {
		theme = v
	}

	this.Layout = theme + "/layout.html"
	this.Data["root"] = "/" + beego.BConfig.WebConfig.ViewsPath + "/" + theme + "/"
	this.TplName = theme + "/" + tpl + ".html"

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = theme + "/header.html"

	this.LayoutSections["menu"] = theme + "/menu.html"

	this.LayoutSections["footer"] = theme + "/footer.html"
}


