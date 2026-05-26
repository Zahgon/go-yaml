// This source inspired by https://github.com/fatih/color.
package printer

type ColorAttribute int

const (
	ColorReset ColorAttribute = iota
	ColorBold
	ColorFaint
	ColorItalic
	ColorUnderline
	ColorBlinkSlow
	ColorBlinkRapid
	ColorReverseVideo
	ColorConcealed
	ColorCrossedOut
)

const (
	ColorFgHiBlack ColorAttribute = iota + 90
	ColorFgHiRed
	ColorFgHiGreen
	ColorFgHiYellow
	ColorFgHiBlue
	ColorFgHiMagenta
	ColorFgHiCyan
	ColorFgHiWhite
)

const (
	ColorResetBold ColorAttribute = iota + 22
	ColorResetItalic
	ColorResetUnderline
	ColorResetBlinking

	ColorResetReversed
	ColorResetConcealed
	ColorResetCrossedOut
)

const escape = "\x1b"

var colorResetMap = map[ColorAttribute]ColorAttribute{
	ColorBold:         ColorResetBold,
	ColorFaint:        ColorResetBold,
	ColorItalic:       ColorResetItalic,
	ColorUnderline:    ColorResetUnderline,
	ColorBlinkSlow:    ColorResetBlinking,
	ColorBlinkRapid:   ColorResetBlinking,
	ColorReverseVideo: ColorResetReversed,
	ColorConcealed:    ColorResetConcealed,
	ColorCrossedOut:   ColorResetCrossedOut,
}

func format(attrs ...ColorAttribute) string { _ = "STUB: not implemented"; return "" }

func unformat(attrs ...ColorAttribute) string { _ = "STUB: not implemented"; return "" }

func colorize(msg string, attrs ...ColorAttribute) string { _ = "STUB: not implemented"; return "" }
