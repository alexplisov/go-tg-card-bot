package bot

import (
	"context"

	"github.com/alexplisov/go-tg-card-bot/internal/bot/handlers"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)

func SetupBotHandler(token string) (*telegohandler.BotHandler, error) {
	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		return nil, err
	}
	updates, _ := bot.UpdatesViaLongPolling(context.Background(), nil)
	handler, err := telegohandler.NewBotHandler(bot, updates)
	handler.Handle(handlers.StartHandler, telegohandler.CommandEqual("start"))
	handler.HandleCallbackQuery(handlers.EducationHandler, telegohandler.CallbackDataEqual("education"))
	handler.HandleCallbackQuery(handlers.OrderHandler, telegohandler.CallbackDataEqual("education"))
	return handler, nil
}
