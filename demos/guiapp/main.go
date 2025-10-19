package main

import (
	"fmt"
	"log"

	applet "github.com/gboulant/dingo-appletgui"
)

const title string = "Exemple d'utilisation du son"

func init() {
	applet.AddApplet("D00", "echelle logarithmique", DEMO00_logscale)
	applet.AddApplet("D01", "son de quintes", DEMO01_quintes)
	applet.AddApplet("D02", "vibrato", DEMO02_vibrato)
	applet.AddApplet("D03", "modulation d'amplitude", DEMO03_amplitude_modulation)
	applet.AddApplet("D04", "modulation de fréquence", DEMO04_frequency_modulation)
	applet.AddApplet("D05", "sounds like a laser", DEMO05_sounds_like_a_laser)
	applet.AddApplet("D06", "echelle musicale", DEMO06_musicalscale)
}

// demo01 shows the standard way to run the applet graphical interface
func demo01() error {
	return applet.StartApplication(title)
}

// demo02 shows a finer way to run the applet graphical interface, in
// particular to customize the actions. In this example, we add an
// action in the buttons list of action.
func demo02() error {
	gui, err := applet.NewAppletGui(title)
	if err != nil {
		return err
	}

	gui.AddAction("Action 1", func() {
		fmt.Println("Action 1")
		gui.TextAppend("Hello action 1")
	})

	gui.Run()
	fmt.Printf("the application terminates")
	return nil
}

func main() {
	//demo := demo01
	demo := demo02
	if err := demo(); err != nil {
		log.Fatal(err)
	}

}
