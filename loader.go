package main

import (
	"github.com/RealAlexandreAI/gen-ui-go/component"
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

	loadComponents(componentRouter)
}

// loadComponents
//
//	@Description:
//	@param router
func loadComponents(router *gin.RouterGroup) {

	router.GET("/button", func(c *gin.Context) {
		comp := component.Button("Click Me!")
		_ = component.Wrap(comp).Render(c, c.Writer)
	})

	router.GET("/list", func(c *gin.Context) {
		comp := component.List([]string{
			"ListItem1",
			"ListItem2",
			"ListItem3",
			"ListItem4",
		})
		_ = component.Wrap(comp).Render(c, c.Writer)
	})

	router.GET("/textarea", func(c *gin.Context) {
		comp := component.Textarea(
			"TextArea label here.",
			`"The quick brown fox jumps over the lazy dog" is an English-language pangram — a sentence that contains all the letters of the alphabet. The phrase is commonly used for touch-typing practice, testing typewriters and computer keyboards, displaying examples of fonts, and other applications involving text where the use of all letters in the alphabet is desired.`)
		_ = component.Wrap(comp).Render(c, c.Writer)
	})

}
