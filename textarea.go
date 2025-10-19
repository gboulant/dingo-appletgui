package appletgui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TextAreaHandler struct {
	TextWidget *widget.TextGrid
	Container  *container.Scroll
}

func NewTextArea() *TextAreaHandler {
	tw := widget.NewTextGrid()
	ct := container.NewVScroll(tw)
	ta := TextAreaHandler{TextWidget: tw, Container: ct}
	tw.ShowLineNumbers = true
	return &ta
}

func (t *TextAreaHandler) Append(text string) {
	// use fyne.Do so that it can be called in a go routine outside of
	// the main loop. Note also that the TxtGrid Append function first
	// go to the next line and add the text. To avoid a blank line at
	// the beginning, we then check if the widget is void and use
	// SetText instead of Append if void.
	fyne.Do(func() {
		text = strings.TrimSuffix(text, "\n")
		if t.TextWidget.Text() == "" {
			t.TextWidget.SetText(text)
		} else {
			t.TextWidget.Append(text)
		}
		t.TextWidget.Refresh()
		t.Container.ScrollToBottom()
	})
}

func (t *TextAreaHandler) Set(text string) {
	fyne.Do(func() {
		t.TextWidget.SetText(text)
		t.TextWidget.Refresh()
	})
}

func (t *TextAreaHandler) Clear() {
	t.Set("")
}
