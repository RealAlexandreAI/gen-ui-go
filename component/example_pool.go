package component

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/emirpasic/gods/v2/maps/linkedhashmap"
)

// nolint
var ExamplePool = linkedhashmap.New[string, templ.Component]()

// RegisterExample
//
//	@Description:
//	@param path
//	@param example
func RegisterExample(path string, example templ.Component) {

	if example == nil {
		return
	}

	if _, ok := ExamplePool.Get(path); ok {
		panic("duplicated path: " + path)
	}

	ExamplePool.Put(path, example)
}

// GetExample
//
//	@Description:
//	@param path
//	@return templ.Component
//	@return error
func GetExample(path string) (templ.Component, error) {

	example, ok := ExamplePool.Get(path)
	if !ok {
		return nil, fmt.Errorf("not found path: %s", path)
	}

	return example, nil
}
