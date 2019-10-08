package views

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
)

const tmplText = `
--> PRESS {{ index . "help" }} TO CLOSE THIS AND CONTINUE. YOU CAN OPEN IT AGAIN WITH {{ index . "help" }} AT ANY TIME. <--
                             _       ___
                            /_\   __| _ )_ _ _____ __ _____ ___
                           / _ \ |_ / _ \ '_/ _ \ V  V (_-</ -_)
                          /_/ \_\/__|___/_| \___/\_/\_//__/\___|
                        Interactive CLI for browsing Azure resources
# Navigation
 
| Action                   | Key(s)
| -------------------------| --------------------
| Select resource          | {{ index . "listup" }} / {{ index . "listdown" }}
| Select menu/JSON         | {{ index . "itemleft" }} / {{ index . "listright" }}
| Go back                  | {{ index . "listback" }}
| Expand/View resource     | {{ index . "listexpand" }}
| Refresh                  | {{ index . "listrefresh" }}
| Filter                   | {{ index . "filter" }}
| Clear filter             | {{ index . "listclearfilter" }}
| Open Command Panel       | {{ index . "commandpanelopen" }}
| Close Command Panel      | {{ index . "commandpanelclose" }}
| Show this help page      | {{ index . "help" }}
| Quit                     | {{ index . "quit" }}
 
# Operations
 
| Action                   | Key(s)
| -------------------------| --------------------
| Toggle browse JSON       | {{ index . "listedit" }}
| Toggle fullscreen        | {{ index . "fullscreen" }}
| Open Azure portal        | {{ index . "listopen" }}
| Delete resource          | {{ index . "listdelete" }}
| Save JSON to clipboard   | {{ index . "copy" }}
| View actions for resource| {{ index . "listactions" }}
| Edit Resource            | {{ index . "listupdate" }}
 
# Status Icons

Deleting:  ☠   Failed:  ⛈   Updating:  ⟳   Resuming/Starting:  ⛅   Provisioning:  ⌛                                                                                                                                  
Creating\Preparing:  🏗   Scaling:  ⚖   Suspended/Suspending:  ⛔   Succeeded:  🌣                                                                                                                                        

For bugs, issue or to contribute visit: https://github.com/lawrencegripper/azbrowse
`

// DrawHelp renders the popup help view
func DrawHelp(keyBindings map[string][]string, v *gocui.View) {

	for k, v := range keyBindings {
		for i, v2 := range v {
			keyBindings[k][i] = strings.ToUpper(v2)
		}
	}

	tmpl, err := template.New("help").Parse(tmplText)
	if err != nil {
		panic("Failed to parse help template. This is a Bug please raise an issue on GH. " + err.Error())
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, keyBindings)
	if err != nil {
		panic("Failed to execute help template. This is a Bug please raise an issue on GH. " + err.Error())
	}

	view := buf.String()

	maxWidth, maxHeight := v.Size()

	paddedView := ""

	lineCount := 0
	scanner := bufio.NewScanner(strings.NewReader(view))
	for scanner.Scan() {
		line := scanner.Text()
		rightPadLen := maxWidth - len(line)
		pad := ""
		if rightPadLen > 0 {
			pad = strings.Repeat(" ", rightPadLen)
		}
		paddedView = fmt.Sprintf("%s%s%s\n", paddedView, line, pad)
		lineCount++
	}

	bottomPadLen := maxHeight - lineCount
	for i := 0; i < bottomPadLen; i++ {
		pad := strings.Repeat(" ", maxWidth)
		paddedView = fmt.Sprintf("%s%s\n", paddedView, pad)
	}

	fmt.Fprint(v, style.Header(paddedView))
}
