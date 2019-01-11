package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/style"
)

var showHelp bool

// ToggleHelpView shows and hides the help view
func ToggleHelpView(g *gocui.Gui) {
	showHelp = !showHelp

	// If we're up and running clear and redraw the view
	// if w.g != nil {
	if showHelp {
		g.Update(func(g *gocui.Gui) error {
			maxX, maxY := g.Size()
			// Padding
			maxX = maxX - 2
			maxY = maxY - 2
			v, err := g.SetView("helpWidget", 1, 1, 140, 32)
			if err != nil && err != gocui.ErrUnknownView {
				panic(err)
			}
			DrawHelp(v)
			return nil
		})
	} else {
		g.DeleteView("helpWidget")
	}
	// }
}

// DrawHelp renders the popup help view
func DrawHelp(v *gocui.View) {

	fmt.Fprint(v, style.Header(`
	--> PRESS CTRL+H TO CLOSE THIS AND CONTINUE. YOU CAN OPEN IT AGAIN WITH CRTL+H AT ANY TIME. <--
                             _       ___
                            /_\   __| _ )_ _ _____ __ _____ ___
                           / _ \ |_ / _ \ '_/ _ \ V  V (_-</ -_)
                          /_/ \_\/__|___/_| \___/\_/\_//__/\___|
                        Interactive CLI for browsing Azure resources
# Navigation

| Key       | Does                 |
| --------- | -------------------- |
| ⇧ / ⇩     | Select resource      |
| ⇦ / ⇨     | Select Menu/JSON     |
| Backspace | Go back              |
| ENTER     | Expand/View resource |
| F5        | Refresh              |
| CTRL+H    | Show this page       |

# Operations

| Key                 | Does                      |                                                                                    |
| ------------------- | ------------------------- | ---------------------------------------------------------------------------------- |
| CTRL+E              | Toggle Browse JSON        | For longer responses you can move the cursor to scroll the doc                     |
| CTLT+F              | Toggle Fullscreen         | Gives a fullscreen view of the JSON for smaller terminals                          |
| CTRL+O (o for open) | Open Portal               | Opens the portal at the currently selected resource                                |
| DEL                 | Delete resource           | The currently selected resource will be deleted (Requires double press to confirm) |
| CTLT+S              | Save JSON to clipboard    | Saves the last JSON response to the clipboard for export                           |
| CTLT+A              | View Actions for resource | This allows things like ListKeys on storage or Restart on VMs                      |

For bugs, issue or to contribute visit: https://github.com/lawrencegripper/azbrowse

--> PRESS CTRL+H TO CLOSE THIS AND CONTINUE. YOU CAN OPEN IT AGAIN WITH CRTL+H AT ANY TIME. <--
`))

}
