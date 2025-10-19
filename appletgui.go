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

// AddApplet is a proxy to the core package function
var AddApplet func(name string, comment string, function func() error) *applet.Applet = applet.AddApplet

type AppletGui struct {
	Title      string
	app        fyne.App
	win        fyne.Window
	ctnActions *fyne.Container
	textArea   *widget.TextGrid
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
	// Text Area (on the right)
	ctnRight := container.NewVBox()
	textArea := widget.NewTextGrid()
	//ctnScroll := container.NewScroll(textArea)
	//ctnRight.Add(ctnScroll)
	ctnRight.Add(textArea)

	// --------------------------------------
	// Action Area (on the left)
	btnQuit := widget.NewButton("Quit", func() {
		fmt.Print("On ferme toutes les fenêtres ... ")
		w.Close()
		fmt.Println("done")
	})

	ctnLeft := container.NewVBox()
	ctnActions := container.NewVBox()
	ctnLeft.Add(ctnActions)
	ctnLeft.Add(layout.NewSpacer())
	ctnLeft.Add(btnQuit)

	// --------------------------------------
	// Packing all together
	w.SetContent(container.NewHBox(ctnLeft, ctnRight))

	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	w.SetMaster()

	// --------------------------------------
	// Setup the redirection from stdout toward the text area.
	h, err := stdrw.NewStdoutHandler(func(line string) {
		g.TextAppend(line)
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
	g.textArea = textArea
	g.stdOutHdl = h
	return nil
}

func (g *AppletGui) TextAppend(text string) {
	// use fyne.Do so that it can be called in a go routine outside of
	// the main loop
	fyne.Do(func() {
		g.textArea.Append(text)
		g.textArea.Refresh()
	})
}

func (g *AppletGui) AddAction(label string, action func()) {
	btn := widget.NewButton(label, action)
	g.ctnActions.Add(btn)
}

func (g *AppletGui) AddApplet(a *applet.Applet) {
	label := fmt.Sprintf("%s - %s", a.Name, a.Comment)
	action := func() {
		g.textArea.SetText("")
		go func() {
			if err := a.Execute(); err != nil {
				s := fmt.Sprintf("err: %s\n", err)
				fyne.Do(func() {
					g.textArea.Append(s)
				})
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

// StartApplication is a standard way to initialize and start the
// graphical application with all applets registered in the applets
// catalog. Nevertheless, you may use directly the base class AppletGui.
func StartApplication(title string) error {
	gui, err := NewAppletGui(title)
	if err != nil {
		return err
	}
	gui.Run()
	return nil
}
