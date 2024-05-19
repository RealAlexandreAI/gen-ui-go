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
			"gowebly, htmx example page, go with htmx",
			"Welcome to example! You're here because it worked out.",
		)

		indexTemplate := component.Layout(
			"Welcome to example!",
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
