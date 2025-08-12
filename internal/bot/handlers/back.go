package handlers

import (
	"github.com/alexplisov/go-tg-card-bot/internal/bot/views"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

func HandleInlineMenuBack(ctx *telegohandler.Context, query telego.CallbackQuery) error {
	view := views.CreateMainMenu()
	if _, err := ctx.Bot().EditMessageText(
		ctx.Context(),
		telegoutil.EditMessageText(
			query.Message.GetChat().ChatID(),
			query.Message.Message().MessageID,
			view.Text,
		).WithParseMode(telego.ModeHTML),
	); err != nil {
		return err
	}
	if _, err := ctx.Bot().EditMessageReplyMarkup(
		ctx.Context(),
		telegoutil.EditMessageReplayMarkup(
			query.Message.GetChat().ChatID(),
			query.Message.Message().MessageID,
			view.Keyboard,
		),
	); err != nil {
		return err
	}
	return ctx.Bot().AnswerCallbackQuery(ctx, telegoutil.CallbackQuery(query.ID))
}
