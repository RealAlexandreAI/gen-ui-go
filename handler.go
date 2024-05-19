package main

import (
	"github.com/RealAlexandreAI/gen-ui-go/component"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// loadHandlers
//
//	@Description:
//	@param r
func loadHandlers(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		metaTags := component.MetaTags(
			"gen-ui-go, htmx, templ",
			"Generative UI in Golang.",
		)

		indexTemplate := component.Layout(
			"Generative UI in Golang.",
			metaTags,
		)

		_ = indexTemplate.Render(c, c.Writer)

	})

	componentRouter := r.Group("component")

	loadExamples(componentRouter)
}

// loadExamples
//
//	@Description:
//	@param router
func loadExamples(router *gin.RouterGroup) {
	component.ExamplePool.Each(func(key string, value templ.Component) {
		router.GET(key, func(c *gin.Context) {
			_ = component.Wrap(value).Render(c, c.Writer)
		})
	})
}
