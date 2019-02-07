package monitor

// Deepin windowmanager detector
type Deepin struct {
}

func (s *Deepin) Name() string {
	return "Deepin"
}

func (s *Deepin) ExecutablesExists() bool {
	return which("deepin-session") != "" && which("dconf") != ""
}

func (s *Deepin) Running() bool {
	return containsE("GDMSESSION", "deepin") || containsE("XDG_SESSION_DESKTOP", "deepin") || containsE("XDG_CURRENT_DESKTOP", "deepin")
}

// SetWallpaper sets the desktop wallpaper, given an image filename.
// The image must exist and be readable.
func (s *Deepin) SetWallpaper(imageFilename string) error {
	return run("dconf write /com/deepin/wrap/gnome/desktop/background/picture-uri \"'" + imageFilename + "'\"")
}