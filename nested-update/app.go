// +build wasm

// The UI is running only on a web browser. Therefore, the build instruction
// above is to compile the code below only when the program is built for the
// WebAssembly (wasm) architecture.

package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo

	Name string
}

type nested struct {
	app.Compo

	NestedName string
}

func (n *nested) Render() app.UI {
	return app.H3().Body(app.Text("nested: " + n.NestedName))
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World!"),
		app.H2().Body(app.Text("top-level: "+h.Name)),
		&nested{NestedName: h.Name},
		app.Input().
			Value(h.Name).
			OnChange(h.OnInputChange),
		app.Text("Type something and press enter"),
	)
}

func (h *hello) OnInputChange(ctx app.Context, e app.Event) {
	h.Name = ctx.JSSrc.Get("value").String()
	h.Update()
}

// The main function is the entry point of the UI. It is where components are
// associated with URL paths and where the UI is started.
func main() {
	app.Route("/", &hello{
		Name: "start value",
	})        // hello component is associated with URL path "/".
	app.Run() // Launches the PWA.
}
