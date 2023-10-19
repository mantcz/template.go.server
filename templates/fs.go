package templates

import (
	"embed"
)

//go:embed *.gohtml
var FS embed.FS

var Emoji = map[string]string{
	"smile": "😀",
	"laugh": "😆",
	"wink": "😉",
	"grin": "😁",
	"rofl": "🤣",
	"joy": "😂",
	"slight_smile": "🙂",
	"upside_down": "🙃",
	"relaxed": "☺️",
	"blush": "😊",
	"heart_eyes": "😍",
	"table_tennis": "🏓",
	"basketball": "🏀",
	"football": "⚽️",
	"baseball": "⚾️",
	"tennis": "🎾",
	"volleyball": "🏐",
	"rugby": "🏉",
	"8ball": "🎱",
	"bowling": "🎳",
	"carrot": "🥕",
	"hot_pepper": "🌶",
	"donut": "🍩",
	"whale": "🐳",
	"octopus": "🐙",
}
