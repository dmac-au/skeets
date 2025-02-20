package main

import (
	"os"

	_ "embed"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

//go:embed ui/main.ui
var mainUI string


func main() {
	app := gtk.NewApplication("com.github.dmac-au.skeets", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	builder := gtk.NewBuilderFromString(mainUI)

	window := builder.GetObject("MainWindow").Cast().(*gtk.Window)

	app.AddWindow(window)
	window.Show()
}
