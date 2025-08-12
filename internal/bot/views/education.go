package views

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
)

const test = `
`

func CreateEducationView() *telego.View {
	return &View{
		Text:     text,
		Keyboard: telegoutil.InlineKeyboard(
			telegoutil.InlineKeyboardButton("Back").WithCallbackData("back"),
		),
	}
