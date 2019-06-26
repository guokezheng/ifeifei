package blog

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
	options           map[string]string
	right             string
}

func (this *MainController) Prepare() {
	this.right = "right.html"
}

func (this *MainController) Index() {
	this.display("index")
}

func (this *MainController) Category() {
	this.right = ""
	this.display("category")
}

func (this *MainController) List() {
	this.right = ""
	this.display("list")
}

func (this *MainController) Message() {
	this.right = ""
	this.display("message")
}

func (this *MainController) About() {
	this.right = ""
	this.display("about")
}

func (this *MainController) Share() {
	this.right = ""
	this.display("share")
}

func (this *MainController) display(tpl string) {
	theme := "m1"
	if v, ok := this.options["theme"]; ok && v != "" {
		theme = v
	}

	this.Layout = theme + "/layout.html"
	this.Data["root"] = "/" + beego.BConfig.WebConfig.ViewsPath + "/" + theme + "/"
	this.TplName = theme + "/" + tpl + ".html"

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["header"] = theme + "/header.html"

	if tpl == "index" {
		this.LayoutSections["banner"] = theme + "/banner.html"
	}
	if this.right != "" {
		this.LayoutSections["right"] = theme +  "/right.html"
	}
	this.LayoutSections["footer"] = theme + "/footer.html"
}


