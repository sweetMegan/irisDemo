package main

import (
	"github.com/kataras/iris"
	"fmt"
)

func main() {
	//获取一个iris对象
	app := iris.New()
	//getAPI(app)
	//postAPI(app)
	postAPI_2(app)

	//在8080端口启动服务
	app.Run(iris.Addr(":8080"))
}
func postAPI_2(app *iris.Application)  {
	//绑定路由/login/{用户名}/{密码}，访问后输出用户名，密码
	app.Handle("POST", "/regist", func(ctx iris.Context) {
		data := &struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		//读取JSON数据
		ctx.ReadJSON(data)
		fmt.Println(data.Username,"==",data.Password)
		ctx.JSON(data)
	})
}
func postAPI(app *iris.Application)  {
	//绑定路由/login/{用户名}/{密码}，访问后输出用户名，密码
	app.Handle("POST", "/regist", func(ctx iris.Context) {
		//获取参数值
		username := ctx.FormValue("username")
		fmt.Println("==",username)
		password := ctx.FormValue("password")
		fmt.Println(password)
		data := &struct {
			Username string
			Password string
		}{
			Username: username,
			Password: password,
		}
		ctx.JSON(data)
	})
}
func getAPI(app *iris.Application)  {
	//绑定路由/login/{用户名}/{密码}，访问后输出用户名，密码
	app.Handle("GET", "/login/{username}/{password}", func(ctx iris.Context) {
		//获取参数值
		username := ctx.Params().Get("username")
		fmt.Println(username)
		password := ctx.Params().Get("password")
		fmt.Println(password)
		data := &struct {
			Username string
			Password string
		}{
			Username: username,
			Password: password,
		}
		ctx.JSON(data)
	})
}