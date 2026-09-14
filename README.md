# fyne-theme-explorer

A desktop app for browsing the current Fyne theme's colors, icons, and sizes.

![GitHub Release](https://img.shields.io/github/v/release/ErikKalkoken/fyne-theme-explorer)
[![Fyne](https://img.shields.io/badge/dynamic/regex?url=https%3A%2F%2Fgithub.com%2FErikKalkoken%2Ffyne-theme-explorer%2Fblob%2Fmain%2Fgo.mod&search=fyne%5C.io%5C%2Ffyne%5C%2Fv2%20(v%5Cd*%5C.%5Cd*%5C.%5Cd*)&replace=%241&label=Fyne&cacheSeconds=https%3A%2F%2Fgithub.com%2Ffyne-io%2Ffyne)](https://github.com/fyne-io/fyne)
[![CI/CD](https://github.com/ErikKalkoken/fyne-theme-explorer/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/ErikKalkoken/fyne-theme-explorer/actions/workflows/ci-cd.yml)
[![codecov](https://codecov.io/gh/ErikKalkoken/fyne-theme-explorer/graph/badge.svg?token=fDk5XvdhOQ)](https://codecov.io/gh/ErikKalkoken/fyne-theme-explorer)
[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/fyne-theme-explorer.svg)](https://pkg.go.dev/github.com/ErikKalkoken/fyne-theme-explorer)
![GitHub License](https://img.shields.io/github/license/ErikKalkoken/fyne-theme-explorer)

## Description

fyne-theme-explorer is a Fyne app for showing details about the default Fyne theme like colors, icons and sizes. This can be very useful when creating your own apps and widgets, e.g. when trying to find the right theme color to use in a new widget.

Features:

- Search, sort, and filter colors, icons, and sizes
- Light / dark / auto theme toggle
- Adjustable icon size and color

All colors, icons and sizes shown are generated directly from the Fyne library, so the list is always complete and matches the version of Fyne currently used by this app.

<img width="914" height="800" alt="Screenshot from 2026-09-14 21-14-55" src="https://github.com/user-attachments/assets/02c36c6c-e739-4065-91ce-5560572dd7bc" />

## Installation

> [!IMPORTANT]
> This app is built with [Fyne](https://fyne.io), which requires a C compiler and a system graphics driver to build and run. See the [Fyne prerequisites](https://docs.fyne.io/started/) for platform-specific setup instructions.

You can run the app directly from the repo with:

```sh
go run github.com/ErikKalkoken/fyne-theme-explorer@latest
```

Or you can install it locally with:

```sh
go install github.com/ErikKalkoken/fyne-theme-explorer@latest
```

Once installed it can be started with:

```sh
fyne-theme-explorer
```
