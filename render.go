package main

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
)

// TemplRender TemplRender
// @Description:
type TemplRender struct {
	Code int
	Data templ.Component
}

// Render Render
//
//	@Description:
//	@receiver t
//	@param w
//	@return error
func (t *TemplRender) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Code)
	if t.Data != nil {
		return t.Data.Render(context.Background(), w)
	}
	return nil
}

// WriteContentType WriteContentType
//
//	@Description:
//	@receiver t
//	@param w
func (t *TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

// Instance Instance
//
//	@Description:
//	@receiver t
//	@param name
//	@param data
//	@return render.Render
func (t *TemplRender) Instance(name string, data interface{}) render.Render {
	if templData, ok := data.(templ.Component); ok {
		return &TemplRender{
			Code: http.StatusOK,
			Data: templData,
		}
	}
	return nil
}
