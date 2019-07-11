package views

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/TylerBrock/colorjson"
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
)

// ItemWidget is response for showing the text response from the Rest requests
type ItemWidget struct {
	x, y      int
	w, h      int
	hideGuids bool
	content   string
	view      *gocui.View
	g         *gocui.Gui
}

// NewItemWidget creates a new instance of ItemWidget
func NewItemWidget(x, y, w, h int, hideGuids bool, content string) *ItemWidget {
	return &ItemWidget{x: x, y: y, w: w, h: h, hideGuids: hideGuids, content: content}
}

// Layout draws the widget in the gocui view
func (w *ItemWidget) Layout(g *gocui.Gui) error {
	w.g = g
	v, err := g.SetView("itemWidget", w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Editable = true
	v.Wrap = true

	w.view = v
	v.Clear()

	if w.content == "" {
		return nil
	}

	if w.hideGuids {
		w.content = stripSecretVals(w.content)
	}

	d := json.NewDecoder(strings.NewReader(w.content))
	d.UseNumber()
	var obj interface{}
	err = d.Decode(&obj)
	if err != nil {
		eventing.SendStatusEvent(eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Failed to display as JSON: " + err.Error(),
			Timeout:    time.Duration(time.Second * 4),
		})
		fmt.Fprint(v, w.content)
		return nil
	}

	f := colorjson.NewFormatter()
	f.Indent = 2
	s, err := f.Marshal(obj)
	if err != nil {
		fmt.Fprint(v, w.content)
	} else {
		fmt.Fprint(v, string(s))
	}

	return nil
}

// SetContent displays the string in the itemview
func (w *ItemWidget) SetContent(content string, title string) {
	w.g.Update(func(g *gocui.Gui) error {
		w.content = content
		// Reset the cursor and origin (scroll poisition)
		// so we don't start at the bottom of a new doc
		w.view.SetCursor(0, 0) //nolint: errcheck
		w.view.SetOrigin(0, 0) //nolint: errcheck
		w.view.Title = title
		return nil
	})
}

// GetContent returns the current content
func (w *ItemWidget) GetContent() string {
	return w.content
}
