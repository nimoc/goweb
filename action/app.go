package action

import (
	"fmt"
	"github.com/og/goweb/core"
	"github.com/og/go-json"
	"log"
)

type routerNameDict struct {
	Home string
}
func RouterNameDict() routerNameDict {
	return routerNameDict {
		Home: "home",
	}
}
func Start(app core.App) {
	app.Render(core.RenderConfig{
		Dir: "./public",
	})
	app.Use(func(c core.Content) core.ResponseError {
		log.Print(fmt.Sprintf("Request %s:%s", c.Req.Method, c.Req.FullURL))
		return c.Next()
	})
	app.Get("/",
		core.RouterConfig{Name:RouterNameDict().Home},
		func(c core.Content) core.ResponseError {
			query := struct {
				Name string `json:"name"`
				Age int `json:"age"`
			}{}
			c.BindQuery(&query)
			c.Cookie().Set("name", query.Name)
			c.Session().Set("id", query.Name)
			view := struct {
				Name string `json:"name"`
			}{}
			view.Name = query.Name
			c.SetJSON(view)
			return c.Next()
		})
	app.Post("/",core.RouterConfig{}, func(c core.Content) core.ResponseError {
		query := struct {
			Title string `json:"title"`
		}{}
		c.BindForm(query)
		c.SetString("you post title: " + query.Title)
		return c.Next()
	})
	app.Get("/end",
		core.RouterConfig{},
		func(c core.Content) core.ResponseError {
			c.EndJSON(gjson.FailMsg("阻止"))
			return c.Next()
		})
	app.Use(func(c core.Content) core.ResponseError {
		log.Print(fmt.Sprintf("Response %s:%s %s", c.Req.Method, c.Req.FullURL, c.Res.Status))
		return c.Next()
	})
}