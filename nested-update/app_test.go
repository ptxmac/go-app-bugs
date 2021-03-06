package main

import (
	"runtime"
	"testing"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/stretchr/testify/require"
)

func testSkipNonWasm(t *testing.T) {
	if goarch := runtime.GOARCH; goarch != "wasm" {
		t.Skip()
	}
}

func TestNested_Render(t *testing.T) {
	testSkipNonWasm(t)

	h := &hello{
		Name: "start",
	}
	t.Run("initial state", func(t *testing.T) {
		require.NoError(t, app.TestMatch(
			h,
			app.TestUIDescriptor{
				Path:     app.TestPath(0, 1, 0),
				Expected: app.Text("top-level: start"),
			},
		))
		require.NoError(t, app.TestMatch(
			h,
			app.TestUIDescriptor{
				Path:     app.TestPath(0, 2, 0, 0),
				Expected: app.Text("nested: start"),
			},
		))
	})
	h.Name = "end"
	h.Update()
	t.Run("updated state", func(t *testing.T) {
		require.NoError(t, app.TestMatch(
			h,
			app.TestUIDescriptor{
				Path:     app.TestPath(0, 1, 0),
				Expected: app.Text("top-level: end"),
			},
		))
		require.NoError(t, app.TestMatch(
			h,
			app.TestUIDescriptor{
				Path:     app.TestPath(0, 2, 0, 0),
				Expected: app.Text("nested: end"),
			},
		))
	})
}
