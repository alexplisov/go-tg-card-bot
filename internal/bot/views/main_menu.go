package views

import (
	"github.com/mymmrac/telego/telegoutil"
)

func CreateMainMenu() *View {
	return &View{
		Text: `
<b>👋 Приветствую!</b>

🙂 Меня зовут Алексей.

🤖 Я разрабатываю Telegram-ботов на TypeScript, Go и Python.

💎 Если хочешь себе бота, то оставь заявку и я помогу тебе с этим.
		`,
		Keyboard: telegoutil.InlineKeyboard(
			telegoutil.InlineKeyboardRow(
				telegoutil.InlineKeyboardButton("Образование").WithCallbackData("education"),
				telegoutil.InlineKeyboardButton("Оставить заявку").WithCallbackData("order"),
			),
		),
	}
}
