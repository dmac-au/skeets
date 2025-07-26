package main

import (
	"os"

	_ "embed"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
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
	_ = adw.StyleManagerGetDefault()
	header := adw.NewHeaderBar()
	header.ShowTitle()
	header.PackStart(gtk.NewButtonFromIconName("open-menu-symbolic"))
	titlebar := adw.NewToolbarView()
	titlebar.AddTopBar(header)
	titlebar.SetBottomBarStyle(adw.ToolbarRaised)
	window := adw.NewApplicationWindow(app)
	window.SetTitle("Skeets")
	window.SetContent(titlebar)
	window.SetVisible(true)
}
