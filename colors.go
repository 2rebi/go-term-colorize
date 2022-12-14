package termcolor

import "fmt"

type ColorType rune

const (
	ColorTypeForeground ColorType = '3'
	ColorTypeBackground ColorType = '4'
)

type ColorCode string

func (code ColorCode) Colorize(typ ColorType, str string) string {
	return fmt.Sprintf("\x1b[%sm%s", fmt.Sprintf(string(code), typ), str)
}

func (code ColorCode) Foreground(str string) string {
	return code.Colorize(ColorTypeForeground, str)
}

func (code ColorCode) Background(str string) string {
	return code.Colorize(ColorTypeBackground, str)
}

const (

	// StdBlack color hex = #000000
	StdBlack ColorCode = "%c0"
	// StdRed color hex = #cc0000
	StdRed ColorCode = "%c1"
	// StdGreen color hex = #4e9a06
	StdGreen ColorCode = "%c2"
	// StdYellow color hex = #c4a000
	StdYellow ColorCode = "%c3"
	// StdBlue color hex = #729fcf
	StdBlue ColorCode = "%c4"
	// StdMagenta color hex = #75507b
	StdMagenta ColorCode = "%c5"
	// StdCyan color hex = #06989a
	StdCyan ColorCode = "%c6"
	// StdWhite color hex = #d3d7cf
	StdWhite ColorCode = "%c7"

	// Black color hex = #000000
	Black ColorCode = "%c8;5;0"
	// Maroon color hex = #800000
	Maroon ColorCode = "%c8;5;1"
	// Green color hex = #008000
	Green ColorCode = "%c8;5;2"
	// Olive color hex = #808000
	Olive ColorCode = "%c8;5;3"
	// Navy color hex = #000080
	Navy ColorCode = "%c8;5;4"
	// PurpleType1 color hex = #800080
	PurpleType1 ColorCode = "%c8;5;5"
	// Teal color hex = #008080
	Teal ColorCode = "%c8;5;6"
	// Silver color hex = #c0c0c0
	Silver ColorCode = "%c8;5;7"
	// Grey color hex = #808080
	Grey ColorCode = "%c8;5;8"
	// Red color hex = #ff0000
	Red ColorCode = "%c8;5;9"
	// Lime color hex = #00ff00
	Lime ColorCode = "%c8;5;10"
	// Yellow color hex = #ffff00
	Yellow ColorCode = "%c8;5;11"
	// Blue color hex = #0000ff
	Blue ColorCode = "%c8;5;12"
	// Fuchsia color hex = #ff00ff
	Fuchsia ColorCode = "%c8;5;13"
	// Aqua color hex = #00ffff
	Aqua ColorCode = "%c8;5;14"
	// White color hex = #ffffff
	White ColorCode = "%c8;5;15"
	// Grey0 color hex = #000000
	Grey0 ColorCode = "%c8;5;16"
	// NavyBlue color hex = #00005f
	NavyBlue ColorCode = "%c8;5;17"
	// DarkBlue color hex = #000087
	DarkBlue ColorCode = "%c8;5;18"
	// Blue3Type1 color hex = #0000af
	Blue3Type1 ColorCode = "%c8;5;19"
	// Blue3Type2 color hex = #0000d7
	Blue3Type2 ColorCode = "%c8;5;20"
	// Blue1 color hex = #0000ff
	Blue1 ColorCode = "%c8;5;21"
	// DarkGreen color hex = #005f00
	DarkGreen ColorCode = "%c8;5;22"
	// DeepSkyBlue4Type1 color hex = #005f5f
	DeepSkyBlue4Type1 ColorCode = "%c8;5;23"
	// DeepSkyBlue4Type2 color hex = #005f87
	DeepSkyBlue4Type2 ColorCode = "%c8;5;24"
	// DeepSkyBlue4Type3 color hex = #005faf
	DeepSkyBlue4Type3 ColorCode = "%c8;5;25"
	// DodgerBlue3 color hex = #005fd7
	DodgerBlue3 ColorCode = "%c8;5;26"
	// DodgerBlue2 color hex = #005fff
	DodgerBlue2 ColorCode = "%c8;5;27"
	// Green4 color hex = #008700
	Green4 ColorCode = "%c8;5;28"
	// SpringGreen4 color hex = #00875f
	SpringGreen4 ColorCode = "%c8;5;29"
	// Turquoise4 color hex = #008787
	Turquoise4 ColorCode = "%c8;5;30"
	// DeepSkyBlue3Type1 color hex = #0087af
	DeepSkyBlue3Type1 ColorCode = "%c8;5;31"
	// DeepSkyBlue3Type2 color hex = #0087d7
	DeepSkyBlue3Type2 ColorCode = "%c8;5;32"
	// DodgerBlue1 color hex = #0087ff
	DodgerBlue1 ColorCode = "%c8;5;33"
	// Green3Type1 color hex = #00af00
	Green3Type1 ColorCode = "%c8;5;34"
	// SpringGreen3Type1 color hex = #00af5f
	SpringGreen3Type1 ColorCode = "%c8;5;35"
	// DarkCyan color hex = #00af87
	DarkCyan ColorCode = "%c8;5;36"
	// LightSeaGreen color hex = #00afaf
	LightSeaGreen ColorCode = "%c8;5;37"
	// DeepSkyBlue2 color hex = #00afd7
	DeepSkyBlue2 ColorCode = "%c8;5;38"
	// DeepSkyBlue1 color hex = #00afff
	DeepSkyBlue1 ColorCode = "%c8;5;39"
	// Green3Type2 color hex = #00d700
	Green3Type2 ColorCode = "%c8;5;40"
	// SpringGreen3Type2 color hex = #00d75f
	SpringGreen3Type2 ColorCode = "%c8;5;41"
	// SpringGreen2Type1 color hex = #00d787
	SpringGreen2Type1 ColorCode = "%c8;5;42"
	// Cyan3 color hex = #00d7af
	Cyan3 ColorCode = "%c8;5;43"
	// DarkTurquoise color hex = #00d7d7
	DarkTurquoise ColorCode = "%c8;5;44"
	// Turquoise2 color hex = #00d7ff
	Turquoise2 ColorCode = "%c8;5;45"
	// Green1 color hex = #00ff00
	Green1 ColorCode = "%c8;5;46"
	// SpringGreen2Type2 color hex = #00ff5f
	SpringGreen2Type2 ColorCode = "%c8;5;47"
	// SpringGreen1 color hex = #00ff87
	SpringGreen1 ColorCode = "%c8;5;48"
	// MediumSpringGreen color hex = #00ffaf
	MediumSpringGreen ColorCode = "%c8;5;49"
	// Cyan2 color hex = #00ffd7
	Cyan2 ColorCode = "%c8;5;50"
	// Cyan1 color hex = #00ffff
	Cyan1 ColorCode = "%c8;5;51"
	// DarkRedType1 color hex = #5f0000
	DarkRedType1 ColorCode = "%c8;5;52"
	// DeepPink4Type1 color hex = #5f005f
	DeepPink4Type1 ColorCode = "%c8;5;53"
	// Purple4Type1 color hex = #5f0087
	Purple4Type1 ColorCode = "%c8;5;54"
	// Purple4Type2 color hex = #5f00af
	Purple4Type2 ColorCode = "%c8;5;55"
	// Purple3 color hex = #5f00d7
	Purple3 ColorCode = "%c8;5;56"
	// BlueViolet color hex = #5f00ff
	BlueViolet ColorCode = "%c8;5;57"
	// Orange4Type1 color hex = #5f5f00
	Orange4Type1 ColorCode = "%c8;5;58"
	// Grey37 color hex = #5f5f5f
	Grey37 ColorCode = "%c8;5;59"
	// MediumPurple4 color hex = #5f5f87
	MediumPurple4 ColorCode = "%c8;5;60"
	// SlateBlue3Type1 color hex = #5f5faf
	SlateBlue3Type1 ColorCode = "%c8;5;61"
	// SlateBlue3Type2 color hex = #5f5fd7
	SlateBlue3Type2 ColorCode = "%c8;5;62"
	// RoyalBlue1 color hex = #5f5fff
	RoyalBlue1 ColorCode = "%c8;5;63"
	// Chartreuse4 color hex = #5f8700
	Chartreuse4 ColorCode = "%c8;5;64"
	// DarkSeaGreen4Type1 color hex = #5f875f
	DarkSeaGreen4Type1 ColorCode = "%c8;5;65"
	// PaleTurquoise4 color hex = #5f8787
	PaleTurquoise4 ColorCode = "%c8;5;66"
	// SteelBlue color hex = #5f87af
	SteelBlue ColorCode = "%c8;5;67"
	// SteelBlue3 color hex = #5f87d7
	SteelBlue3 ColorCode = "%c8;5;68"
	// CornflowerBlue color hex = #5f87ff
	CornflowerBlue ColorCode = "%c8;5;69"
	// Chartreuse3Type1 color hex = #5faf00
	Chartreuse3Type1 ColorCode = "%c8;5;70"
	// DarkSeaGreen4Type2 color hex = #5faf5f
	DarkSeaGreen4Type2 ColorCode = "%c8;5;71"
	// CadetBlueType1 color hex = #5faf87
	CadetBlueType1 ColorCode = "%c8;5;72"
	// CadetBlueType2 color hex = #5fafaf
	CadetBlueType2 ColorCode = "%c8;5;73"
	// SkyBlue3 color hex = #5fafd7
	SkyBlue3 ColorCode = "%c8;5;74"
	// SteelBlue1Type1 color hex = #5fafff
	SteelBlue1Type1 ColorCode = "%c8;5;75"
	// Chartreuse3Type2 color hex = #5fd700
	Chartreuse3Type2 ColorCode = "%c8;5;76"
	// PaleGreen3Type1 color hex = #5fd75f
	PaleGreen3Type1 ColorCode = "%c8;5;77"
	// SeaGreen3 color hex = #5fd787
	SeaGreen3 ColorCode = "%c8;5;78"
	// Aquamarine3 color hex = #5fd7af
	Aquamarine3 ColorCode = "%c8;5;79"
	// MediumTurquoise color hex = #5fd7d7
	MediumTurquoise ColorCode = "%c8;5;80"
	// SteelBlue1Type2 color hex = #5fd7ff
	SteelBlue1Type2 ColorCode = "%c8;5;81"
	// Chartreuse2Type1 color hex = #5fff00
	Chartreuse2Type1 ColorCode = "%c8;5;82"
	// SeaGreen2 color hex = #5fff5f
	SeaGreen2 ColorCode = "%c8;5;83"
	// SeaGreen1Type1 color hex = #5fff87
	SeaGreen1Type1 ColorCode = "%c8;5;84"
	// SeaGreen1Type2 color hex = #5fffaf
	SeaGreen1Type2 ColorCode = "%c8;5;85"
	// Aquamarine1Type1 color hex = #5fffd7
	Aquamarine1Type1 ColorCode = "%c8;5;86"
	// DarkSlateGray2 color hex = #5fffff
	DarkSlateGray2 ColorCode = "%c8;5;87"
	// DarkRedType2 color hex = #870000
	DarkRedType2 ColorCode = "%c8;5;88"
	// DeepPink4Type2 color hex = #87005f
	DeepPink4Type2 ColorCode = "%c8;5;89"
	// DarkMagentaType1 color hex = #870087
	DarkMagentaType1 ColorCode = "%c8;5;90"
	// DarkMagentaType2 color hex = #8700af
	DarkMagentaType2 ColorCode = "%c8;5;91"
	// DarkVioletType1 color hex = #8700d7
	DarkVioletType1 ColorCode = "%c8;5;92"
	// PurpleType2 color hex = #8700ff
	PurpleType2 ColorCode = "%c8;5;93"
	// Orange4Type2 color hex = #875f00
	Orange4Type2 ColorCode = "%c8;5;94"
	// LightPink4 color hex = #875f5f
	LightPink4 ColorCode = "%c8;5;95"
	// Plum4 color hex = #875f87
	Plum4 ColorCode = "%c8;5;96"
	// MediumPurple3Type1 color hex = #875faf
	MediumPurple3Type1 ColorCode = "%c8;5;97"
	// MediumPurple3Type2 color hex = #875fd7
	MediumPurple3Type2 ColorCode = "%c8;5;98"
	// SlateBlue1 color hex = #875fff
	SlateBlue1 ColorCode = "%c8;5;99"
	// Yellow4Type1 color hex = #878700
	Yellow4Type1 ColorCode = "%c8;5;100"
	// Wheat4 color hex = #87875f
	Wheat4 ColorCode = "%c8;5;101"
	// Grey53 color hex = #878787
	Grey53 ColorCode = "%c8;5;102"
	// LightSlateGrey color hex = #8787af
	LightSlateGrey ColorCode = "%c8;5;103"
	// MediumPurple color hex = #8787d7
	MediumPurple ColorCode = "%c8;5;104"
	// LightSlateBlue color hex = #8787ff
	LightSlateBlue ColorCode = "%c8;5;105"
	// Yellow4Type2 color hex = #87af00
	Yellow4Type2 ColorCode = "%c8;5;106"
	// DarkOliveGreen3Type1 color hex = #87af5f
	DarkOliveGreen3Type1 ColorCode = "%c8;5;107"
	// DarkSeaGreen color hex = #87af87
	DarkSeaGreen ColorCode = "%c8;5;108"
	// LightSkyBlue3Type1 color hex = #87afaf
	LightSkyBlue3Type1 ColorCode = "%c8;5;109"
	// LightSkyBlue3Type2 color hex = #87afd7
	LightSkyBlue3Type2 ColorCode = "%c8;5;110"
	// SkyBlue2 color hex = #87afff
	SkyBlue2 ColorCode = "%c8;5;111"
	// Chartreuse2Type2 color hex = #87d700
	Chartreuse2Type2 ColorCode = "%c8;5;112"
	// DarkOliveGreen3Type2 color hex = #87d75f
	DarkOliveGreen3Type2 ColorCode = "%c8;5;113"
	// PaleGreen3Type2 color hex = #87d787
	PaleGreen3Type2 ColorCode = "%c8;5;114"
	// DarkSeaGreen3Type1 color hex = #87d7af
	DarkSeaGreen3Type1 ColorCode = "%c8;5;115"
	// DarkSlateGray3 color hex = #87d7d7
	DarkSlateGray3 ColorCode = "%c8;5;116"
	// SkyBlue1 color hex = #87d7ff
	SkyBlue1 ColorCode = "%c8;5;117"
	// Chartreuse1 color hex = #87ff00
	Chartreuse1 ColorCode = "%c8;5;118"
	// LightGreenType1 color hex = #87ff5f
	LightGreenType1 ColorCode = "%c8;5;119"
	// LightGreenType2 color hex = #87ff87
	LightGreenType2 ColorCode = "%c8;5;120"
	// PaleGreen1Type1 color hex = #87ffaf
	PaleGreen1Type1 ColorCode = "%c8;5;121"
	// Aquamarine1Type2 color hex = #87ffd7
	Aquamarine1Type2 ColorCode = "%c8;5;122"
	// DarkSlateGray1 color hex = #87ffff
	DarkSlateGray1 ColorCode = "%c8;5;123"
	// Red3Type1 color hex = #af0000
	Red3Type1 ColorCode = "%c8;5;124"
	// DeepPink4Type3 color hex = #af005f
	DeepPink4Type3 ColorCode = "%c8;5;125"
	// MediumVioletRed color hex = #af0087
	MediumVioletRed ColorCode = "%c8;5;126"
	// Magenta3Type1 color hex = #af00af
	Magenta3Type1 ColorCode = "%c8;5;127"
	// DarkVioletType2 color hex = #af00d7
	DarkVioletType2 ColorCode = "%c8;5;128"
	// PurpleType3 color hex = #af00ff
	PurpleType3 ColorCode = "%c8;5;129"
	// DarkOrange3Type1 color hex = #af5f00
	DarkOrange3Type1 ColorCode = "%c8;5;130"
	// IndianRedType1 color hex = #af5f5f
	IndianRedType1 ColorCode = "%c8;5;131"
	// HotPink3Type1 color hex = #af5f87
	HotPink3Type1 ColorCode = "%c8;5;132"
	// MediumOrchid3 color hex = #af5faf
	MediumOrchid3 ColorCode = "%c8;5;133"
	// MediumOrchid color hex = #af5fd7
	MediumOrchid ColorCode = "%c8;5;134"
	// MediumPurple2Type1 color hex = #af5fff
	MediumPurple2Type1 ColorCode = "%c8;5;135"
	// DarkGoldenrod color hex = #af8700
	DarkGoldenrod ColorCode = "%c8;5;136"
	// LightSalmon3Type1 color hex = #af875f
	LightSalmon3Type1 ColorCode = "%c8;5;137"
	// RosyBrown color hex = #af8787
	RosyBrown ColorCode = "%c8;5;138"
	// Grey63 color hex = #af87af
	Grey63 ColorCode = "%c8;5;139"
	// MediumPurple2Type2 color hex = #af87d7
	MediumPurple2Type2 ColorCode = "%c8;5;140"
	// MediumPurple1 color hex = #af87ff
	MediumPurple1 ColorCode = "%c8;5;141"
	// Gold3Type1 color hex = #afaf00
	Gold3Type1 ColorCode = "%c8;5;142"
	// DarkKhaki color hex = #afaf5f
	DarkKhaki ColorCode = "%c8;5;143"
	// NavajoWhite3 color hex = #afaf87
	NavajoWhite3 ColorCode = "%c8;5;144"
	// Grey69 color hex = #afafaf
	Grey69 ColorCode = "%c8;5;145"
	// LightSteelBlue3 color hex = #afafd7
	LightSteelBlue3 ColorCode = "%c8;5;146"
	// LightSteelBlue color hex = #afafff
	LightSteelBlue ColorCode = "%c8;5;147"
	// Yellow3Type1 color hex = #afd700
	Yellow3Type1 ColorCode = "%c8;5;148"
	// DarkOliveGreen3Type3 color hex = #afd75f
	DarkOliveGreen3Type3 ColorCode = "%c8;5;149"
	// DarkSeaGreen3Type2 color hex = #afd787
	DarkSeaGreen3Type2 ColorCode = "%c8;5;150"
	// DarkSeaGreen2Type1 color hex = #afd7af
	DarkSeaGreen2Type1 ColorCode = "%c8;5;151"
	// LightCyan3 color hex = #afd7d7
	LightCyan3 ColorCode = "%c8;5;152"
	// LightSkyBlue1 color hex = #afd7ff
	LightSkyBlue1 ColorCode = "%c8;5;153"
	// GreenYellow color hex = #afff00
	GreenYellow ColorCode = "%c8;5;154"
	// DarkOliveGreen2 color hex = #afff5f
	DarkOliveGreen2 ColorCode = "%c8;5;155"
	// PaleGreen1Type2 color hex = #afff87
	PaleGreen1Type2 ColorCode = "%c8;5;156"
	// DarkSeaGreen2Type2 color hex = #afffaf
	DarkSeaGreen2Type2 ColorCode = "%c8;5;157"
	// DarkSeaGreen1Type1 color hex = #afffd7
	DarkSeaGreen1Type1 ColorCode = "%c8;5;158"
	// PaleTurquoise1 color hex = #afffff
	PaleTurquoise1 ColorCode = "%c8;5;159"
	// Red3Type2 color hex = #d70000
	Red3Type2 ColorCode = "%c8;5;160"
	// DeepPink3Type1 color hex = #d7005f
	DeepPink3Type1 ColorCode = "%c8;5;161"
	// DeepPink3Type2 color hex = #d70087
	DeepPink3Type2 ColorCode = "%c8;5;162"
	// Magenta3Type2 color hex = #d700af
	Magenta3Type2 ColorCode = "%c8;5;163"
	// Magenta3Type3 color hex = #d700d7
	Magenta3Type3 ColorCode = "%c8;5;164"
	// Magenta2Type1 color hex = #d700ff
	Magenta2Type1 ColorCode = "%c8;5;165"
	// DarkOrange3Type2 color hex = #d75f00
	DarkOrange3Type2 ColorCode = "%c8;5;166"
	// IndianRedType2 color hex = #d75f5f
	IndianRedType2 ColorCode = "%c8;5;167"
	// HotPink3Type2 color hex = #d75f87
	HotPink3Type2 ColorCode = "%c8;5;168"
	// HotPink2 color hex = #d75faf
	HotPink2 ColorCode = "%c8;5;169"
	// Orchid color hex = #d75fd7
	Orchid ColorCode = "%c8;5;170"
	// MediumOrchid1Type1 color hex = #d75fff
	MediumOrchid1Type1 ColorCode = "%c8;5;171"
	// Orange3 color hex = #d78700
	Orange3 ColorCode = "%c8;5;172"
	// LightSalmon3Type2 color hex = #d7875f
	LightSalmon3Type2 ColorCode = "%c8;5;173"
	// LightPink3 color hex = #d78787
	LightPink3 ColorCode = "%c8;5;174"
	// Pink3 color hex = #d787af
	Pink3 ColorCode = "%c8;5;175"
	// Plum3 color hex = #d787d7
	Plum3 ColorCode = "%c8;5;176"
	// Violet color hex = #d787ff
	Violet ColorCode = "%c8;5;177"
	// Gold3Type2 color hex = #d7af00
	Gold3Type2 ColorCode = "%c8;5;178"
	// LightGoldenrod3 color hex = #d7af5f
	LightGoldenrod3 ColorCode = "%c8;5;179"
	// Tan color hex = #d7af87
	Tan ColorCode = "%c8;5;180"
	// MistyRose3 color hex = #d7afaf
	MistyRose3 ColorCode = "%c8;5;181"
	// Thistle3 color hex = #d7afd7
	Thistle3 ColorCode = "%c8;5;182"
	// Plum2 color hex = #d7afff
	Plum2 ColorCode = "%c8;5;183"
	// Yellow3Type2 color hex = #d7d700
	Yellow3Type2 ColorCode = "%c8;5;184"
	// Khaki3 color hex = #d7d75f
	Khaki3 ColorCode = "%c8;5;185"
	// LightGoldenrod2 color hex = #d7d787
	LightGoldenrod2 ColorCode = "%c8;5;186"
	// LightYellow3 color hex = #d7d7af
	LightYellow3 ColorCode = "%c8;5;187"
	// Grey84 color hex = #d7d7d7
	Grey84 ColorCode = "%c8;5;188"
	// LightSteelBlue1 color hex = #d7d7ff
	LightSteelBlue1 ColorCode = "%c8;5;189"
	// Yellow2 color hex = #d7ff00
	Yellow2 ColorCode = "%c8;5;190"
	// DarkOliveGreen1Type1 color hex = #d7ff5f
	DarkOliveGreen1Type1 ColorCode = "%c8;5;191"
	// DarkOliveGreen1Type2 color hex = #d7ff87
	DarkOliveGreen1Type2 ColorCode = "%c8;5;192"
	// DarkSeaGreen1Type2 color hex = #d7ffaf
	DarkSeaGreen1Type2 ColorCode = "%c8;5;193"
	// Honeydew2 color hex = #d7ffd7
	Honeydew2 ColorCode = "%c8;5;194"
	// LightCyan1 color hex = #d7ffff
	LightCyan1 ColorCode = "%c8;5;195"
	// Red1 color hex = #ff0000
	Red1 ColorCode = "%c8;5;196"
	// DeepPink2 color hex = #ff005f
	DeepPink2 ColorCode = "%c8;5;197"
	// DeepPink1Type1 color hex = #ff0087
	DeepPink1Type1 ColorCode = "%c8;5;198"
	// DeepPink1Type2 color hex = #ff00af
	DeepPink1Type2 ColorCode = "%c8;5;199"
	// Magenta2Type2 color hex = #ff00d7
	Magenta2Type2 ColorCode = "%c8;5;200"
	// Magenta1 color hex = #ff00ff
	Magenta1 ColorCode = "%c8;5;201"
	// OrangeRed1 color hex = #ff5f00
	OrangeRed1 ColorCode = "%c8;5;202"
	// IndianRed1Type1 color hex = #ff5f5f
	IndianRed1Type1 ColorCode = "%c8;5;203"
	// IndianRed1Type2 color hex = #ff5f87
	IndianRed1Type2 ColorCode = "%c8;5;204"
	// HotPinkType1 color hex = #ff5faf
	HotPinkType1 ColorCode = "%c8;5;205"
	// HotPinkType2 color hex = #ff5fd7
	HotPinkType2 ColorCode = "%c8;5;206"
	// MediumOrchid1Type2 color hex = #ff5fff
	MediumOrchid1Type2 ColorCode = "%c8;5;207"
	// DarkOrange color hex = #ff8700
	DarkOrange ColorCode = "%c8;5;208"
	// Salmon1 color hex = #ff875f
	Salmon1 ColorCode = "%c8;5;209"
	// LightCoral color hex = #ff8787
	LightCoral ColorCode = "%c8;5;210"
	// PaleVioletRed1 color hex = #ff87af
	PaleVioletRed1 ColorCode = "%c8;5;211"
	// Orchid2 color hex = #ff87d7
	Orchid2 ColorCode = "%c8;5;212"
	// Orchid1 color hex = #ff87ff
	Orchid1 ColorCode = "%c8;5;213"
	// Orange1 color hex = #ffaf00
	Orange1 ColorCode = "%c8;5;214"
	// SandyBrown color hex = #ffaf5f
	SandyBrown ColorCode = "%c8;5;215"
	// LightSalmon1 color hex = #ffaf87
	LightSalmon1 ColorCode = "%c8;5;216"
	// LightPink1 color hex = #ffafaf
	LightPink1 ColorCode = "%c8;5;217"
	// Pink1 color hex = #ffafd7
	Pink1 ColorCode = "%c8;5;218"
	// Plum1 color hex = #ffafff
	Plum1 ColorCode = "%c8;5;219"
	// Gold1 color hex = #ffd700
	Gold1 ColorCode = "%c8;5;220"
	// LightGoldenrod2Type1 color hex = #ffd75f
	LightGoldenrod2Type1 ColorCode = "%c8;5;221"
	// LightGoldenrod2Type2 color hex = #ffd787
	LightGoldenrod2Type2 ColorCode = "%c8;5;222"
	// NavajoWhite1 color hex = #ffd7af
	NavajoWhite1 ColorCode = "%c8;5;223"
	// MistyRose1 color hex = #ffd7d7
	MistyRose1 ColorCode = "%c8;5;224"
	// Thistle1 color hex = #ffd7ff
	Thistle1 ColorCode = "%c8;5;225"
	// Yellow1 color hex = #ffff00
	Yellow1 ColorCode = "%c8;5;226"
	// LightGoldenrod1 color hex = #ffff5f
	LightGoldenrod1 ColorCode = "%c8;5;227"
	// Khaki1 color hex = #ffff87
	Khaki1 ColorCode = "%c8;5;228"
	// Wheat1 color hex = #ffffaf
	Wheat1 ColorCode = "%c8;5;229"
	// Cornsilk1 color hex = #ffffd7
	Cornsilk1 ColorCode = "%c8;5;230"
	// Grey100 color hex = #ffffff
	Grey100 ColorCode = "%c8;5;231"
	// Grey3 color hex = #080808
	Grey3 ColorCode = "%c8;5;232"
	// Grey7 color hex = #121212
	Grey7 ColorCode = "%c8;5;233"
	// Grey11 color hex = #1c1c1c
	Grey11 ColorCode = "%c8;5;234"
	// Grey15 color hex = #262626
	Grey15 ColorCode = "%c8;5;235"
	// Grey19 color hex = #303030
	Grey19 ColorCode = "%c8;5;236"
	// Grey23 color hex = #3a3a3a
	Grey23 ColorCode = "%c8;5;237"
	// Grey27 color hex = #444444
	Grey27 ColorCode = "%c8;5;238"
	// Grey30 color hex = #4e4e4e
	Grey30 ColorCode = "%c8;5;239"
	// Grey35 color hex = #585858
	Grey35 ColorCode = "%c8;5;240"
	// Grey39 color hex = #626262
	Grey39 ColorCode = "%c8;5;241"
	// Grey42 color hex = #6c6c6c
	Grey42 ColorCode = "%c8;5;242"
	// Grey46 color hex = #767676
	Grey46 ColorCode = "%c8;5;243"
	// Grey50 color hex = #808080
	Grey50 ColorCode = "%c8;5;244"
	// Grey54 color hex = #8a8a8a
	Grey54 ColorCode = "%c8;5;245"
	// Grey58 color hex = #949494
	Grey58 ColorCode = "%c8;5;246"
	// Grey62 color hex = #9e9e9e
	Grey62 ColorCode = "%c8;5;247"
	// Grey66 color hex = #a8a8a8
	Grey66 ColorCode = "%c8;5;248"
	// Grey70 color hex = #b2b2b2
	Grey70 ColorCode = "%c8;5;249"
	// Grey74 color hex = #bcbcbc
	Grey74 ColorCode = "%c8;5;250"
	// Grey78 color hex = #c6c6c6
	Grey78 ColorCode = "%c8;5;251"
	// Grey82 color hex = #d0d0d0
	Grey82 ColorCode = "%c8;5;252"
	// Grey85 color hex = #dadada
	Grey85 ColorCode = "%c8;5;253"
	// Grey89 color hex = #e4e4e4
	Grey89 ColorCode = "%c8;5;254"
	// Grey93 color hex = #eeeeee
	Grey93 ColorCode = "%c8;5;255"
)

func MakeColorCode(r, g, b int) ColorCode {
	return ColorCode(fmt.Sprintf("%%c8;2;%d;%d;%d", r%256, g%256, b%256))
}
