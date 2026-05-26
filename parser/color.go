package parser

const (
	colorFgHiBlack int = iota + 90
	colorFgHiRed
	colorFgHiGreen
	colorFgHiYellow
	colorFgHiBlue
	colorFgHiMagenta
	colorFgHiCyan
)

var colorTable = []int{
	colorFgHiRed,
	colorFgHiGreen,
	colorFgHiYellow,
	colorFgHiBlue,
	colorFgHiMagenta,
	colorFgHiCyan,
}

func colorize(idx int, content string) string { _ = "STUB: not implemented"; return "" }
