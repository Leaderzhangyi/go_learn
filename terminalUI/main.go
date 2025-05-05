package main

import (
	"github.com/rivo/tview"
)

func main() {
	form := tview.NewForm()
	form.AddInputField("Name", "", 20, nil, nil)
	form.AddInputField("Email", "", 20, nil, nil)

	app := tview.NewApplication().
		SetRoot(form, true).
		EnableMouse(true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
