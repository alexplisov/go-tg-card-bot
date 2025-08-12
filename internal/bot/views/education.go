package views

import (
	"github.com/mymmrac/telego/telegoutil"
)

func CreateEducationView() *View {
	return &View{
		Text: `
👨‍🎓 Я имею диплом бакалавра в области прикладной информатики от Уральского Государственного Экономического Университета.
		`,
		Keyboard: telegoutil.InlineKeyboard(
			telegoutil.InlineKeyboardRow(
				telegoutil.InlineKeyboardButton("Back").WithCallbackData("back"),
			),
		),
	}
}
