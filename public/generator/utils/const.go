package utils

// Const for creating Esc-sequence
const (
	Esc = "\u001b"
	CSI = Esc + "["
)

// Preset for background and foreground(text)
// material design colors.
const (
	/* text colors */
	FgRed    = "38;2;255;0;17"
	FgGreen  = "38;2;0;190;82"
	FgYellow = "38;2;251;250;34"
	FgBlue   = "38;2;28;131;246"
	FgPurple = "38;2;174;0;176"
	FgPink   = "38;2;251;0;89"
	FgCyan   = "38;2;;0;190;215"
	FgOrange = "38;2;251;98;00"
	FgBrown  = "38;2;126;83;70"
	FgGray   = "38;2;158;158;158"
	FgWhite  = "38;2;255;255;255"
	FgBlack  = "38;2;0;0;0"

	FgReset = "39" // sets foreground color to default

	/* background colors */
	BgRed    = "48;2;255;0;17"
	BgGreen  = "48;2;0;190;82"
	BgYellow = "48;2;251;250;34"
	BgBlue   = "48;2;28;131;246"
	BgPurple = "48;2;174;0;176"
	BgPink   = "48;2;251;0;89"
	BgCyan   = "48;2;;0;190;215"
	BgOrange = "48;2;251;98;00"
	BgBrown  = "48;2;126;83;70"
	BgGray   = "48;2;158;158;158"
	BgWhite  = "48;2;255;255;255"
	BgBlack  = "48;2;0;0;0"

	BgReset = "49" // sets background color to default
)
