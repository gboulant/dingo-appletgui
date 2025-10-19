package appletgui

import applet "github.com/gboulant/dingo-applet"

// AddApplet is a proxy to the core package function (just for
// convenience of code writing when registering the applets)
var AddApplet func(name string, comment string, function func() error) *applet.Applet = applet.AddApplet

// StartApplication is athe standard way to initialize and start the
// graphical application with all applets registered in the applets
// catalog. Nevertheless, you may use directly the base class AppletGui
// (see demos examples).
func StartApplication(title string) error {
	gui, err := NewAppletGui(title)
	if err != nil {
		return err
	}
	gui.Run()
	return nil
}
