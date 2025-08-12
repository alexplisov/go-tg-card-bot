package handlers

import (
	"github.com/alexplisov/go-tg-card-bot/internal/bot/views"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

func StartHandler(ctx *telegohandler.Context, update telego.Update) error {
	view := views.CreateMainMenu()
	message := telegoutil.
		Message(telegoutil.ID(update.Message.Chat.ID), view.Text).
		WithParseMode(telego.ModeHTML).
		WithReplyMarkup(view.Keyboard)
	if _, err := ctx.Bot().SendMessage(ctx, message); err != nil {
		return err
	}
	return nil
}
