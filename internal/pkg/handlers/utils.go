package handlers

// DrawStatus converts a status string to an icon
func DrawStatus(s string) string {
	switch s {
	case "Deleting":
		return "☠"
	case "Failed":
		return "⛈"
	case "Updating":
		return "⟳"
	case "Resuming":
		return "⛅"
	case "Starting":
		return "⛅"
	case "Provisioning":
		return "⌛"
	case "Creating":
		return "🏗"
	case "Preparing":
		return "🏗"
	case "Scaling":
		return "⚖"
	case "Suspended":
		return "⛔"
	case "Suspending":
		return "⛔"
	case "Succeeded":
		return "🌣"
	}
	return ""
}
