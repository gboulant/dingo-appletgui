# appletgui - a graphical user interface for applets

**contact**: [Guillaume Boulant](mailto:gboulant@gmail.com?subject=dingo-appletgui)

The appletgui package provides a graphical user interface for executing
a set of different little applications, called applets. The concept of
applet is defined in the
[applet](https://github.com/gboulant/dingo-applet) package. It consists
in a standard go function without argument and returning an
error:

```go
func DEMO00_logscale() error {
    fmt.Println("Executing demo DEMO00_logscale")
    // Do something
    // ...
    return nil
}
```

The core applet package provides a simple command line interface for
running the applets. This gui package can be used to create a graphical
user interface to play the applets.

Like for the core applet package, you can register the applet functions
with an identifier name identifier, short description and a function to
execute. This can be done eather with the core applet package or the
appletgui one (that contain a proxy to to the NewExample applet
function):

```go
import appgui "github.com/gboulant/dingo-appletgui"
import applet "github.com/gboulant/dingo-applet"

appgui.NewExample("D00", "echelle logarithmique", DEMO00_logscale)
appgui.NewExample("D01", "son de quintes", DEMO01_quintes)
applet.NewExample("D02", "exemple 02", DEMO02_hello)
...
```

And finaly, the main function should execute the `StartExampleApp`, with
the main window title as argument :

```go
appgui.StartExampleApp("Exemple d'utilisation du son")
```

This function starts a graphical user interface that let the user play
the different registered applets:

![appletgui](demos/guiapp/guiapp.png)
