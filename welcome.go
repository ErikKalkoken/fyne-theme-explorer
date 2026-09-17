package main

import (
	_ "embed"
	"net/url"
	"runtime/debug"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//go:embed Icon.png
var iconBytes []byte

func parseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	return u
}

// appVersion returns the app version from the Fyne metadata (FyneApp.toml or
// packaged build), falling back to the module version embedded by
// `go install pkg@version` when no Fyne metadata was loaded.
func appVersion() string {
	if meta := fyne.CurrentApp().Metadata(); meta.ID != "" {
		return meta.Version
	}
	if bi, ok := debug.ReadBuildInfo(); ok {
		return bi.Main.Version
	}
	return "unknown"
}

func makeWelcome() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Welcome to Fyne Theme Explorer", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title.SizeName = theme.SizeNameSubHeadingText

	version := widget.NewLabelWithStyle(appVersion(), fyne.TextAlignCenter, fyne.TextStyle{})

	logo := canvas.NewImageFromResource(fyne.NewStaticResource("Icon.png", iconBytes))
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(96, 96))

	description := widget.NewLabelWithStyle(
		"A desktop app for browsing the current Fyne theme's colors, icons, and sizes.\n\nSelect a tab from the list on the left to get started.",
		fyne.TextAlignCenter,
		fyne.TextStyle{},
	)

	footer := container.NewHBox(
		layout.NewSpacer(),
		widget.NewHyperlink("GitHub", parseURL(repoURL)),
		widget.NewLabel("-"),
		widget.NewHyperlink("Fyne Documentation", parseURL("https://docs.fyne.io/")),
		layout.NewSpacer(),
	)

	content := container.NewCenter(container.NewVBox(
		title,
		version,
		logo,
		description,
	))

	return container.NewBorder(nil, footer, nil, nil, content)
}
