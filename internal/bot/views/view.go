package views

import "github.com/mymmrac/telego"

type View struct {
	Text     string
	Keyboard *telego.InlineKeyboardMarkup
}
