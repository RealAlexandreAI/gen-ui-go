// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package component

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BodyContent() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container\"><div class=\"left-side\"><button class=\"component-button\" hx-get=\"/component/button\" hx-target=\"#display-area\" hx-swap=\"beforeend scroll:bottom show:bottom\">Button</button> <button class=\"component-button\" hx-get=\"/component/list\" hx-target=\"#display-area\" hx-swap=\"beforeend scroll:bottom show:bottom\">List</button> <button class=\"component-button\" hx-get=\"/component/textarea\" hx-target=\"#display-area\" hx-swap=\"beforeend scroll:bottom show:bottom\">TextArea</button></div><div class=\"right-side\"><div id=\"display-area\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Style() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n        body, html {\n            margin: 0;\n            padding: 0;\n            height: 100%;\n            overflow: hidden;\n            font-family: sans-serif;\n        }\n        .container {\n            display: flex;\n            justify-content: space-between;\n            padding: 10px;\n            padding-bottom: 100px;\n        }\n        .left-side {\n            background-color: #f2f2f2;\n            padding: 10px;\n            flex: 1;\n        }\n        .component-button {\n            width: 100%;\n            padding: 10px;\n            margin-bottom: 10px;\n            background-color: #4CAF50;\n            color: white;\n            text-align: center;\n            border: none;\n            cursor: pointer;\n            outline: none;\n        }\n        .right-side {\n            background-color: #e0e0e0;\n            flex: 3;\n            overflow-y: auto;\n            padding: 10px;\n            position: relative;\n            min-height: 0;\n        }\n        #display-area {\n            display: flex;\n            flex-direction: column;\n            align-items: flex-end;\n            position: relative;\n            padding-bottom: 20px;\n            width: 96%;\n            bottom: 0;\n        }\n\n        #display-area > hr {\n            border: 0;\n            height: 1px;\n            padding-bottom: 10px;\n            background-color: #ccc;\n        }\n\n        #display-area > div.generated-component {\n            background-color: white;\n            padding: 15px;\n            border-radius: 5px;\n            margin-bottom: 10px;\n        }\n\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
