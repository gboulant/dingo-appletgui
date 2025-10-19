package appletgui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TextAreaWidget struct {
	TextWidget *widget.TextGrid
}

func NewTextArea() *TextAreaWidget {
	ta := TextAreaWidget{TextWidget: widget.NewTextGrid()}
	ta.TextWidget.ShowLineNumbers = true
	return &ta
}

func (t *TextAreaWidget) Append(text string) {
	// use fyne.Do so that it can be called in a go routine outside of
	// the main loop
	fyne.Do(func() {
		t.TextWidget.Append(text)
		t.TextWidget.Refresh()
	})
}

func (t *TextAreaWidget) Set(text string) {
	fyne.Do(func() {
		t.TextWidget.SetText(text)
		t.TextWidget.Refresh()
	})
}

func (t *TextAreaWidget) Clear() {
	t.Set("")
}
