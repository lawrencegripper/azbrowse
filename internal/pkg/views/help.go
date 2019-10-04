package views

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/stuartleeks/gocui"
)

// DrawHelp renders the popup help view
func DrawHelp(keyBindings map[string][]string, v *gocui.View) {

	for k, v := range keyBindings {
		for i, v2 := range v {
			keyBindings[k][i] = strings.ToUpper(v2)
		}
	}

	view := fmt.Sprintf(`
--> PRESS %s TO CLOSE THIS AND CONTINUE. YOU CAN OPEN IT AGAIN WITH %s AT ANY TIME. <--
                             _       ___
                            /_\   __| _ )_ _ _____ __ _____ ___
                           / _ \ |_ / _ \ '_/ _ \ V  V (_-</ -_)
                          /_/ \_\/__|___/_| \___/\_/\_//__/\___|
                        Interactive CLI for browsing Azure resources
 
# Navigation
 
| Action                   | Key(s)
| -------------------------| --------------------
| Select resource          | %s / %s
| Select menu/JSON         | %s / %s
| Go back                  | %s
| Expand/View resource     | %s
| Refresh                  | %s
| Show this help page      | %s
| Quit                     | %s
 
# Operations
 
| Action                   | Key(s)
| -------------------------| --------------------
| Toggle browse JSON       | %s
| Toggle fullscreen        | %s
| Open Azure portal        | %s
| Delete resource          | %s
| Save JSON to clipboard   | %s
| View actions for resource| %s
| Edit Resource            | %s
 
For bugs, issue or to contribute visit: https://github.com/lawrencegripper/azbrowse
 
# Status Icons
 
Deleting:  ☠   Failed:  ⛈   Updating:  ⟳   Resuming/Starting:  ⛅   Provisioning:  ⌛                                                                                                                                  
Creating\Preparing:  🏗   Scaling:  ⚖   Suspended/Suspending:  ⛔   Succeeded:  🌣                                                                                                                                        
 
--> PRESS %s TO CLOSE THIS AND CONTINUE. YOU CAN OPEN IT AGAIN WITH %s AT ANY TIME. <--

`, keyBindings["help"],
		keyBindings["help"],
		keyBindings["listup"],
		keyBindings["listdown"],
		keyBindings["itemleft"],
		keyBindings["listright"],
		keyBindings["listback"],
		keyBindings["listexpand"],
		keyBindings["listrefresh"],
		keyBindings["help"],
		keyBindings["quit"],
		keyBindings["listedit"],
		keyBindings["fullscreen"],
		keyBindings["listopen"],
		keyBindings["listdelete"],
		keyBindings["copy"],
		keyBindings["listactions"],
		keyBindings["listupdate"],
		keyBindings["help"],
		keyBindings["help"])

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
