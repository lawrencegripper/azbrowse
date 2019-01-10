package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/style"
)

// ToggleHelpView shows and hides the help view
func ToggleHelpView(g *gocui.Gui) {

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
                                                                                                                                                                                                                                                
| Key         | Does                 |                                                                                                                                                                                                                                                 
| ----------- | -------------------- |                                                                                                                                                                                                            
| ↑/↓         | Select resource      |                                                                                                                                                                                                           
| Backspace/← | Go back              |                                                                                                                                                                                                           
| ENTER/→     | Expand/View resource |                                                                                                                                                                                                           
| F5          | Refresh              |                                                                                                                                                                                                           
| CTRL+H      | Show this page       |                                                                                                                                                                                                           
                                                                                                                                                                                                           
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
