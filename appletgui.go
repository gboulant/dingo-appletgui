package appletgui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	applet "github.com/gboulant/dingo-applet"
	stdrw "github.com/gboulant/dingo-stdrw"
)

var NewExample func(name string, comment string, function func() error) *applet.Example = applet.NewExample

func StartExampleApp(title string) {
	a := app.New()
	w := a.NewWindow(title)

	reader, _ := stdrw.NewStdoutReader()
	defer reader.Close()

	ctnRight := container.NewVBox()
	text := widget.NewTextGrid()
	ctnRight.Add(text)

	names := applet.GetExampleNames()
	btnDemos := make([]*widget.Button, len(names))
	for i, name := range names {
		example, _ := applet.GetExample(name)
		btnlabel := fmt.Sprintf("%s - %s", example.Name, example.Comment)
		btnDemos[i] = widget.NewButton(btnlabel, func() {
			text.SetText("")
			go func() {
				if err := example.Execute(); err != nil {
					s := fmt.Sprintf("err: %s\n", err)
					fyne.Do(func() {
						text.Append(s)
					})
				}
			}()
		})
		btnDemos[i].Alignment = widget.ButtonAlignLeading
	}
	btnQuit := widget.NewButton("Quit", func() {
		fmt.Print("On ferme toutes les fenêtres ... ")
		w.Close()
		fmt.Println("done")
	})

	ctnLeft := container.NewVBox()
	for _, btn := range btnDemos {
		ctnLeft.Add(btn)
	}
	ctnLeft.Add(layout.NewSpacer())
	ctnLeft.Add(btnQuit)

	go func() {
		for {
			line, _ := reader.ReadString('\n')
			fyne.Do(func() {
				text.Append(line)
				text.Refresh()
			})
		}
	}()

	w.SetContent(container.NewHBox(ctnLeft, ctnRight))
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	w.SetMaster()
	w.Show()
	a.Run()
}
