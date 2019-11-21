package config

// Settings to enable different behavior (specified at runtime)
type Settings struct {
	EnableTracing bool
	HideGuids     bool
	NavigateToID  string
}
