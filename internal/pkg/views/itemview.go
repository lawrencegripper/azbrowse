package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-xmlfmt/xmlfmt"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/stuartleeks/colorjson"
	"github.com/stuartleeks/gocui"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/quick"
	"github.com/alecthomas/chroma/styles"
)

// ItemWidget is response for showing the text response from the Rest requests
type ItemWidget struct {
	x, y            int
	w, h            int
	hideGuids       bool
	node            *expanders.TreeNode
	originalContent string // unformatted - for copying
	content         string
	contentType     interfaces.ExpanderResponseType
	view            *gocui.View
	shouldRender    bool
	g               *gocui.Gui
}

var _ interfaces.ItemWidget = &ItemWidget{}

// NewItemWidget creates a new instance of ItemWidget
func NewItemWidget(x, y, w, h int, hideGuids bool, shouldRender bool, content string) *ItemWidget {
	configureYAMLHighlighting()

	return &ItemWidget{x: x, y: y, w: w, h: h, hideGuids: hideGuids, shouldRender: shouldRender, content: content}
}

// Layout draws the widget in the gocui view
func (w *ItemWidget) Layout(g *gocui.Gui) error {
	w.g = g

	x0, y0, x1, y1 := getViewBounds(g, w.x, w.y, w.w, w.h)

	v, err := g.SetView("itemWidget", x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Editable = true
	v.Wrap = true
	w.view = v
	width, height := v.Size()
	expanders.ItemWidgetHeight = height
	expanders.ItemWidgetWidth = width
	v.Clear()

	if w.shouldRender {
		if w.content == "" {
			return nil
		}

		if w.hideGuids {
			w.content = StripSecretVals(w.content)
		}
		fmt.Fprint(v, w.content)
	}

	return nil
}

// PageDown move the view down a page
func (w *ItemWidget) PageDown() {
	_, maxHeight := w.view.Size()
	x, y := w.view.Origin()
	w.view.SetCursor(0, 0) //nolint: errcheck

	maxY := strings.Count(w.content, "\n")
	y = y + maxHeight
	if y > maxY {
		y = maxY
	}

	err := w.view.SetOrigin(x, y)
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Failed to execute pagedown: " + err.Error(),
			Timeout:    time.Duration(time.Second * 4),
		})
	}
}

// PageUp move the view a page up
func (w *ItemWidget) PageUp() {
	_, maxHeight := w.view.Size()
	x, y := w.view.Origin()
	w.view.SetCursor(0, 0) //nolint: errcheck

	y = y - maxHeight
	// Check we haven't overshot
	if y < 0 {
		y = 0
	}
	err := w.view.SetOrigin(x, y)
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Failed to execute pagedown: " + err.Error(),
			Timeout:    time.Duration(time.Second * 4),
		})
	}
}

// SetContent displays the string in the itemview
func (w *ItemWidget) SetContent(content string, contentType interfaces.ExpanderResponseType, title string) {
	w.SetContentWithNode(nil, content, contentType, title)
}

// SetContentWithNode displays the string in the itemview and tracks the associated node
func (w *ItemWidget) SetContentWithNode(node *expanders.TreeNode, content string, contentType interfaces.ExpanderResponseType, title string) {
	w.g.Update(func(g *gocui.Gui) error {
		w.node = node
		w.originalContent = content
		w.content = content
		w.contentType = contentType

		if w.hideGuids {
			w.content = StripSecretVals(w.content)
		}
		switch w.contentType {
		case interfaces.ResponseJSON:
			d := json.NewDecoder(strings.NewReader(w.content))
			d.UseNumber()
			var obj interface{}
			err := d.Decode(&obj)
			if err != nil {
				eventing.SendStatusEvent(&eventing.StatusEvent{
					InProgress: false,
					Failure:    true,
					Message:    "Failed to display as JSON: " + err.Error(),
					Timeout:    time.Duration(time.Second * 4),
				})
			}

			f := colorjson.NewFormatter()
			f.Indent = 2
			s, err := f.Marshal(obj)
			if err == nil {
				w.content = string(s)
			}
		case interfaces.ResponseYAML:
			var buf bytes.Buffer
			err := quick.Highlight(&buf, w.content, "YAML-azbrowse", "terminal", "azbrowse")
			if err == nil {
				w.content = buf.String()
			}

		case interfaces.ResponseTerraform:
			var buf bytes.Buffer
			err := quick.Highlight(&buf, w.content, "Terraform", "terminal", "azbrowse")
			if err == nil {
				w.content = buf.String()
			}

		case interfaces.ResponseXML:
			formattedContent := strings.TrimSpace(xmlfmt.FormatXML(w.content, "", "  "))
			formattedContent = strings.ReplaceAll(formattedContent, "\r", "")
			w.content = formattedContent
		}

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
	return w.originalContent
}

// GetContentType returns the current content type
func (w *ItemWidget) GetContentType() interfaces.ExpanderResponseType {
	return w.contentType
}

// GetNode returns the TreeNode associated with the currently displayed content (or nil if content is not related to a node)
func (w *ItemWidget) GetNode() *expanders.TreeNode {
	return w.node
}

// SetShouldRender set the shouldRender value of this item
func (w *ItemWidget) SetShouldRender(val bool) {
	w.shouldRender = val
}

func configureYAMLHighlighting() {
	lexer := chroma.MustNewLexer(
		&chroma.Config{
			Name:      "YAML-azbrowse",
			Aliases:   []string{"yaml"},
			Filenames: []string{"*.yaml", "*.yml"},
			MimeTypes: []string{"text/x-yaml"},
		},
		chroma.Rules{
			"root": {
				chroma.Include("whitespace"),
				{Pattern: `#.*`, Type: chroma.Comment, Mutator: nil},
				{Pattern: `!![^\s]+`, Type: chroma.CommentPreproc, Mutator: nil},
				{Pattern: `&[^\s]+`, Type: chroma.CommentPreproc, Mutator: nil},
				{Pattern: `\*[^\s]+`, Type: chroma.CommentPreproc, Mutator: nil},
				{Pattern: `^%include\s+[^\n\r]+`, Type: chroma.CommentPreproc, Mutator: nil},
				{Pattern: `([>|+-]\s+)(\s+)((?:(?:.*?$)(?:[\n\r]*?)?)*)`, Type: chroma.ByGroups(chroma.StringDoc, chroma.StringDoc, chroma.StringDoc), Mutator: nil},
				chroma.Include("value"),
				{Pattern: `[?:,\[\]]`, Type: chroma.Punctuation, Mutator: nil},
				{Pattern: `.`, Type: chroma.Text, Mutator: nil},
			},
			"value": {
				{Pattern: chroma.Words(``, `\b`, "true", "false", "null"), Type: chroma.KeywordConstant, Mutator: nil},
				{Pattern: `"(?:\\.|[^"])*"`, Type: chroma.StringDouble, Mutator: nil},
				{Pattern: `'(?:\\.|[^'])*'`, Type: chroma.StringSingle, Mutator: nil},
				{Pattern: `\d\d\d\d-\d\d-\d\d([T ]\d\d:\d\d:\d\d(\.\d+)?(Z|\s+[-+]\d+)?)?`, Type: chroma.LiteralDate, Mutator: nil},
				{Pattern: `\b[+\-]?(0x[\da-f]+|0o[0-7]+|(\d+\.?\d*|\.?\d+)(e[\+\-]?\d+)?|\.inf|\.nan)\b`, Type: chroma.Number, Mutator: nil},
				{Pattern: `\b([\w]+)([ \t]*)([:]+)([ \t]*)(\d+\.?\d*|\.?\d+)(\s)`, Type: chroma.ByGroups(chroma.Text, chroma.Whitespace, chroma.Punctuation, chroma.Whitespace, chroma.Number, chroma.Whitespace), Mutator: nil},
				{Pattern: `\b([\w]+)([ \t]*)([:]+)([ \t]*)(true)\b`, Type: chroma.ByGroups(chroma.Text, chroma.Whitespace, chroma.Punctuation, chroma.Whitespace, chroma.LiteralStringBoolean), Mutator: nil},
				{Pattern: `\b([\w]+)([ \t]*)([:]+)([ \t]*)(false)\b`, Type: chroma.ByGroups(chroma.Text, chroma.Whitespace, chroma.Punctuation, chroma.Whitespace, chroma.LiteralStringBoolean), Mutator: nil},
				{Pattern: `\b([\w]+)([ \t]*)([:]+)([ \t]*)([\w\./\-_]+)\b`, Type: chroma.ByGroups(chroma.Text, chroma.Whitespace, chroma.Punctuation, chroma.Whitespace, chroma.StringDouble), Mutator: nil},
				// {Pattern: `\b(?:[\w]+)(?:[:\b]+)(?:[\w+])\b`, Type: ByGroups(KeywordReserved, Punctuation, StringDouble), Mutator: nil},
				{Pattern: `\b[\w]+\b`, Type: chroma.Text, Mutator: nil},
			},
			"whitespace": {
				{Pattern: `\s+`, Type: chroma.Whitespace, Mutator: nil},
			},
		},
	)

	lexers.Register(lexer)

	style := chroma.MustNewStyle(
		"azbrowse",
		chroma.StyleEntries{
			chroma.LiteralString:        "#00aa00",
			chroma.LiteralStringBoolean: "#b3760e",
			chroma.LiteralNumber:        "#0099ff",
			chroma.Keyword:              "#0099ff",
			chroma.Comment:              "#006600",
		},
	)
	styles.Register(style)
}

// SetHideGuids sets the HideGuids option
func (w *ItemWidget) SetHideGuids(value bool) {
	w.hideGuids = value
}
