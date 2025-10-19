package appletgui

import (
	"embed"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	applet "github.com/gboulant/dingo-applet"
	stdrw "github.com/gboulant/dingo-stdrw"
)

//go:embed icons/*.png
var icons embed.FS

type AppletGui struct {
	Title      string
	app        fyne.App
	win        fyne.Window
	ctnActions *fyne.Container
	TextArea   *TextAreaHandler
	stdOutHdl  *stdrw.StdoutHandler
}

func NewAppletGui(title string) (*AppletGui, error) {
	gui := AppletGui{Title: title}
	if err := gui.Setup(); err != nil {
		return nil, err
	}

	names := applet.GetAppletNames()
	for _, name := range names {
		a, _ := applet.GetApplet(name)
		gui.AddApplet(a)
	}

	return &gui, nil
}

func (g *AppletGui) Setup() error {
	a := app.New()
	w := a.NewWindow(g.Title)

	// --------------------------------------
	// Text Area (center of the BorderLayout)
	textArea := NewTextArea()
	ctnCenter := textArea.Container
	// Note that the TextAreaHandler is not a widget and can not be
	// displayed. It is composed of a scrollable Container containing a
	// TextWidget, both can be handled.

	// --------------------------------------
	// Action Area (on the left)
	btnQuit := widget.NewButton("Quit", func() {
		log.Print("On ferme toutes les fenêtres ... ")
		w.Close()
		log.Println("done")
	})
	icondata, _ := icons.ReadFile("icons/quit.png")
	btnQuit.Icon = fyne.NewStaticResource("quit", icondata)

	// Define a customizable actions container embedded in a fixed
	// container that will be placed in the left boder of the window
	// border layout
	ctnActions := container.NewVBox()
	ctnLeft := container.NewVBox()
	ctnLeft.Add(ctnActions)
	ctnLeft.Add(layout.NewSpacer())
	ctnLeft.Add(btnQuit)

	// --------------------------------------
	// Packing all together inn a border layout
	w.SetContent(container.NewBorder(nil, nil, ctnLeft, nil, ctnCenter))
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	w.SetMaster()

	// --------------------------------------
	// Setup the redirection from stdout toward the text area.
	h, err := stdrw.NewStdoutHandler(func(line string) {
		g.TextArea.Append(line)
	})
	if err != nil {
		return err
	}
	// At this step the redirection is not activated yet. It will
	// activated when running the application

	// --------------------------------------
	// Keep references in the struct for further usage
	g.app = a
	g.win = w
	g.ctnActions = ctnActions
	g.TextArea = textArea
	g.stdOutHdl = h
	return nil
}

func (g *AppletGui) AddAction(label string, action func()) {
	btn := widget.NewButton(label, action)
	btn.Alignment = widget.ButtonAlignLeading
	g.ctnActions.Add(btn)
}

func (g *AppletGui) AddApplet(a *applet.Applet) {
	label := fmt.Sprintf("%s - %s", a.Name, a.Comment)
	action := func() {
		g.TextArea.Clear()
		go func() {
			if err := a.Execute(); err != nil {
				s := fmt.Sprintf("err: %s\n", err)
				g.TextArea.Append(s)
			}
		}()
	}
	g.AddAction(label, action)
}

func (g *AppletGui) Run() {
	g.stdOutHdl.Start()
	defer g.stdOutHdl.Stop()
	g.win.Show()
	g.app.Run()
}
