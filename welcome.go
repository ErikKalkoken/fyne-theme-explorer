package main

import (
	_ "embed"
	"net/url"

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

func makeWelcome() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Welcome to Fyne Theme Explorer", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title.SizeName = theme.SizeNameSubHeadingText

	version := widget.NewLabelWithStyle(fyne.CurrentApp().Metadata().Version, fyne.TextAlignCenter, fyne.TextStyle{})

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
