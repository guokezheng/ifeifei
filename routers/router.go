package routers

import (
	"github.com/astaxie/beego"
	"ifeifei/controllers/admin"
	"ifeifei/controllers/blog"
)

func init() {

	//首页
    beego.Router("/", &blog.MainController{},"*:Index")
	beego.Router("/index.html", &blog.MainController{}, "*:Index")

    //软文分享
	beego.Router("/share.html", &blog.MainController{},"*:Share")

	//博客列表
	beego.Router("/list.html", &blog.MainController{},"*:List")

	//归类归档
	beego.Router("/category.html", &blog.MainController{},"*:Category")

    //关于
	beego.Router("/about.html", &blog.MainController{},"*:About")

	//留言
	beego.Router("/message.html", &blog.MainController{},"*:Message")

	//后台管理主页
	beego.Router("/admin", &admin.IndexController{},"*:Index")
	beego.Router("/admin/user.html", &admin.IndexController{},"*:User")
	beego.Router("/admin/user/index", &admin.UserController{},"*:Index")

	beego.Router("/admin/user/list", &admin.UserController{},"*:List")
	beego.Router("/admin/user/edit", &admin.UserController{},"*:Edit")


}
