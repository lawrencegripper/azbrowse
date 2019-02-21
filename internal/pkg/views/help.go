package views

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
)

type keyBindings struct {
	help string
}

// ToggleHelpView shows and hides the help view
func ToggleHelpView(g *gocui.Gui) {

	// }
}

// DrawHelp renders the popup help view
func DrawHelp(keyBindings map[string]string, v *gocui.View) {

	for k, v := range keyBindings {
		keyBindings[k] = strings.ToUpper(v)
	}

	fmt.Fprint(v, style.Header(fmt.Sprintf(`
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
                                                                                                                                                                                                           
# Operations                                                                                                                                                                                                             
                                                                                                                                                                                                           
| Action                   | Key(s)                                                                                                                                                                                                                                                             
| -------------------------| --------------------                                                                                                                                                                                                                                                             
| Toggle browse JSON       | %s                                                                                                                                                                                                 
| Toggle fullscreen        | %s                                                                                                                                                                                                 
| Open Azure portal        | %s                                                                                                                                                                                                 
| Delete resource          | %s                                                                                                                                                                                                 
| Save JSON to clipboard   | %s                                                                                                                                                                                                 
| View actions for resource| %s                                                                                                                                                                                                                                                                                                                                                                                                           
                                                                                                                                                                                                           
For bugs, issue or to contribute visit: https://github.com/lawrencegripper/azbrowse                                                                                                                                                                                                                                 
                                                                                                                                                                                                           
# Status Icons                                                                                                                                                                                                           
                                                                                                                                                                                                           
Deleting: ☠ Failed: ⛈  Updating: ⟳  Resuming/Starting:    ⛅  Provisioning: ⌛                                                                                                                                                                                                            
Creating\Preparing: 🏗  Scaling:  ⚖   Suspended/Suspending: ⛔  Succeeded:   🌣                                                                                                                                                                                                                                                             
                                                                                                                                                                                                                                                                                                                                                            
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
		keyBindings["listedit"],
		keyBindings["fullscreen"],
		keyBindings["listopen"],
		keyBindings["delete"],
		keyBindings["save"],
		keyBindings["listactions"],
		keyBindings["help"],
		keyBindings["help"])))
}
