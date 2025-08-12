package handlers

import (
	"github.com/alexplisov/go-tg-card-bot/internal/bot/views"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

func EducationHandler(ctx *telegohandler.Context, query telego.CallbackQuery) error {
	view := views.CreateEducationView()
	if _, err := ctx.Bot().EditMessageText(
		ctx.Context(),
		telegoutil.EditMessageText(
			query.Message.GetChat().ChatID(),
			query.Message.Message().MessageID,
			view.Text,
		),
	); err != nil {
		return err
	}
	if _, err := ctx.Bot().EditMessageReplyMarkup(
		ctx.Context(),
		telegoutil.EditMessageReplayMarkup()(
			query.Message.GetChat().ChatID(),
			query.Message.Message().MessageID,
			view.ReplyMarkup,
		),
	); err != nil {
		return err
	}
	return nil
}
