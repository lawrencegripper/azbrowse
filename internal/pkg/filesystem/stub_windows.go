// +build windows

package filesystem

// Run starts a fuze based filesystem over the Azure ARM API
func Run(mountpoint string, filterToSub string, editMode bool, demoMode bool) (func(), error) {
	panic("Fuse not supported on windows")
}
