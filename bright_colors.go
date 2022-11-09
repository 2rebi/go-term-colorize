package termcolor

import "fmt"

type BrightColorType string

const (
	BrightColorTypeForeground BrightColorType = "9"
	BrightColorTypeBackground BrightColorType = "10"
)

type BrightColorCode string

func (code BrightColorCode) Colorize(typ BrightColorType, str string) string {
	return fmt.Sprintf("\x1b[%sm%s", fmt.Sprintf(string(code), typ), str)
}

func (code BrightColorCode) Foreground(str string) string {
	return code.Colorize(BrightColorTypeForeground, str)
}

func (code BrightColorCode) Background(str string) string {
	return code.Colorize(BrightColorTypeBackground, str)
}

const (
	// BrightBlack color hex = #555753
	BrightBlack BrightColorCode = "%s0"
	// BrightRed color hex = #ef2929
	BrightRed BrightColorCode = "%s1"
	// BrightGreen color hex = #8ae234
	BrightGreen BrightColorCode = "%s2"
	// BrightYellow color hex = #fce94f
	BrightYellow BrightColorCode = "%s3"
	// BrightBlue color hex = #32afff
	BrightBlue BrightColorCode = "%s4"
	// BrightMagenta color hex = #ad7fa8
	BrightMagenta BrightColorCode = "%s5"
	// BrightCyan color hex = #34e2e2
	BrightCyan BrightColorCode = "%s6"
	// BrightWhite color hex = #ffffff
	BrightWhite BrightColorCode = "%s7"
)
