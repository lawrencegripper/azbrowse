package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/go-xmlfmt/xmlfmt"
	"github.com/lawrencegripper/azbrowse/internal/pkg/eventing"
	"github.com/lawrencegripper/azbrowse/internal/pkg/expanders"
	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/stuartleeks/colorjson"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/quick"
	"github.com/alecthomas/chroma/styles"
)

// ItemWidget is response for showing the text response from the Rest requests
type ItemWidget struct {
	x, y                 int
	w, h                 int
	hideGuids            bool
	node                 *expanders.TreeNode
	originalContent      string // unformatted - for copying
	content              string
	unfilteredContent    string // formatted content, before filter applied
	contentType          interfaces.ExpanderResponseType
	view                 *gocui.View
	shouldRender         bool
	FullscreenKeyBinding string
	ActionKeyBinding     string
	g                    *gocui.Gui
	filterHandler        func(s string) error
	filterString         string
	title                string
	// track if we need to re-render the layout or is it the same content?
	hasChanged bool
}

var _ interfaces.ItemWidget = &ItemWidget{}

// NewItemWidget creates a new instance of ItemWidget
func NewItemWidget(x, y, w, h int, hideGuids bool, shouldRender bool, content string, filterHandler func(s string) error) *ItemWidget {
	configureYAMLHighlighting()

	return &ItemWidget{
		x: x, y: y, w: w, h: h,
		hideGuids:     hideGuids,
		shouldRender:  shouldRender,
		content:       content,
		filterHandler: filterHandler,
	}
}

// Layout draws the widget in the gocui view
func (w *ItemWidget) Layout(g *gocui.Gui) error {
	w.g = g

	x0, y0, x1, y1 := getViewBounds(g, w.x, w.y, w.w, w.h)

	v, err := g.SetView("itemWidget", x0, y0, x1, y1, 0)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Editable = true
	v.KeybindOnEdit = true
	v.Wrap = true
	v.Subtitle = fmt.Sprintf("[%s-> Fullscreen|%s -> Actions|Type to filter]", strings.ToUpper(w.FullscreenKeyBinding), strings.ToUpper(w.ActionKeyBinding))
	v.Editor = gocui.EditorFunc(func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
		switch key {
		case gocui.KeySpace:
			v.EditWrite(' ')
			return
		case gocui.KeyArrowDown:
			v.MoveCursor(0, 1)
			return
		case gocui.KeyArrowUp:
			v.MoveCursor(0, -1)
			return
		case gocui.KeyArrowLeft:
			v.MoveCursor(-1, 0)
			return
		case gocui.KeyArrowRight:
			v.MoveCursor(1, 0)
			return
		}

		// If we type any chars other than / (which pops up filtering) start filtering
		if ch != 0 && ch != rune('/') {
			w.filterHandler(string(ch))
		}
	})
	w.view = v
	width, height := v.Size()
	expanders.ItemWidgetHeight = height
	expanders.ItemWidgetWidth = width

	// only update the content if we should render
	// and there has been a change to the content
	// SetContentWithNode sets hasChanged = true when
	// called we reset it to false after doing a layout
	if w.shouldRender && w.hasChanged {
		w.hasChanged = false
		if w.content == "" {
			return nil
		}

		if w.hideGuids {
			w.content = StripSecretVals(w.content)
		}
		v.Clear()
		fmt.Fprint(v, w.content)
	}

	if w.filterString != "" {
		v.Title = w.title + " [filter=" + w.filterString + "]"
	}

	return nil
}

// PageDown move the view down a page
func (w *ItemWidget) PageDown() {
	_, maxHeight := w.view.Size()
	x, y := w.view.Origin()

	maxY := strings.Count(w.content, "\n")
	y = y + maxHeight
	if y > maxY {
		y = maxY
	}

	w.view.SetOrigin(x, y) //nolint: errcheck
	err := w.view.SetCursor(x, y)
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

	y = y - maxHeight
	// Check we haven't overshot
	if y < 0 {
		y = 0
	}

	w.view.SetOrigin(x, y) //nolint: errcheck
	err := w.view.SetCursor(x, y)
	if err != nil {
		eventing.SendStatusEvent(&eventing.StatusEvent{
			InProgress: false,
			Failure:    true,
			Message:    "Failed to execute pagedown: " + err.Error(),
			Timeout:    time.Duration(time.Second * 4),
		})
	}
}

// SetFilter sets the filter to be applied to list items
func (w *ItemWidget) SetFilter(filterString string) {
	w.filterString = filterString

	var currentContent []string
	if w.unfilteredContent == "" {
		currentContent = strings.Split(w.content, "\n")
		w.unfilteredContent = w.content
	} else {
		currentContent = strings.Split(w.unfilteredContent, "\n")
	}

	var filteredResult strings.Builder
	// ensure the content isn't under the filter resultbox
	filteredResult.WriteString("\n\n")

	for _, line := range currentContent {
		if strings.Contains(strings.ToLower(line), strings.ToLower(filterString)) {
			line = highlightText(line, filterString)
			filteredResult.WriteString(line + "\n")
		}
	}

	w.content = filteredResult.String()
	w.hasChanged = true

	w.g.Update(func(gui *gocui.Gui) error {
		return nil
	})
}

// ClearFilter clears a filter if applied
func (w *ItemWidget) ClearFilter() {
	w.filterString = ""
	w.content = w.unfilteredContent
	w.hasChanged = true

	w.g.Update(func(gui *gocui.Gui) error {
		return nil
	})
}

// SetContent displays the string in the itemview
func (w *ItemWidget) SetContent(content string, contentType interfaces.ExpanderResponseType, title string) {
	w.SetContentWithNode(nil, content, contentType, title)
}

// SetContentWithNode displays the string in the itemview and tracks the associated node
func (w *ItemWidget) SetContentWithNode(node *expanders.TreeNode, content string, contentType interfaces.ExpanderResponseType, title string) {
	// Clear down filter stuff
	w.filterString = ""
	w.unfilteredContent = ""

	w.view.Clear()
	w.hasChanged = true
	w.node = node
	w.originalContent = content
	w.content = content
	w.contentType = contentType
	w.title = title

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

	// Update the view
	w.g.Update(func(*gocui.Gui) error { return nil })
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
